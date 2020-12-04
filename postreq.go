package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// SendReq 发送 dis 数据
func SendReq(req *http.Request) {
	c := http.Client{}

	resp, err := c.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	disResp := DISResponse{}

	content, err := ioutil.ReadAll(resp.Body)

	resp.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(content, &disResp)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("HW DIS response content : %d\n", disResp.FailedRecordCount)

}
