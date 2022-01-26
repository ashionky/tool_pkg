/**
 * @Author pibing
 * @create 2022/1/25 3:05 PM
 */

package util

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

// 发送GET请求
func HttpGet(url string) (string, error) {
	// 超时时间：默认5秒
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)

	return string(result), err
}

// 发送POST请求
func HttpPost(url string, data interface{}, contentType string) (string, error) {
	if contentType == "" {
		contentType = "application/json"
	}
	jsonStr, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	// 超时时间：默认5秒
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	result, err := ioutil.ReadAll(resp.Body)
	return string(result), err
}
