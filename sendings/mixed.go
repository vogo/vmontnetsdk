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
	"strings"

	"github.com/vogo/vmontnetsdk/cores"
)

type MixedSendRequest struct {
	Mobiles  string   `json:"mobile" comment:"短信接收的手机号:多个手机号请用英文逗号分隔,最大500个号码"`
	Contents []string `json:"content" comment:"短信内容:多个内容以英文逗号分隔,信息内容与手机号顺序一一对应,如果信息内容数量与手机号个数不一致将返回错误"`
	CustID   string   `json:"custid,omitempty" comment:"用户自定义流水号:该条短信在业务系统内的ID,比如订单号或者短信发送记录的流水号,填写后发送状态返回值内将包含用户自定义流水号"`
	ExData   string   `json:"exdata,omitempty" comment:"自定义扩展数据:额外提供的最大64个长度的ASCII字符串,填写后,状态报告返回时将会包含这部分数据"`
}

type SendMixedResponse struct {
	Result int    `json:"result" comment:"个性化群发请求处理结果:0-成功,非0-失败"`
	Desc   string `json:"desc" comment:"应答结果描述,当result非0时,为错误描述"`
	MsgID  int64  `json:"msgid" comment:"平台流水号:非0,64位整型,对应Java和C#的long,不可用int解析,result非0时,msgid为0"`
	CustID string `json:"custid,omitempty" comment:"用户自定义流水号:默认与请求报文中的custid保持一致,若请求报文中没有custid参数或值为空,则返回由梦网生成的代表本批短信的唯一编号,result非0时,custid为空"`
}

// SendMixed 发送个性化短信（不同手机号对应不同内容）
func (s *SendingService) SendMixed(req *MixedSendRequest) (*SendMixedResponse, error) {
	// 验证手机号码
	if req.Mobiles == "" {
		return nil, errors.New("mobiles cannot be empty")
	}

	// 验证短信内容
	if len(req.Contents) == 0 {
		return nil, errors.New("contents cannot be empty")
	}

	// 验证手机号和内容数量是否匹配
	mobiles := strings.Split(req.Mobiles, ",")
	if len(mobiles) != len(req.Contents) {
		return nil, errors.New("number of mobiles and contents must match")
	}

	// 构建请求参数
	params := make(map[string]string)

	// 添加鉴权参数
	authParams := cores.GenerateAuthParams(s.client.Config)
	for k, v := range authParams {
		params[k] = v
	}

	// 添加必要参数
	params["mobile"] = req.Mobiles

	// 对内容进行编码并拼接
	encodedContents := make([]string, len(req.Contents))
	for i, content := range req.Contents {
		// 应用签名前缀
		content = s.client.Config.ApplySignature(content)
		encodedContents[i] = cores.EncodeContent(content)
	}
	params["content"] = strings.Join(encodedContents, ",")

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
	respBody, err := s.client.DoRequest("POST", cores.PathSendMixed, params)
	if err != nil {
		return nil, err
	}

	// 解析响应
	var resp SendMixedResponse
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
