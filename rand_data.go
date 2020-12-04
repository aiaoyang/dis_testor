package main

import (
	"math/rand"
	"time"

	uuid "github.com/satori/go.uuid"
)

// RandData 生成随机数据
func RandData() []TestData {
	res := make([]TestData, 0)
	rand.Seed(time.Now().Unix())

	for i := 0; i < 100; i++ {

		uid := uuid.Must(uuid.NewV4())
		t := TestData{
			UserName: randUserName(6),
			Password: randPassword(10),
			UUID:     uid.String(),
		}
		res = append(res, t)
	}
	return res
}

var pack = [3][2]int{{65, 26}, {97, 26}, {48, 10}}

var base, scope int

func randUserName(userNameLen int) string {
	buf := make([]byte, userNameLen)
	randStr(userNameLen-1, 3, &buf)
	return string(buf)
}

func randPassword(passwordLen int) string {
	buf := make([]byte, passwordLen)
	randStr(passwordLen-1, 3, &buf)
	return string(buf)
}

func randStr(index int, kind int, res *[]byte) {
	if kind > 2 {
		if index < 0 {
			return
		}
		rand.Seed(time.Now().UnixNano())

		// fmt.Printf("index : %d\n", index)
		k := rand.Intn(3)
		(*res)[index] = byte(pack[k][0] + rand.Intn(pack[k][1]))
		randStr(index-1, kind, res)
	}
}
