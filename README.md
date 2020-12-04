```

	// 生成随机数据
	tData := RandData()

	// 将数据转换成dis记录
	records := NewRecords("", "", "", tData...)

	// 生成dis结构内容
	disContent := NewDISContent("hw_channel_name", "hw_channel_id")

	// 将记录添加到dis结构内容中
	disContent.AddRecord(records...)

	// 获取tokencache
	token := NewTokenWithCache("username", "password", "hw_account_name", "hw_project_id")

	// 生成http请求
	req := NewDISRequest("post", disURL, disContent, token.Token())

	// 发送请求
	SendReq(req)

```