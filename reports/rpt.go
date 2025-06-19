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

package reports

import (
	"encoding/json"
	"fmt"

	"github.com/vogo/vmontnetsdk/cores"
)

// RptItem 状态报告项
type RptItem struct {
	Mobile    string `json:"mobile" comment:"手机号码"`
	MsgID     int64  `json:"msgid" comment:"平台流水号:对应下行请求返回结果中的msgid"`
	Status    int    `json:"status" comment:"接收状态:0-成功,非0-失败"`
	ReceTime  string `json:"recvtime" comment:"状态报告返回时间,格式:YYYY-MM-DD HH:MM:SS"`
	NotifyURL string `json:"notifyurl" comment:"通知URL"`
	CustID    string `json:"custid,omitempty" comment:"用户自定义流水号:对应下行请求时填写的custid"`
	ExData    string `json:"exdata,omitempty" comment:"下行时填写的扩展数据"`
}

// GetRptResponse 状态报告查询响应
type GetRptResponse struct {
	Result int       `json:"result" comment:"请求处理结果:0-成功,非0-失败"`
	Desc   string    `json:"desc" comment:"结果描述,当result非0时为错误描述"`
	Rpts   []RptItem `json:"rpts" comment:"状态报告列表,result非0时为空"`
}

// GetRpt 查询状态报告
func (s *ReportService) GetRpt(maxCount int) (*GetRptResponse, error) {
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
	respBody, err := s.client.DoRequest("POST", cores.PathGetReport, params)
	if err != nil {
		return nil, err
	}

	// 解析响应
	var resp GetRptResponse
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
