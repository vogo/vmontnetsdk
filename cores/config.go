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
	// API基础URL列表，第一个为主地址，其他为备份地址
	BaseURLs []string
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
// baseURLs 参数可以传入多个URL，第一个为主地址，其他为备份地址
func NewConfig(baseURLs []string, userID, password string) *Config {
	return &Config{
		BaseURLs:    baseURLs,
		UserID:      userID,
		Password:    password,
		FixedString: "00000000",
	}
}

// NewConfigWithSingleURL 创建新的配置（使用单个URL、用户ID和密码）
// 为了兼容旧版本API
func NewConfigWithSingleURL(baseURL, userID, password string) *Config {
	return NewConfig([]string{baseURL}, userID, password)
}

// NewConfigWithAPIKey 创建新的配置（使用API密钥）
// baseURLs 参数可以传入多个URL，第一个为主地址，其他为备份地址
func NewConfigWithAPIKey(baseURLs []string, apiKey string) *Config {
	return &Config{
		BaseURLs:    baseURLs,
		APIKey:      apiKey,
		FixedString: "00000000",
	}
}

// NewConfigWithAPIKeyAndSingleURL 创建新的配置（使用单个URL和API密钥）
// 为了兼容旧版本API
func NewConfigWithAPIKeyAndSingleURL(baseURL, apiKey string) *Config {
	return NewConfigWithAPIKey([]string{baseURL}, apiKey)
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

// GetBaseURL 获取当前可用的BaseURL
// 默认返回第一个URL（主地址）
func (c *Config) GetBaseURL() string {
	if len(c.BaseURLs) == 0 {
		return ""
	}
	return c.BaseURLs[0]
}
