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

package sendings

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/vogo/vmontnetsdk/cores"
	"github.com/vogo/vogo/vlog"
)

type SingleSendRequest struct {
	Mobile  string `json:"mobile" comment:"短信接收的手机号:只能填一个手机号"`
	Content string `json:"content" comment:"短信内容:最大支持1000个字(含签名),发送时请预留至少10个字的签名长度"`
	CustID  string `json:"custid,omitempty" comment:"用户自定义流水号:该条短信在业务系统内的ID,比如订单号或者短信发送记录的流水号,填写后发送状态返回值内将包含用户自定义流水号"`
	ExData  string `json:"exdata,omitempty" comment:"自定义扩展数据:额外提供的最大64个长度的ASCII字符串,填写后,状态报告返回时将会包含这部分数据"`
}

type SendSingleResponse struct {
	Result int    `json:"result" comment:"短信发送请求处理结果:0-成功,非0-失败"`
	Desc   string `json:"desc" comment:"应答结果描述,当result非0时,为错误描述"`
	MsgID  int64  `json:"msgid" comment:"平台流水号:非0,64位整型,对应Java和C#的long,不可用int解析,result非0时,msgid为0"`
	CustID string `json:"custid,omitempty" comment:"用户自定义流水号:默认与请求报文中的custid保持一致,若请求报文中没有custid参数或值为空,则返回由梦网生成的代表本批短信的唯一编号,result非0时,custid为空"`
}

// SendSingle 发送单条短信
func (s *SendingService) SendSingle(req *SingleSendRequest) (*SendSingleResponse, error) {
	// 验证手机号码
	if !cores.ValidateMobile(req.Mobile) {
		return nil, errors.New("invalid mobile number")
	}

	// 验证短信内容
	if req.Content == "" {
		return nil, errors.New("content cannot be empty")
	}

	// 构建请求参数
	params := make(map[string]string)

	// 添加鉴权参数
	authParams := cores.GenerateAuthParams(s.client.Config)
	for k, v := range authParams {
		params[k] = v
	}

	// 添加必要参数
	params["mobile"] = req.Mobile
	params["content"] = cores.EncodeContent(req.Content)

	// 添加可选参数
	if s.client.Config.SvrType != "" {
		params["svrtype"] = s.client.Config.SvrType
	}

	if s.client.Config.Exno != "" {
		params["exno"] = s.client.Config.Exno
	}

	if req.CustID != "" {
		params["custid"] = req.CustID
	}

	if req.ExData != "" {
		params["exdata"] = req.ExData
	}

	// 发送请求
	apiURL := s.client.Config.BaseURL + cores.PathSendSingle
	respBody, err := s.client.DoRequest("POST", apiURL, params)
	if err != nil {
		return nil, err
	}

	// 解析响应
	var commonResp *cores.CommonResponse
	commonResp, err = s.client.ParseCommonResponse(respBody)
	if err != nil {
		vlog.Errorf("[send_single] failed to parse common response: %v, respBody: %s", err, respBody)
		return nil, err
	}

	// 检查响应状态
	err = s.client.CheckResponse(commonResp)
	if err != nil {
		vlog.Errorf("[send_single] failed to check response: %v, respBody: %s", err, respBody)
		return nil, err
	}

	// 解析完整响应
	var resp SendSingleResponse
	err = json.Unmarshal(respBody, &resp)
	if err != nil {
		vlog.Errorf("[send_single] failed to parse response: %v, respBody: %s", err, respBody)
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}

	return &resp, nil
}
