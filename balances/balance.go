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

package balances

import (
	"encoding/json"
	"fmt"

	"github.com/vogo/vmontnetsdk/cores"
)

// GetBalanceResponse 余额查询响应
type GetBalanceResponse struct {
	Result  int    `json:"result" comment:"请求处理结果:0-成功,非0-失败"`
	Desc    string `json:"desc" comment:"结果描述,当result非0时为错误描述"`
	Balance int    `json:"balance" comment:"短信余额条数"`
}

// GetBalance 查询账户余额
func (s *BalanceService) GetBalance() (*GetBalanceResponse, error) {
	// 构建请求参数
	params := make(map[string]string)

	// 添加鉴权参数
	authParams := cores.GenerateAuthParams(s.Client.Config)
	for k, v := range authParams {
		params[k] = v
	}

	// 发送请求
	respBody, err := s.Client.DoRequest("POST", cores.PathGetBalance, params)
	if err != nil {
		return nil, err
	}

	// 解析响应
	var resp GetBalanceResponse
	err = json.Unmarshal(respBody, &resp)
	if err != nil {
		return nil, err
	}

	// 检查响应状态
	if resp.Result != 0 {
		desc, _ := cores.DecodeContent(resp.Desc)
		return &resp, fmt.Errorf("API error: code=%d, desc=%s", resp.Result, desc)
	}

	return &resp, nil
}
