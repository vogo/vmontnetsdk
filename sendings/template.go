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
	"net/url"
	"strings"

	"github.com/vogo/vmontnetsdk/cores"
)

type TemplateSendRequest struct {
	TmplID  string            `json:"tmplid" comment:"短信模版编号:长度最大20位字符"`
	Mobiles string            `json:"mobile" comment:"短信接收的手机号:多个手机号请用英文逗号分隔,最大1000个号码"`
	Params  map[string]string `json:"params" comment:"变量名和变量值:一个模板变量名对应一个变量值,多个变量使用key=value的方式进行拼接"`
	CustID  string            `json:"custid,omitempty" comment:"用户自定义流水号:该条短信在业务系统内的ID,比如订单号或者短信发送记录的流水号,填写后发送状态返回值内将包含用户自定义流水号"`
	ExData  string            `json:"exdata,omitempty" comment:"自定义扩展数据:额外提供的最大64个长度的ASCII字符串,填写后,状态报告返回时将会包含这部分数据"`
}

type SendTemplateResponse struct {
	Result int    `json:"result" comment:"短信发送请求处理结果:0-成功,非0-失败"`
	Desc   string `json:"desc" comment:"应答结果描述,当result非0时,为错误描述"`
	MsgID  int64  `json:"msgid" comment:"平台流水号:非0,64位整型,对应Java和C#的long,不可用int解析,result非0时,msgid为0"`
	CustID string `json:"custid,omitempty" comment:"用户自定义流水号:默认与请求报文中的custid保持一致,若请求报文中没有custid参数或值为空,则返回由梦网生成的代表本批短信的唯一编号,result非0时,custid为空"`
}

// SendTemplate 发送模板短信
func (s *SendingService) SendTemplate(req *TemplateSendRequest) (*SendTemplateResponse, error) {
	// 验证模板ID
	if req.TmplID == "" {
		return nil, errors.New("template ID cannot be empty")
	}

	// 验证手机号码
	if req.Mobiles == "" {
		return nil, errors.New("mobiles cannot be empty")
	}

	// 验证模板参数
	if len(req.Params) == 0 {
		return nil, errors.New("template params cannot be empty")
	}

	// 构建请求参数
	params := make(map[string]string)

	// 添加鉴权参数
	authParams := cores.GenerateAuthParams(s.client.Config)
	for k, v := range authParams {
		params[k] = v
	}

	// 添加必要参数
	params["tmplid"] = req.TmplID
	params["mobile"] = req.Mobiles

	// 处理模板参数
	// 第一次编码:对每个变量值进行URL编码
	var contentParts []string
	for key, value := range req.Params {
		encodedValue := url.QueryEscape(value)
		contentParts = append(contentParts, fmt.Sprintf("%s=%s", key, encodedValue))
	}

	// 拼接参数
	content := strings.Join(contentParts, "&")

	// 第二次编码:对整体内容进行URL编码
	params["content"] = url.QueryEscape(content)

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
	respBody, err := s.client.DoRequest("POST", cores.PathSendTemplate, params)
	if err != nil {
		return nil, err
	}

	// 解析响应
	var resp SendTemplateResponse
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
