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

package main

import (
	"fmt"
	"log"

	"github.com/vogo/vmontnetsdk/balances"
	"github.com/vogo/vmontnetsdk/cores"
	"github.com/vogo/vmontnetsdk/reports"
	"github.com/vogo/vmontnetsdk/sendings"
	"github.com/vogo/vogo/vos"
)

func main() {
	// 创建配置
	config := cores.NewConfig(
		vos.EnsureEnvString("MONTNETS_API_URL"),
		vos.EnsureEnvString("MONTNETS_ACCOUNT"),
		vos.EnsureEnvString("MONTNETS_PASSWORD"),
	)
	// 设置使用明文密码
	config.UsePlainPwd = false
	// 设置业务类型
	config.SvrType = "0"
	// 设置扩展号
	config.Exno = ""

	// 创建客户端
	client := cores.NewClient(config)
	// 设置响应格式为JSON
	client.ResponseFormat = cores.ResponseFormatJSON

	// 创建发送服务
	sendingService := sendings.NewSendingService(client)

	// 发送单条短信
	request := &sendings.SingleSendRequest{
		Mobile:  vos.EnsureEnvString("MONTNETS_TEST_MOBILE"),
		Content: vos.EnsureEnvString("MONTNETS_TEST_CONTENT"),
	}

	resp, err := sendingService.SendSingle(request)
	if err != nil {
		log.Fatalf("发送短信失败: %v", err)
	}

	fmt.Printf("发送成功,消息ID: %d\n", resp.MsgID)

	// 创建状态报告服务
	reportService := reports.NewReportService(client)

	// 查询状态报告
	rptResp, err := reportService.GetRpt(10) // 最多获取10条状态报告
	if err != nil {
		log.Printf("查询状态报告失败: %v", err)
	} else {
		fmt.Printf("成功获取状态报告,共 %d 条\n", len(rptResp.Rpts))
		for i, rpt := range rptResp.Rpts {
			fmt.Printf("报告 #%d: 手机号=%s, 状态=%d, 消息ID=%d\n",
				i+1, rpt.Mobile, rpt.Status, rpt.MsgID)
		}
	}

	// 创建余额查询服务
	balanceService := balances.NewBalanceService(client)

	// 查询账户余额
	balanceResp, err := balanceService.GetBalance()
	if err != nil {
		log.Printf("查询账户余额失败: %v", err)
		return
	} else {
		fmt.Printf("账户余额: %d 条\n", balanceResp.Balance)
	}
}
