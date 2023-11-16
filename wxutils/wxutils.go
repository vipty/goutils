// wxutils/wxutils.go
package wxutils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// SendWX 发送消息到企业微信机器人
func SendWX(key, message string) error {
	webhookURL := "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=" + key

	// 构建消息体
	msg := map[string]interface{}{
		"msgtype": "markdown",
		"markdown": map[string]string{
			"content": message,
		},
	}

	// 将消息内容转换为 JSON 格式
	messageData, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("JSON 格式化错误: %v", err)
	}

	// 发送 HTTP POST 请求到企业微信机器人
	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(messageData))
	if err != nil {
		return fmt.Errorf("HTTP 请求错误: %v", err)
	}
	defer resp.Body.Close()

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP 请求失败，状态码：%d", resp.StatusCode)
	}

	// fmt.Println("消息发送成功！")
	return nil
}
