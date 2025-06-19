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
	"net/url"
	"regexp"
	"strings"
)

// Config 梦网短信配置
type Config struct {
	// API基础URL
	BaseURL string
	// 用户ID
	UserID string
	// 密码
	Password string
	// API密钥
	APIKey string
	// 是否使用明文密码
	UsePlainPwd bool
	// 固定字符串
	FixedString string
	// 业务类型
	SvrType string
	// 扩展号
	Exno string
}

// NewConfig 创建新的配置（使用用户ID和密码）
func NewConfig(baseURL, userID, password string) *Config {
	return &Config{
		BaseURL:     baseURL,
		UserID:      userID,
		Password:    password,
		FixedString: "00000000",
	}
}

// NewConfigWithAPIKey 创建新的配置（使用API密钥）
func NewConfigWithAPIKey(baseURL, apiKey string) *Config {
	return &Config{
		BaseURL:     baseURL,
		APIKey:      apiKey,
		FixedString: "00000000",
	}
}

// EncodeContent URL编码内容
func EncodeContent(content string) string {
	return url.QueryEscape(content)
}

// DecodeContent URL解码内容
func DecodeContent(content string) (string, error) {
	return url.QueryUnescape(content)
}

// ValidateMobile 验证手机号码
func ValidateMobile(mobile string) bool {
	// 简单验证手机号码格式（中国大陆手机号）
	pattern := `^1[3-9]\d{9}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(mobile)
}

// ValidateMobiles 验证多个手机号码
func ValidateMobiles(mobiles string) bool {
	// 分割手机号码
	mobileList := strings.Split(mobiles, ",")
	for _, mobile := range mobileList {
		if !ValidateMobile(mobile) {
			return false
		}
	}
	return true
}
