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

package cores

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// ResponseFormat 响应格式
type ResponseFormat string

const (
	// ResponseFormatJSON JSON格式
	ResponseFormatJSON ResponseFormat = "json"
	// ResponseFormatXML XML格式
	ResponseFormatXML ResponseFormat = "xml"
	// ResponseFormatURLEncoded URL编码格式
	ResponseFormatURLEncoded ResponseFormat = "urlencode"
)

// Client 梦网短信客户端
type Client struct {
	Config         *Config
	httpClient     *http.Client
	ResponseFormat ResponseFormat
}

// CommonResponse 通用响应结构
type CommonResponse struct {
	Result int    `json:"result"`
	Desc   string `json:"desc"`
}

// NewClient 创建新的客户端
func NewClient(config *Config) *Client {
	return &Client{
		Config: config,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		ResponseFormat: ResponseFormatJSON,
	}
}

// DoRequest 执行HTTP请求
func (c *Client) DoRequest(method, urlStr string, params map[string]string) ([]byte, error) {
	var req *http.Request
	var err error

	switch method {
	case http.MethodGet:
		// 构建GET请求URL
		queryValues := url.Values{}
		for k, v := range params {
			queryValues.Set(k, v)
		}
		urlStr = fmt.Sprintf("%s?%s", urlStr, queryValues.Encode())
		req, err = http.NewRequest(method, urlStr, nil)
		if err != nil {
			return nil, err
		}
	case http.MethodPost:
		// 根据响应格式构建POST请求体
		var body io.Reader

		switch c.ResponseFormat {
		case ResponseFormatJSON:
			jsonData, err := json.Marshal(params)
			if err != nil {
				return nil, err
			}
			body = bytes.NewBuffer(jsonData)
			req, err = http.NewRequest(method, urlStr, body)
			if err != nil {
				return nil, err
			}
			req.Header.Set("Content-Type", "application/json")
		case ResponseFormatURLEncoded:
			formValues := url.Values{}
			for k, v := range params {
				formValues.Set(k, v)
			}
			body = strings.NewReader(formValues.Encode())
			req, err = http.NewRequest(method, urlStr, body)
			if err != nil {
				return nil, err
			}
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		default:
			return nil, errors.New("unsupported response format")
		}
	default:
		return nil, errors.New("unsupported HTTP method")
	}

	// 执行请求
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应内容
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}

// ParseCommonResponse 解析通用响应
func (c *Client) ParseCommonResponse(respBody []byte) (*CommonResponse, error) {
	var resp CommonResponse
	err := json.Unmarshal(respBody, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CheckResponse 检查响应是否成功
func (c *Client) CheckResponse(resp *CommonResponse) error {
	if resp.Result != 0 {
		desc, _ := DecodeContent(resp.Desc)
		return fmt.Errorf("API error: code=%d, desc=%s", resp.Result, desc)
	}
	return nil
}
