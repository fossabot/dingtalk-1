/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package dingtalk

import (
	"github.com/zhaoyunxing92/dingtalk/v2/domain/message"
	"testing"
)
import "github.com/zhaoyunxing92/dingtalk/v2/request"
import "github.com/stretchr/testify/assert"

func TestDingTalk_SendTemplateMessage(t *testing.T) {

	res, err := isv.SendTemplateMessage(
		request.NewSendTemplateMessage(1332307896, "80424a44cb444c53a071aae34c0fd140").
			SetUserIds("manager7556", "manager7556").
			SetData("msg", "发送模板消息测试").
			Build())

	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestDingTalk_GetMessageProgress(t *testing.T) {

	res, err := isv.GetMessageProgress(1332307896, 474307154979)

	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestDingTalk_SendCorpConversationMessage_Text(t *testing.T) {
	msg := message.NewTextMessage("hello dubbo-go")

	res, err := isv.SendCorpConversationMessage(
		request.NewCorpConversationMessage(msg).
			SetUserIds("manager7556").
			Build())

	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestDingTalk_SendCorpConversationMessage_Link(t *testing.T) {

	msg := message.NewLinkMessage("消息标题，建议100字符以内",
		"消息描述，建议500字符以内", "@lADOADmaWMzazQKA", "https://github.com/zhaoyunxing92/dingtalk")

	res, err := isv.SendCorpConversationMessage(
		request.NewCorpConversationMessage(msg).
			SetUserIds("manager7556").
			Build())

	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestDingTalk_SendCorpConversationMessage_Image(t *testing.T) {

	msg := message.NewImageMessages("@lADOADmaWMzazQKA")

	res, err := isv.SendCorpConversationMessage(
		request.NewCorpConversationMessage(msg).
			SetUserIds("manager7556").
			Build())

	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestDingTalk_SendCorpConversationMessage_File(t *testing.T) {
	msg := message.NewFileMessage("@lADOADmaWMzazQKA")

	res, err := isv.SendCorpConversationMessage(
		request.NewCorpConversationMessage(msg).
			SetUserIds("manager7556").
			Build())

	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestDingTalk_SendCorpConversationMessage_Voice(t *testing.T) {
	voice := message.NewVoiceMessage("@lADOADmaWMzazQKA", 10)

	res, err := client.SendCorpConversationMessage(
		request.NewCorpConversationMessage(voice).
			SetUserIds("manager7556").
			Build())

	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestDingTalk_SendCorpConversationMessage_OA(t *testing.T) {
	msg := message.NewOaMessage("消息的头部标题", "FFBBBBBB", "http://dingtalk.com",
		"http://dingtalk.com")

	msg.Body.Title = "消息体标题"
	msg.Body.Content = "消息体的内容，最多显示3行"
	msg.Body.MediaId = "@lADOADmaWMzazQKA"
	msg.Body.Forms = []message.Form{
		{Key: "姓名", Value: "张三"},
		{Key: "年龄", Value: "20"},
		{Key: "身高", Value: "1.8m"},
		{Key: "体重", Value: "70kg"},
		{Key: "学历", Value: "本科"},
		{Key: "爱好", Value: "打球、听歌"},
	}
	msg.Rich.Num = "15.6"
	msg.Rich.Unit = "元"

	msg.Author = "李四"

	msg.StatusBar.Value = "进行中"
	msg.StatusBar.BackColor = "0xFFF65E5E"

	res, err := isv.SendCorpConversationMessage(
		request.NewCorpConversationMessage(msg).
			SetUserIds("manager7556").
			Build())

	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestDingTalk_SendCorpConversationMessage_Markdown(t *testing.T) {
	markdown := `# 这是支持markdown的文本
## 标题2
* 列表1

![alt 啊](https://img.alicdn.com/tps/TB1XLjqNVXXXXc4XVXXXXXXXXXX-170-64.png)`

	msg := message.NewMarkDownMessage("首屏会话透出的展示内容", markdown)

	res, err := client.SendCorpConversationMessage(
		request.NewCorpConversationMessage(msg).
			SetUserIds("manager164").
			Build())

	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestDingTalk_SendCorpConversationMessage_Card(t *testing.T) {
	content := `# 这是支持markdown的文本
## 标题2
* 列表1

![alt 啊](https://img.alicdn.com/tps/TB1XLjqNVXXXXc4XVXXXXXXXXXX-170-64.png)`

	msg := message.NewCardMessage(content)

	msg.Title = "是透出到会话列表和通知的文案"
	msg.SingleTitle = "查看详情"
	msg.SingleUrl = "https://open.dingtalk.com"

	res, err := client.SendCorpConversationMessage(
		request.NewCorpConversationMessage(msg).
			SetUserIds("manager164").
			Build())

	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestDingTalk_SendCorpConversationMessage_CardButtons(t *testing.T) {
	content := `# 这是支持markdown的文本
## 标题2
* 列表1

![alt 啊](https://img.alicdn.com/tps/TB1XLjqNVXXXXc4XVXXXXXXXXXX-170-64.png)`

	msg := message.NewCardMessage(content)

	msg.Title = "是透出到会话列表和通知的文案"
	msg.BtnOrientation = "1"
	msg.CardButtons = []message.CardButton{
		{"一个按钮", "https://www.taobao.com"},
		{"两个按钮", "https://www.tmall.com"},
	}

	res, err := client.SendCorpConversationMessage(
		request.NewCorpConversationMessage(msg).
			SetUserIds("manager164").
			Build())

	assert.Nil(t, err)
	assert.NotNil(t, res)
}
