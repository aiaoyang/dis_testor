package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// accesskey accessid 类方法
func sign(accessKey, accessID string, url string) {

}

// ==============================================
type tokenRequest struct {
	Auth `json:"auth"`
}
type Auth struct {
	Identity `json:"identity"`
	Scope    `json:"scope"`
}
type Identity struct {
	Methods  []string `json:"methods"`
	Password `json:"password"`
}
type Password struct {
	User `json:"user"`
}
type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Domain   `json:"domain"`
}
type Domain struct {
	Name string `json:"name"`
}
type Scope struct {
	Domain  `json:"domain"`
	Project `json:"project"`
}
type Project struct {
	ID string `json:"id"`
}

// CacheToken token 类方法
type CacheToken struct {
	token      string
	expireTime time.Duration
	genTime    time.Time
	expired    bool
	tokenRequest
}

// Token 获取缓存中的token
func (c *CacheToken) Token() string {

	if c.expired || time.Now().Add(-c.expireTime).After(c.genTime) {
		log.Printf("none token ,gen one\n")
		c.token = c.reqToken()
		c.genTime = time.Now()
		c.expired = false
	}
	return c.token
}

var tokenURL = "https://iam.cn-east-3.myhuaweicloud.com/v3/auth/tokens"

// 发送获取token请求
func (c *CacheToken) reqToken() string {

	rawStr, err := json.Marshal(c.tokenRequest)
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{}
	req, err := http.NewRequest("POST", tokenURL, bytes.NewReader(rawStr))
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	return resp.Header.Get("X-Subject-Token")
}

// NewTokenWithCache 生成带缓存的token结构体
func NewTokenWithCache(userName, password, domain, projectID string) CacheToken {

	reqContent := tokenRequest{}
	reqContent.Auth.Methods = []string{"password"}
	reqContent.Identity.Password.User.Name = userName
	reqContent.Identity.Password.User.Password = password
	reqContent.Identity.Password.User.Domain.Name = domain
	reqContent.Scope.Project.ID = projectID

	cacheToken := CacheToken{
		expireTime:   time.Second * 86400,
		genTime:      time.Now(),
		expired:      true,
		tokenRequest: reqContent,
	}

	cacheToken.token = cacheToken.reqToken()
	cacheToken.expired = false

	return cacheToken

}
