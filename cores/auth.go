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
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

// AuthType 鉴权类型
type AuthType int

const (
	// AuthTypeMD5 MD5加密鉴权
	AuthTypeMD5 AuthType = 0
	// AuthTypePlain 明文鉴权
	AuthTypePlain AuthType = 1
	// AuthTypeAPIKey APIKey鉴权
	AuthTypeAPIKey AuthType = 2
)

// GetTimestamp 获取时间戳,格式为MMDDHHMMSS
func GetTimestamp() string {
	now := time.Now()
	return fmt.Sprintf("%02d%02d%02d%02d%02d", now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
}

// GenerateMD5Password 生成MD5加密的密码
// 将userid值大写、固定字符串00000000、明文pwd、timestamp依次拼接成字符串后,再进行MD5加密
func GenerateMD5Password(userID, password, fixedString, timestamp string) string {
	// 将userID转为大写
	userID = strings.ToUpper(userID)
	// 拼接字符串
	str := fmt.Sprintf("%s%s%s%s", userID, fixedString, password, timestamp)
	// MD5加密
	hasher := md5.New()
	hasher.Write([]byte(str))
	return hex.EncodeToString(hasher.Sum(nil))
}

// GenerateAuthParams 根据配置生成鉴权参数
func GenerateAuthParams(config *Config) map[string]string {
	params := make(map[string]string)

	// 判断使用哪种鉴权方式
	if config.APIKey != "" {
		// APIKey鉴权
		params["apikey"] = config.APIKey
	} else {
		// 用户名密码鉴权
		params["userid"] = config.UserID

		if config.UsePlainPwd {
			// 明文密码
			params["pwd"] = config.Password
		} else {
			// MD5加密密码
			timestamp := GetTimestamp()
			params["timestamp"] = timestamp
			params["pwd"] = GenerateMD5Password(config.UserID, config.Password, config.FixedString, timestamp)
		}
	}

	return params
}
