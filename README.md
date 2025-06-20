# 梦网短信SDK (vmontnetsdk) - https://www.montnets.com

这是一个基于梦网云通讯平台API接口的Go语言SDK,提供简单易用的短信发送功能。

## 功能特性

- 支持多种鉴权方式（MD5加密、明文密码、APIKey）
- 支持单条短信发送、批量发送、个性化群发、模板发送等多种发送方式
- 支持状态报告查询、上行短信查询、余额查询等功能
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

	"github.com/vogo/vmontnetsdk/balances"
	"github.com/vogo/vmontnetsdk/cores"
	"github.com/vogo/vmontnetsdk/reports"
	"github.com/vogo/vmontnetsdk/sendings"
)

func main() {
	// 创建配置（支持多个API地址，第一个为主地址，其他为备份地址）
	config := cores.NewConfig(
		[]string{
			"https://api.montnets.com",      // 主地址
			"https://api2.montnets.com",     // 备份地址1
			"https://api3.montnets.com",     // 备份地址2
		},
		"J10003",                         // 用户ID
		"111111",                         // 密码
	)
	// 如果使用APIKey鉴权方式,可以使用以下方式创建配置
	// config := cores.NewConfigWithAPIKey(
	// 	[]string{"https://api.montnets.com", "https://api2.montnets.com"},
	// 	"abade5589e2798f82f006bbc36d269ce"
	// )
	
	// 也可以使用兼容旧版本的方式创建配置（仅支持单个地址）
	// config := cores.NewConfigWithSingleURL(
	// 	"https://api.montnets.com",
	// 	"J10003",
	// 	"111111",
	// )

	// 设置其他配置项
	config.UsePlainPwd = false  // 使用MD5加密密码（默认）
	config.SvrType = "0"       // 设置业务类型
	config.Exno = "0006"       // 设置扩展号
	config.Signature = "【梦网科技】" // 短信签名，如果设置，所有短信内容前会自动添加该签名前缀

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
	
	// 创建状态报告服务
	reportService := reports.NewReportService(client)

	// 查询状态报告
	rptResp, err := reportService.GetRpt(10) // 最多获取10条状态报告
	if err != nil {
		log.Printf("查询状态报告失败: %v", err)
	} else {
		fmt.Printf("成功获取状态报告,共 %d 条\n", len(rptResp.Rpts))
	}
	
	// 创建余额查询服务
	balanceService := balances.NewBalanceService(client)

	// 查询账户余额
	balanceResp, err := balanceService.GetBalance()
	if err != nil {
		log.Printf("查询账户余额失败: %v", err)
	} else {
		fmt.Printf("账户余额: %d 条\n", balanceResp.Balance)
	}
}
```

## 配置说明

### Config 配置项

| 配置项 | 类型 | 必填 | 说明 |
| --- | --- | --- | --- |
| BaseURLs | []string | 是 | API基础URL列表，第一个为主地址，其他为备份地址 |
| UserID | string | 是* | 用户ID（*使用API密钥时可不填） |
| Password | string | 是* | 密码（*使用API密钥时可不填） |
| APIKey | string | 是* | API密钥（*使用用户ID和密码时可不填） |
| UsePlainPwd | bool | 否 | 是否使用明文密码，默认为false（使用MD5加密） |
| FixedString | string | 否 | 固定字符串，默认为"00000000" |
| SvrType | string | 否 | 业务类型 |
| Exno | string | 否 | 扩展号 |
| Signature | string | 否 | 短信签名，如果设置，所有短信内容前会自动添加该签名前缀 |

## 鉴权方式

### 用户名密码鉴权

```go
config := cores.NewConfig(
    []string{"https://api.montnets.com"},
    "J00000",
    "123456",
)
```

### API密钥鉴权

```go
config := cores.NewConfigWithAPIKey(
    []string{"https://api.montnets.com"},
    "9SDK-EMY-0999-BDSKS",
)
```

## 短信签名

短信签名是短信内容的前缀，通常用于标识发送方，如【梦网科技】。SDK支持自动添加短信签名前缀，只需在配置中设置 `Signature` 字段：

```go
config.Signature = "【梦网科技】"
```

设置签名后，SDK会自动检查短信内容是否已包含签名，如果没有，会自动在内容前添加签名前缀。这适用于所有发送方式（单发、群发、个性化发送等）。

## 目录结构

- `cores/`: 核心功能包
  - `config.go`: 配置相关
  - `auth.go`: 鉴权相关
  - `client.go`: HTTP客户端
  - `consts.go`: 常量定义
- `sendings/`: 发送功能包
  - `service.go`: 发送服务定义
  - `single.go`: 单条发送实现
  - `batch.go`: 批量发送实现
  - `multi.go`: 个性化群发实现
  - `mixed.go`: 个性化群发实现（另一种方式）
  - `template.go`: 模板发送实现
- `reports/`: 状态报告功能包
  - `service.go`: 状态报告服务定义
  - `rpt.go`: 状态报告查询实现
- `mos/`: 上行短信功能包
  - `service.go`: 上行短信服务定义
  - `mo.go`: 上行短信查询实现
- `balances/`: 余额查询功能包
  - `service.go`: 余额查询服务定义
  - `balance.go`: 余额查询实现
- `example/`: 使用示例

## 参与贡献

我们非常欢迎您参与贡献，帮助我们改进这个SDK！

### 如何贡献

1. Fork 本仓库
2. 创建您的特性分支 (`git checkout -b feature/amazing-feature`)
3. 提交您的更改 (`git commit -m 'Add some amazing feature'`)
4. 推送到分支 (`git push origin feature/amazing-feature`)
5. 提交 Pull Request

### 问题反馈

如果您在使用过程中遇到任何问题，或者有任何建议和想法，欢迎通过以下方式反馈：

- 提交 [GitHub Issues](https://github.com/vogo/vmontnetsdk/issues)
- 在 Pull Request 中详细描述您的改进建议

### 贡献指南

- 请确保您的代码符合 Go 语言规范和项目的代码风格
- 添加必要的测试用例
- 更新相关文档
- 保持代码简洁、清晰，并添加必要的注释

## 许可证

[Apache License 2.0](LICENSE)