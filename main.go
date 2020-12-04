package main

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
)

// TestData 测试用数据
type TestData struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
	UUID     string `json:"uuid"`
}

// HuaweiCloudAuth 华为云账号认证配置
type HuaweiCloudAuth struct {
	UserName      string `yaml:"userName"`
	Password      string `yaml:"password"`
	Domain        string `yaml:"domain"`
	ProjectID     string `yaml:"projectId"`
	DisURL        string `yaml:"disUrl"`
	HuaweiDisConf `yaml:"huaweiDisConf"`
}

// HuaweiDisConf 华为云DIS通道信息
type HuaweiDisConf struct {
	StreamName string `yaml:"steamName"`
	StreamID   string `yaml:"steamId"`
}

func main() {
	log.SetFlags(log.Llongfile | log.Ldate)
	v := viper.New()
	v.SetConfigFile("conf.yaml")
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	conf := HuaweiCloudAuth{}
	v.Unmarshal(&conf)
	log.Printf("viper: %v\nconfig : %v\n", v.AllKeys(), conf)
	// return

	// 生成随机数据
	tData := RandData()

	// 将数据转换成dis记录
	records := NewRecords("", "", "", tData...)

	// 生成dis结构内容
	disContent := NewDISContent(conf.StreamName, conf.StreamID)
	// disContent := NewDISContent("dis-pFKZ", "eVHrcNWlNR2XHybkUdd")

	// 将记录添加到dis结构内容中
	disContent.AddRecord(records...)

	// 获取tokencache
	token := NewTokenWithCache(conf.UserName, conf.Password, conf.Domain, conf.ProjectID)

	count := 0
	// for {

	fmt.Printf("count: %d\n", count)
	count++
	// 生成http请求
	req := NewDISRequest("post", conf.DisURL, disContent, token.Token())

	// 发送请求
	SendReq(req)
	time.Sleep(time.Second)
	// }

}
