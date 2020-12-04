package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// DISContent dis 请求结构体
type DISContent struct {
	// Console控制台已创建的通道名称。通道名称由字母、数字、下划线和中划线组成，长度为1～64字符。
	StreamName string `json:"stream_name"`

	// 通道唯一标识符。当使用stream_name没有找到对应通道且stream_id不为空时，会使用stream_id去查找通道。
	StreamID string `json:"stream_id"`

	// List record，record为对象结构体。
	Records []Record `json:"records"`
}

// Record dis 条目
type Record struct {
	// Base64-encoded binary data object
	// 需要上传的数据。上传的数据为序列化之后的二进制数据（序列化后是Base64编码格式）。
	Data string `json:"data"`

	// 	用于明确数据需要写入分区的哈希值，此哈希值将覆盖“partition_key”的哈希值。取值范围：0~long.max
	ExplicitHashKey string `json:"explicit_hash_key"`

	// 分区的唯一标识符。
	PartitionID string `json:"partition_id"`

	// 	数据将写入的分区。
	PartitionKey string `json:"partition_key"`
}

// DISResponse dis 返回结构体
type DISResponse struct {
	FailedRecordCount int               `json:"failed_record_count"`
	Records           []ResponseRecords `json:"records"`
}

// ResponseRecords 返回记录结构体
type ResponseRecords struct {
	// 错误码。
	ErrorCode string `json:"error_code"`

	// 错误消息。
	ErrorMessage string `json:"error_message"`

	// 分区ID。
	PartitionID string `json:"partition_id"`

	// 序列号。序列号是每个记录的唯一标识符。序列号由DIS在数据生产者调用PutRecords操作以添加数据到DIS数据通道时DIS服务自动分配的。同一分区键的序列号通常会随时间变化增加。PutRecords请求之间的时间段越长，序列号越大。
	SequenceNumber string `json:"sequence_number"`
}

// AddRecord add Record to request
func (d *DISContent) AddRecord(record ...Record) {
	d.Records = append(d.Records, record...)
}

func (r *DISResponse) String() string {
	jsonStr, err := json.Marshal(r)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%s", jsonStr)
}

// NewRecords NewRecords
func NewRecords(ExplicitHashKey, PartitionID, PartitionKey string, Data ...TestData) []Record {

	// 初始化空 buffer
	buf := bytes.NewBuffer(nil)

	// 初始化 b64 加密实例
	b64Encoder := base64.NewEncoder(base64.StdEncoding, buf)
	// defer b64Encoder.Close()

	records := make([]Record, 0)

	for _, data := range Data {

		// 将record结构体数据转为 >> json格式，[]byte数据类型 <<
		jstr, err := json.Marshal(data)

		if err != nil {
			log.Fatal(err)
		}
		// fmt.Printf("%s\n", jstr)
		// log.Printf("data : %s\n", jstr)

		// 初始化一个空 record 实例
		record := Record{}

		// 将 json []byte 内容进行 b64 加密
		b64Encoder.Write(jstr)
		b64Encoder.Close()
		// 将 b64加密的内容转为 string 格式，存入 record 实例
		record.Data = buf.String()

		// fmt.Printf("buf len : %d\n", buf.Len())
		fmt.Println(record.Data)

		// 重置 buffer
		buf.Reset()

		// 将 record 存入records
		records = append(records, record)

	}

	return records

}

// NewDISRequest 初始化dis请求为http.Request结构体
func NewDISRequest(method, value string, disReq DISContent, token string) *http.Request {

	switch strings.ToUpper(method) {
	case "POST":
		method = "POST"

	case "GET":
		method = "GET"

	case "DELETE":
		method = "DELETE"

	default:
		return nil

	}

	content, err := json.Marshal(disReq)

	if err != nil {
		log.Fatal(err)
	}

	body := bytes.NewReader(content)

	req, err := http.NewRequest(method, value, body)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("X-Auth-Token", token)

	req.Header.Set("Content-Type", "application/json")

	return req
}

// NewDISContent NewDISContent
func NewDISContent(streamName, streamID string) DISContent {
	return DISContent{
		StreamName: streamName,
		StreamID:   streamID,
		Records:    make([]Record, 0),
	}
}
