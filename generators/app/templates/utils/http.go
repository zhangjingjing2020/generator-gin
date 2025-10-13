package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type httpHandler struct {
}

var HttpHandler = new(httpHandler)

// GenericHTTPCallWithHeaders 是一个带有自定义头部信息的通用HTTP调用方法
func (h *httpHandler) GenericHTTPCallWithHeaders(method, url string, requestBody []byte, headers map[string]string, response interface{}) (err error) {
	// 创建一个新的HTTP请求对象
	req, err := http.NewRequest(method, url, bytes.NewBufferString(string(requestBody)))
	if err != nil {
		err = fmt.Errorf("HTTP请求创建失败:%s", err)
		return
	}

	// 设置默认的Content-Type头部信息
	req.Header.Set("Content-Type", "application/json")

	// 如果用户提供了自定义的头部信息，则覆盖默认值
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// 发起HTTP请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		err = fmt.Errorf("HTTP请求失败:%s", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应结果
	body, err := io.ReadAll(resp.Body)
	// fmt.Print("body===>>", string(body))
	if err != nil {
		err = fmt.Errorf("读取响应结果失败:%s", err)
		return
	}

	// 解析响应结果
	if err = json.Unmarshal(body, response); err != nil {
		err = fmt.Errorf("解析响应结果失败:%s", err.Error())
		return
	}

	return
}
