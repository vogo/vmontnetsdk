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

package mos

import (
	"encoding/json"
	"fmt"

	"github.com/vogo/vmontnetsdk/cores"
)

// MoItem 上行短信项
type MoItem struct {
	Mobile    string `json:"mobile" comment:"手机号码"`
	SpNumber  string `json:"spnumber" comment:"完整的通道号"`
	Content   string `json:"content" comment:"短信内容"`
	ExNo      string `json:"exno" comment:"下行时填写的扩展号"`
	RecvTime  string `json:"recvtime" comment:"上行返回的时间,格式:YYYY-MM-DD HH:MM:SS"`
	NotifyURL string `json:"notifyurl" comment:"通知URL"`
}

// GetMoResponse 上行短信查询响应
type GetMoResponse struct {
	Result int      `json:"result" comment:"请求处理结果:0-成功,非0-失败"`
	Desc   string   `json:"desc" comment:"结果描述,当result非0时为错误描述"`
	Mos    []MoItem `json:"mos" comment:"上行短信列表,result非0时为空"`
}

// GetMo 查询上行短信
func (s *MoService) GetMo(maxCount int) (*GetMoResponse, error) {
	// 构建请求参数
	params := make(map[string]string)

	// 添加鉴权参数
	authParams := cores.GenerateAuthParams(s.client.Config)
	for k, v := range authParams {
		params[k] = v
	}

	// 添加可选参数
	if maxCount > 0 {
		params["retsize"] = fmt.Sprintf("%d", maxCount)
	}

	// 发送请求
	respBody, err := s.client.DoRequest("POST", cores.PathGetMo, params)
	if err != nil {
		return nil, err
	}

	// 解析响应
	var resp GetMoResponse
	err = json.Unmarshal(respBody, &resp)
	if err != nil {
		return nil, err
	}

	// 检查响应状态
	if resp.Result != 0 {
		desc, _ := cores.DecodeContent(resp.Desc)
		return &resp, fmt.Errorf("API error: code=%d, desc=%s", resp.Result, desc)
	}

	// 解码短信内容
	for i := range resp.Mos {
		resp.Mos[i].Content, _ = cores.DecodeContent(resp.Mos[i].Content)
	}

	return &resp, nil
}
