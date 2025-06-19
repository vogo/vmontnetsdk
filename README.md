# 梦网短信SDK (vmontnetsdk) - https://www.montnets.com

这是一个基于梦网云通讯平台API接口的Go语言SDK,提供简单易用的短信发送功能。

## 功能特性

- 支持多种鉴权方式（MD5加密、明文密码、APIKey）
- 支持单条短信发送
- 支持JSON、URL编码格式的请求和响应
- 提供完整的错误处理和参数验证

## 安装

```bash
go get github.com/vogo/vmontnetsdk
```

## 使用示例

### 单条短信发送

```go
package main

import (
	"fmt"
	"log"

	"github.com/vogo/vmontnetsdk/cores"
	"github.com/vogo/vmontnetsdk/sendings"
)

func main() {
	// 创建配置
	config := cores.NewConfig("ip", 80, "J10003", "111111")
	// 如果使用APIKey鉴权方式,可以使用以下方式创建配置
	// config := cores.NewConfigWithAPIKey("ip", 80, "abade5589e2798f82f006bbc36d269ce")

	// 设置其他配置项
	config.UseHTTPS = true      // 使用HTTPS
	config.UsePlainPwd = false  // 使用MD5加密密码
	config.SvrType = "SMS001"   // 设置业务类型
	config.ExNo = "0006"        // 设置扩展号

	// 创建客户端
	client := cores.NewClient(config)
	client.ResponseFormat = cores.ResponseFormatJSON // 设置响应格式为JSON

	// 创建发送服务
	sendingService := sendings.NewSendingService(client)

	// 发送单条短信
	resp, err := sendingService.SendSingle(&sendings.SingleSendRequest{
		Mobile:  "13800138000",
		Content: "验证码:6666,打死都不要告诉别人哦！",
		CustID:  "b3d0a2783d31b21b8573",
		ExData:  "exdata000002",
	})

	if err != nil {
		log.Fatalf("发送短信失败: %v", err)
	}

	fmt.Printf("发送短信成功,平台流水号: %d, 自定义流水号: %s\n", resp.MsgID, resp.CustID)
}
```

## 配置说明

### Config 配置项

| 配置项 | 类型 | 说明 |
| --- | --- | --- |
| ServerIP | string | 服务器地址 |
| ServerPort | int | 服务器端口 |
| UseHTTPS | bool | 是否使用HTTPS |
| UserID | string | 用户账号 |
| Password | string | 用户密码 |
| APIKey | string | APIKey (如果使用APIKey鉴权方式,则UserID和Password可不填) |
| UsePlainPwd | bool | 是否使用明文密码 |
| FixedString | string | 固定字符串,用于MD5加密 |
| SvrType | string | 业务类型 |
| ExNo | string | 扩展号 |

## 鉴权方式

### 1. 账号+密码的MD5加密鉴权

```go
config := cores.NewConfig("ip", 80, "J10003", "111111")
config.UsePlainPwd = false  // 使用MD5加密密码
```

### 2. 账号+密码的明文鉴权

```go
config := cores.NewConfig("ip", 80, "J10003", "111111")
config.UsePlainPwd = true  // 使用明文密码
```

### 3. APIKey鉴权

```go
config := cores.NewConfigWithAPIKey("ip", 80, "abade5589e2798f82f006bbc36d269ce")
```

## 目录结构

- `cores/`: 核心功能包
  - `config.go`: 配置相关
  - `auth.go`: 鉴权相关
  - `client.go`: HTTP客户端
  - `response.go`: 响应结构定义
- `sendings/`: 发送功能包
  - `service.go`: 发送服务定义
  - `single.go`: 单条发送实现
- `example/`: 使用示例

## 待实现功能

- 相同内容群发接口 (send_batch)
- 个性化群发接口 (send_multi)
- 个性化群发接口 (send_mixed)
- 模板发送接口 (send_template)
- 获取上行接口 (get_mo)
- 获取状态报告接口 (get_rpt)
- 查询余额接口 (get_balance)

## 许可证

[Apache License 2.0](LICENSE)