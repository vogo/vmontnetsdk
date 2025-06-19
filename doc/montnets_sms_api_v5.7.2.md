梦网平台 API 接口说明(V5.7.2)

https://www.montnets.com/
深圳市梦网科技发展有限公司二○一七年九月


目录

1. [接口简介 4]
2. [接口说明 4]
    1. [单条发送接口 send_single 4]
    2. [相同内容群发接口 send_batch 9]
    3. [个性化群发接口 send_multi 13]
    4. [个性化群发接口 send_mixed 19]
    5. [模板发送接口 send_template 24]
    6. [获取上行接口 get_mo 29]
    7. [获取状态报告接口 get_rpt 34]
    8. [查询余额接口 get_balance 40]
    9. [推送上行接口 43]
    10. [推送状态报告接口 46]
3. [规则说明 50]
    1. [鉴权规则 50]
    2. [手机号码规则 51]
    3. [匹配状态报告规则 52]
    4. [内容加密规则 53]
4. [错误代码5. [注意事项 57]

# 接口简介

梦网云通讯平台 API 接口是梦网凭借多年的开发经验,专为用户提供的简单易用的 API调用服务,皆在为第三方开发者在应用内快速、高效、低成本集成梦网云通讯平台的各项业务提供了一站式服务。API 采用了当前流行的 REST 风格设计和实现,简洁易懂,支持多种主流开发语言的调用,支持 JSON、XML、x-www-form-urlencoded 三种数据传递方式,用户可以根据实际的开发需求选择合适的数据传递方式,兼容了大部分用户对 API 接口的 需求。另外为了方便用户进行对接,我们提供了将 API 封装好的 SDK 开发包,针对通过 SDK开发包接入的用户,能方便其快速使用平台提供的功能,SDK 的详细使用方法可查看《梦网平台 SDK 接口说明》文档。

接口应答、推送请求实际包含的字段可能超出文档中列举的部分,建议用户在解析时忽略这些字段,请勿当作错误处理。

# 接口说明

## 1. 单条发送接口 send_single

### 请求URL

<https://ip:port/sms/v2/std/send_single> <http://ip:port/sms/v2/std/send_single>

### 请求方式

POST/GET

### 请求参数说明

| 参数 | 类型 | 是否必须 | 描述 | 示例 |
| --- | --- | --- | --- | --- |
| userid | string | 是   | 用户账号:长度最大 6 个字符,统一大写,如提交参数中包含 apikey,则可以不用填写该参数及 pwd,两种鉴权方式中只<br><br>能选择一种方式来进行鉴权 | 示例:J10003 |
| pwd | string | 是   | 用户密码:定长小写 32 位字符,如提交参数中包含 apikey,则可以不用填写该参数及 userid,两种鉴权方式中只能选择一种方式来进行鉴权。密码规则详见“[3.1 鉴权规则]” | 示例:<br><br>密码明文模式:111111密码加密模式:<br><br>账号:J10003密码:111111<br><br>固定字符串:00000000时间戳:0803192020<br><br>MD5 加密之前的对应字符串: J1000300000000111111080 3192020<br><br>MD5 加密之后的密码字符串: 26dad7f364507df18f3841c<br><br>c9c4ff94d |
| mobile | string | 是   | 短信接收的手机号:只能填一个手机号。号码规则详见“[3.2]<br><br>手机号码规则” | 示例:138xxxxxxxx |
| content | string | 是   | 短信内容:最大支持 1000 个字(含签名),发送时请预留至少 10 个字的签名长度,一个字母或一个汉字都视为一个字。本字段支持加密传输,若不使用加密功能编码方式: UrlEncode(" 验 证 码 : 6666,打死都不要告诉<br><br>别人哦！","UTF-8")<br><br>若使用加密功能详见“[3.4 内]容加密规则” | 示例:<br><br>短信内容:“验证码:6666,打死都不要告诉别人哦！”短信内容进行urlencode 编码后:<br><br>%e9%aa%8c%e8%af%81<br><br>%e7%a0%81%ef%bc%9a6 666%ef%bc%8c%e6%89%93%e6%ad%bb%e9%83% bd%e4%b8%8d%e8%a6%81%e5%91%8a%e8%af%8<br><br>9%e5%88%ab%e4%ba%ba<br><br>%e5%93%a6%ef%bc%81 |
| timestamp | string | 否   | 时间戳:24 小时制格式:<br><br>MMDDHHMMSS,即月日时<br><br>分秒,定长 10 位,月、日、时、分、秒每段不足 2 位时左补 0,密码选择 MD5 加密方式时必填该参数,密码选择明文方式<br><br>时则不用填写 | 示例:0803192020 |
| svrtype | string | 否   | 业务类型:最大可支持 32 个<br><br>长度的英文字母、数字组合的字符串 | 示例:SMS001 |
| exno | string | 否   | 扩展号:长度不能超过 6 位,注意通道号+扩展号的总长度不能超过20 位,若超出则exno无效,如不需要扩展号则不用<br><br>提交此字段或填空 | 示例:0006 |
| custid | string | 否   | 用户自定义流水号:该条短信在您业务系统内的 ID,比如订单号或者短信发送记录的流水号。填写后发送状态返回值内将包含用户自定义流水号。最大可支持 64 位的 ASCII 字符串:字母、数字、下划线、减号,如不需要则不用提交此<br><br>字段或填空 | 示例:<br><br>b3d0a2783d31b21b8573 |
| exdata | string | 否   | 自定义扩展数据:额外提供的最大 64 个长度的 ASCII 字符串:字母、数字、下划线、减号,作为自定义扩展数据,填写后,状态报告返回时将会包含这部分数据,如不需要则不<br><br>用提交此字段或填空 |     |


### 返回参数说明

| 参数 | 类型 | 描述 | 示例 |
| --- | --- | --- | --- |
| result | int | 短信发送请求处理结果:<br><br>0:成功<br><br>非 0:失败,详见 [4 错误代码表] | 示例:0 |
| desc | string | 应答结果描述,当 result 非 0 时,为错误描述<br><br>编码方式:urlencode(UTF-8) | 示例:<br><br>“鉴权失败”<br><br>urlencode(UTF-8)编码:<br><br>%e9%89%b4%e6%9d%83<br><br>%e5%a4%b1%e8%b4%a5 |
| msgid | long (64 位) | 平台流水号:非 0,64 位整型,对应Java 和 C#的 long,不可用 int 解析。<br><br>result 非 0 时,msgid 为 0 | 示例:9223372036854775808<br><br>注意:msgid 允许出现负数 |
| custid | string | 用户自定义流水号:默认与请求报文中的 custid 保持一致,若请求报文中没有 custid 参数或值为空,则返回由梦网生成的代表本批短信的唯一编号<br><br>result 非 0 时,custid 为空 | 示例:<br><br>b3d0a2783d31b21b8573 |


### Post 请求示例

urlencode
```
userid=J10003&pwd=26dad7f364507df18f3841cc9c4ff94d&mobile=138xxxxxxxx &content=%e9%aa%8c%e8%af%81%e7%a0%81%ef%bc%9a6666%ef%bc%8c%e6%89%93%e 6%ad%bb%e9%83%bd%e4%b8%8d%e8%a6%81%e5%91%8a%e8%af%89%e5%88%ab%e4%ba%b a%e5%93%a6%ef%bc%81&timestamp=0803192020&svrtype=SMS001&exno=0006&cus
tid=b3d0a2783d31b21b8573&exdata=exdata000002
```

JSON
```json
{"userid":"J10003","pwd":"26dad7f364507df18f3841cc9c4ff94d","mobile": "138xxxxxxxx","content":"%e9%aa%8c%e8%af%81%e7%a0%81%ef%bc%9a6666%ef% bc%8c%e6%89%93%e6%ad%bb%e9%83%bd%e4%b8%8d%e8%a6%81%e5%91%8a%e8%af%89% e5%88%ab%e4%ba%ba%e5%93%a6%ef%bc%81","timestamp":"0803192020","svrtyp e":"SMS001","exno":"0006","custid":"b3d0a2783d31b21b8573","exdata":"ex data000002"}
```

XML
```xml
<?xml version=1.0 encoding=utf-8?>
<mtreq>
<userid>J10003</userid>
<pwd>26dad7f364507df18f3841cc9c4ff94d </pwd>
<mobile>138xxxxxxxx</mobile>
<content>%e9%aa%8c%e8%af%81%e7%a0%81%ef%bc%9a6666%ef%bc%8c%e6%89%93%e6%ad%bb%e9%83%bd%e4%b8%8d%e8%a6%81%e5%91%8a%e8%af%89%e5%88%ab%e4% ba%ba%e5%93%a6%ef%bc%81</content>
<timestamp>0803192020</timestamp>
<svrtype>SMS001</svrtype>
<exno>0006</exno>
<custid>b3d0a2783d31b21b8573</custid>
<exdata>exdata000002</exdata>
</mtreq>
```

### GET 请求示例

URLENCODE

用户名密码鉴权方式
```
userid=J10003&pwd=26dad7f364507df18f3841cc9c4ff94d&mobile=138xxxxxxxx &content=%e9%aa%8c%e8%af%81%e7%a0%81%ef%bc%9a6666%ef%bc%8c%e6%89%93%e 6%ad%bb%e9%83%bd%e4%b8%8d%e8%a6%81%e5%91%8a%e8%af%89%e5%88%ab%e4%ba%b a%e5%93%a6%ef%bc%81&timestamp=0803192020&svrtype=SMS001&exno=0006&cus tid=b3d0a2783d31b21b8573&exdata=exdata000002
```
apikey 鉴权方式
```
apikey=abade5589e2798f82f006bbc36d269ce&mobile=138xxxxxxxx&content=%e 9%aa%8c%e8%af%81%e7%a0%81%ef%bc%9a6666%ef%bc%8c%e6%89%93%e6%ad%bb%e9%83%bd%e4%b8%8d%e8%a6%81%e5%91%8a%e8%af%89%e5%88%ab%e4%ba%ba%e5%93%a6% ef%bc%81&svrtype=SMS001&exno=0006&custid=b3d0a2783d31b21b8573&exdata= exdata000002
```
### 发送成功返回示例

JSON (urlencode 与 JSON 请求都以 JSON 返回数据)
```json
{
"result":0, "desc":%e6%88%90%e5%8a%9f, "msgid":9223372036854775808,
"custid":"b3d0a2783d31b21b8573"

}
```

XML
```xml
<?xml version=1.0 encoding=utf-8?>
<mtrsp>
<result>0</result>
<desc>%e6%88%90%e5%8a%9f</desc>
<msgid>9223372036854775808</msgid>
<custid>b3d0a2783d31b21b8573</custid>
</mtrsp>
```

### 发送失败返回示例

JSON (urlencode 与 JSON 请求都以 JSON 返回数据)
```json
{
"result":-100999, "desc":%e7%b3%bb%e7%bb%9f%e5%86%85%e9%83%a8%e9%94%99%e8%af%af, "msgid":0,
"custid":""

}
```

XML
```xml
<?xml version=1.0 encoding=utf-8?>
<mtrsp>
<result>-100999</result>
<desc>

%e7%b3%bb%e7%bb%9f%e5%86%85%e9%83%a8%e9%94%99%e8%af%af

</desc>
<msgid>0</msgid>
<custid></custid>
</mtrsp>
```

## 2. 相同内容群发接口 send_batch

### 请求URL

<https://ip:port/sms/v2/std/send_batch> <http://ip:port/sms/v2/std/send_batch>

### 请求方式

POST/GET

### 请求参数说明

| 参数 | 类型 | 是否必须 | 描述 | 示例 |
| --- | --- | --- | --- | --- |
| userid | string | 是   | 用户账号:长度最大 6 个字符,统一大写,如提交参数中包含 apikey,则可以不用填写该参数及 pwd,两种鉴权方式中只<br><br>能选择一种方式来进行鉴权 | 示例:J10003 |
| pwd | string | 是   | 用户密码:定长小写 32 位字符,如提交参数中包含 apikey,则可以不用填写该参数及 userid,两种鉴权方式中只能选择一种方式来进行鉴权。密码规则详见“[3.1 鉴权规则]” | 示例:<br><br>密码明文模式:111111密码加密模式:<br><br>账号:J10003密码:111111<br><br>固定字符串:00000000时间戳:0803192020<br><br>MD5 加密之前的对应字符<br><br>串:J1000300000000111111080 3192020<br><br>MD5 加密之后的密码字符<br><br>串: 26dad7f364507df18f3841c c9c4ff94d |
| mobile | string | 是   | 短信接收的手机号:多个手机号请用英文逗号分隔,最大 1000 个号码。号码规则详见<br><br>“[3.2 手机号码规则]” | 示例:<br><br>138xxxxxxxx,130xxxxxxxx |
| content | string | 是   | 短信内容:最大支持 1000 个字(含签名),发送时请预留至少 10 个字的签名长度,一个字母或一个汉字都视为一个字。本字段支持加密传输,若不使用加密功能编码方式: UrlEncode("验证码:6666,打死都不要告诉别人哦！ ","UTF-8")<br><br>若使用加密功能详见“[3.4 内]容加密规则” | 示例:<br><br>短信内容:“验证码:6666,打死都不要告诉别人哦！”短信内容进行urlencode 编码后:<br><br>%e9%aa%8c%e8%af%81<br><br>%e7%a0%81%ef%bc%9a6 666%ef%bc%8c%e6%89%93%e6%ad%bb%e9%83% bd%e4%b8%8d%e8%a6%81%e5%91%8a%e8%af%8<br><br>9%e5%88%ab%e4%ba%ba<br><br>%e5%93%a6%ef%bc%81 |
| timestamp | string | 否   | 时间戳:24 小时制格式:<br><br>MMDDHHMMSS,即月日时<br><br>分秒,定长 10 位,月、日、时、分、秒每段不足 2 位时左补 0,密码选择 MD5 加密方式时必填该参数,密码选择明文方式时则不用填写 | 示例:0803192020 |
| svrtype | string | 否   | 业务类型:最大可支持 32 个长度的英文字母、数字组合的<br><br>字符串 | 示例:SMS001 |
| exno | string | 否   | 扩展号:长度不能超过 6 位,注意通道号+扩展号的总长度不能超过20 位,若超出则exno无效,如不需要扩展号则不用<br><br>提交此字段或填空 | 示例:0006 |
| custid | string | 否   | 用户自定义流水号:该条短信在您业务系统内的 ID,比如订单号或者短信发送记录的流水号。填写后发送状态返回值内将包含用户自定义流水号。<br><br>最大可支持 64 位的 ASCII 字符串:字母、数字、下划线、减号, 如不需要则不用提交<br><br>此字段或填空 | 示例:<br><br>b3d0a2783d31b21b8573 |
| exdata | string | 否   | 自定义扩展数据:额外提供的最大 64 个长度的 ASCII 字符串:字母、数字、下划线、减号,作为自定义扩展数据,填写后,状态报告返回时将会包含这部分数据,如不需要则不<br><br>用提交此字段或填空 |     |


### 返回参数说明

| 参数 | 类型 | 描述 | 示例 |
| --- | --- | --- | --- |
| result | int | 相同内容群发请求处理结果:<br><br>0:成功<br><br>非 0:失败,详见 [4 错误代码表] | 示例:0 |
| desc | string | 应答结果描述,当 result 非 0 时,为错误描述<br><br>编码方式:urlencode(UTF-8) | 示例:<br><br>“鉴权失败”<br><br>urlencode(UTF-8)编码:<br><br>%e9%89%b4%e6%9d%83<br><br>%e5%a4%b1%e8%b4%a5 |
| msgid | long (64 位) | 平台流水号:非 0,64 位整型,对应<br><br>Java 和 C#的 long,不可用 int 解析。<br><br>result 非 0 时,msgid 为 0 | 示例:<br><br>9223372036854775808<br><br>注意:msgid 允许出现负数 |
| custid | string | 用户自定义流水号:默认与请求报文中的 custid 保持一致,若请求报文中没有 custid 参数或值为空,则返回由梦网生成的代表本批短信的唯一编号<br><br>result 非 0 时,custid 为空 | 示例:<br><br>b3d0a2783d31b21b8573 |


### Post 请求示例

urlencode
```
userid=J10003&pwd=26dad7f364507df18f3841cc9c4ff94d&mobile=138xxxxxxxx, 130xxxxxxxx,180xxxxxxxx&content=%e9%aa%8c%e8%af%81%e7%a0%81%ef%bc%9a6 666%ef%bc%8c%e6%89%93%e6%ad%bb%e9%83%bd%e4%b8%8d%e8%a6%81%e5%91%8a%e8 %af%89%e5%88%ab%e4%ba%ba%e5%93%a6%ef%bc%81&timestamp=0803192020&svrty pe=SMS001&exno=0006&custid=b3d0a2783d31b21b8573&exdata
=exdata000002
```

JSON
```json
{"userid":"J10003","pwd":"26dad7f364507df18f3841cc9c4ff94d","mobile": "138xxxxxxxx,130xxxxxxxx,180xxxxxxxx","content":"%e9%aa%8c%e8%af%81%e 7%a0%81%ef%bc%9a6666%ef%bc%8c%e6%89%93%e6%ad%bb%e9%83%bd%e4%b8%8d%e8% a6%81%e5%91%8a%e8%af%89%e5%88%ab%e4%ba%ba%e5%93%a6%ef%bc%81","timesta mp":"0803192020","svrtype":"SMS001","exno":"0006","custid":"b3d0a2783d 31b21b8573","exdata":"exdata000002"}
```

XML
```xml
<?xml version=1.0 encoding=utf-8?>
<mtreq>
<userid>J10003</userid>
<pwd>26dad7f364507df18f3841cc9c4ff94d </pwd>
<mobile>138xxxxxxxx,130xxxxxxxx,180xxxxxxxx</mobile>
<content>%e9%aa%8c%e8%af%81%e7%a0%81%ef%bc%9a6666%ef%bc%8c%e6%89%93%e6%ad%bb%e9%83%bd%e4%b8%8d%e8%a6%81%e5%91%8a%e8%af%89%e5%88%ab%e4% ba%ba%e5%93%a6%ef%bc%81</content>
<timestamp>0803192020</timestamp>
<svrtype>SMS001</svrtype>
<exno>0006</exno>
<custid>b3d0a2783d31b21b8573</custid>
<exdata>exdata000002</exdata>
</mtreq>
```

### GET 请求示例

URLENCODE

用户名密码鉴权方式

userid=J10003&pwd=26dad7f364507df18f3841cc9c4ff94d&mobile=138xxxxxxxx, 139xxxxxxxx&content=%e9%aa%8c%e8%af%81%e7%a0%81%ef%bc%9a6666%ef%bc%8c %e6%89%93%e6%ad%bb%e9%83%bd%e4%b8%8d%e8%a6%81%e5%91%8a%e8%af%89%e5%88 %ab%e4%ba%ba%e5%93%a6%ef%bc%81&timestamp=0803192020&svrtype=SMS001&ex no=0006&custid=b3d0a2783d31b21b8573&exdata=exdata000002

apikey 鉴权方式

apikey=abade5589e2798f82f006bbc36d269ce&mobile=138xxxxxxxx,139xxxxxxx x&content=%e9%aa%8c%e8%af%81%e7%a0%81%ef%bc%9a6666%ef%bc%8c%e6%89%93% e6%ad%bb%e9%83%bd%e4%b8%8d%e8%a6%81%e5%91%8a%e8%af%89%e5%88%ab%e4%ba% ba%e5%93%a6%ef%bc%81&svrtype=SMS001&exno=0006&custid=b3d0a2783d31b21b

8573&exdata=exdata000002

### 发送成功返回示例

JSON (urlencode 与 JSON 请求都以 JSON 返回数据)
```json
{
"result":0, "desc":%e6%88%90%e5%8a%9f, "msgid":9223372036854775808,
"custid":"b3d0a2783d31b21b8573"

}
```

XML
```xml
<?xml version=1.0 encoding=utf-8?>
<mtrsp>
<result>0</result>
<desc>%e6%88%90%e5%8a%9f</desc>
<msgid>9223372036854775808</msgid>
<custid>b3d0a2783d31b21b8573</custid>
</mtrsp>
```

### 发送失败返回示例

JSON (urlencode 与 JSON 请求都以 JSON 返回数据)
```json
{
"result":-100999, "desc":"%e7%b3%bb%e7%bb%9f%e5%86%85%e9%83%a8%e9%94%99%e8%af%af", "msgid":0,
"custid":""

}
```

XML
```xml
<?xml version=1.0 encoding=utf-8?>
<mtrsp>
<result>-100999</result>
<desc>

%e7%b3%bb%e7%bb%9f%e5%86%85%e9%83%a8%e9%94%99%e8%af%af

</desc>
<msgid>0</msgid>
<custid></custid>
</mtrsp>
```

## 3. 个性化群发接口 send_multi

### 请求URL

<https://ip:port/sms/v2/std/send_multi> <http://ip:port/sms/v2/std/send_multi>

### 请求方式

POST/GET

### 请求参数说明

| 参数 | 类型 | 是否必须 | 描述 | 示例 |
| --- | --- | --- | --- | --- |
| userid | string | 是   | 用户账号:长度最大 6 个字符,统一大写,如提交参数中包含 apikey,则可以不用填写该参数及 pwd,两种鉴权方式中只<br><br>能选择一种方式来进行鉴权 | 示例:J10003 |
| pwd | string | 是   | 用户密码:定长小写 32 位字符,如提交参数中包含 apikey,则可以不用填写该参数及 userid,两种鉴权方式中只能选择一种方式来进行鉴权。密码规则详见“[3.1 鉴权规则]” | 示例:<br><br>密码明文模式:111111密码加密模式:<br><br>账号:J10003密码:111111<br><br>固定字符串:00000000时间戳:0803192020<br><br>MD5 加密之前的对应字符串: J1000300000000111111080 3192020<br><br>MD5 加密之后的密码字符串: 26dad7f364507df18f3841c<br><br>c9c4ff94d |
| timestamp | string | 否   | 时间戳:24 小时制格式:<br><br>MMDDHHMMSS,即月日时<br><br>分秒,定长 10 位,月、日、时、分、秒每段不足 2 位时左补 0,密码选择 MD5 加密方式时必填该参数,密码选择明文方式时则不用填写 | 示例:0803192020 |
| multimt | string | 是   | 个性化信息详情: 详见下表<br><br>3-2。<br><br>multimt 中最多可支持 500 个手机号的个性化信息 |     |


multimt包结构参数说明:

| 参数 | 类型 | 是否必须 | 描述 | 示例 |
| --- | --- | --- | --- | --- |
| mobile | string | 是   | 单个手机号: 号码规则详见<br><br>“[3.2 手机号码规则]” |     |
| content | string | 是   | 短信内容:最大支持 1000 个字(含签名),发送时请预留至少 10 个字的签名长度,一个字母或一个汉字都视为一个字。本字段支持加密传输,若不使用加密功能编码方式: UrlEncode("验证码:6666,打死 都 不 要 告 诉 别 人 哦 ！ ","UTF-8")<br><br>若使用加密功能详见“[3.4 内]容加密规则” | 示例:<br><br>短信内容:“验证码:6666,打死都不要告诉别人哦！”短信内容进行urlencode 编码后:<br><br>%e9%aa%8c%e8%af%81<br><br>%e7%a0%81%ef%bc%9a6 666%ef%bc%8c%e6%89%93%e6%ad%bb%e9%83% bd%e4%b8%8d%e8%a6%81%e5%91%8a%e8%af%8<br><br>9%e5%88%ab%e4%ba%ba<br><br>%e5%93%a6%ef%bc%81 |
| svrtype | string | 否   | 业务类型:最大可支持 32 个长度的英文字母、数字组合的<br><br>字符串 | 示例:SMS001 |
| exno | string | 否   | 扩展号:长度不能超过 6 位,注意通道号+扩展号的总长度不能超过20 位,若超出则exno无效,如不需要扩展号则不用<br><br>提交此字段或填空 | 示例:0006 |
| custid | string | 否   | 用户自定义流水号:该批次短信在您业务系统内的 ID,比如订单号或者短信发送记录的流水 号,multimt 包内的所有 custid建议与第一个 custid 保持一致,代表同一批次的短信。填写后发送状态返回值内将包含这个用户自定义流水号。最大可支持 64 位的 ASCII 字符串:字母、数字、下划线、减号,如不需要<br><br>则不用提交此字段或填空 | 示例:<br><br>b3d0a2783d31b21b8573 |
| exdata | string | 否   | 自定义扩展数据:额外提供的最大 64 个长度的 ASCII 字符串:字母、数字、下划线、减号,作为自定义扩展数据,填写后,状态报告返回时将会包含这部分数据,如不需要则不<br><br>用提交此字段或填空 |     |


### 返回参数说明

| 参数 | 类型 | 描述 | 示例 |
| --- | --- | --- | --- |
| result | int | 个性化群发请求处理结果:<br><br>0:成功<br><br>非 0:失败,详见 [4 错误代码表] | 示例:0 |
| desc | string | 应答结果描述,当 result 非 0 时,为错误描述<br><br>编码方式:urlencode(UTF-8) | 示例:<br><br>“鉴权失败”<br><br>urlencode(UTF-8)编码:<br><br>%e9%89%b4%e6%9d%83<br><br>%e5%a4%b1%e8%b4%a5 |
| msgid | long (64 位) | 平台流水号:返回个性化群发第一条记录中的 msgid,非 0,64 位整型, 对应 Java 和 C#的 long,不可用 int 解析。<br><br>result 非 0 时,msgid 为 0 | 示例:<br><br>9223372036854775808<br><br>注意:msgid 允许出现负数 |
| custid | string | 用户自定义流水号:默认与请求报文 multimt 包结构内第一条数据的 custid保持一致,若请求报文中没有 custid参数或值为空,则返回由梦网生成的代表本批短信的唯一编号<br><br>result 非 0 时,custid 为空 |     |


### Post 请求示例

urlencode
```
userid=J10003&pwd=26dad7f364507df18f3841cc9c4ff94d&timestamp=08031920 20&multimt=[{"mobile":"138xxxxxxxx","content":"%e9%aa%8c%e8%af%81%e7% a0%81%ef%bc%9a6666%ef%bc%8c%e6%89%93%e6%ad%bb%e9%83%bd%e4%b8%8d%e8%a6 %81%e5%91%8a%e8%af%89%e5%88%ab%e4%ba%ba%e5%93%a6%ef%bc%81","svrtype": "SMS001","exno":"0006","custid":"b3d0a2783d31b21b8573","exdata":"exda ta000001"},{"mobile":"131xxxxxxxx","content":"%e9%aa%8c%e8%af%81%e7%a 0%81%ef%bc%9a6666%ef%bc%8c%e6%89%93%e6%ad%bb%e9%83%bd%e4%b8%8d%e8%a6%81%e5%91%8a%e8%af%89%e5%88%ab%e4%ba%ba%e5%93%a6%ef%bc%81","svrtype":" SMS002","exno":"0007","custid":"b3d0a2783d31b21b8573","exdata":"exdat a000002"}] ```

JSON
```json
{"userid":"J10003","pwd":"26dad7f364507df18f3841cc9c4ff94d","timestam p":"0803192020","multimt":[
    {"mobile":"138xxxxxxxx","content":"%e9%aa%8c%e8%af%81%e7%a0%81%ef%bc%9a6666%ef%bc%8c%e6%89%93%e6%ad%bb%e9%83%bd%e4%b8%8d%e8%a6%81%e5%91%8a%e8%af%89%e5%88%ab%e4%ba%ba%e5%93%a6%ef%bc%81","svrtype":"SMS001","exno":"0006","custid":"b3d0a2783d31b21b8573", "exdata":"exdata000001"},
    {"mobile":"131xxxxxxxx","content":"%e9%aa%8c%e8%af%81%e7%a0%81%ef%bc%9a6666%ef%bc%8c%e6%89%93%e6%ad%bb%e9%83%bd%e 4%b8%8d%e8%a6%81%e5%91%8a%e8%af%89%e5%88%ab%e4%ba%ba%e5%93%a6%ef%bc%8 1","svrtype":"SMS002","exno":"0007","custid":"b3d0a2783d31b21b8573","exdata":"exdata000002"}
]}
```

XML
```xml
<?xml version=1.0 encoding=utf-8?>
<mtreq>
<userid>J10003</userid>
<pwd>26dad7f364507df18f3841cc9c4ff94d </pwd>
<timestamp>0803192020</timestamp>
<multimt>
<mt>
<mobile>138xxxxxxxx</mobile>
<content>%e9%aa%8c%e8%af%81%e7%a0%81%ef%bc%9a6666%ef%bc%8c%e6%8 9%93%e6%ad%bb%e9%83%bd%e4%b8%8d%e8%a6%81%e5%91%8a%e8%af%89%e5%88%ab%e 4%ba%ba%e5%93%a6%ef%bc%81</content>
<svrtype>SMS001</svrtype>
<exno>0006</exno>
<custid>b3d0a2783d31b21b8573</custid>
<exdata>exdata000002</exdata>
</mt>
<mt>
<mobile>131xxxxxxxx</mobile>
<content>%e9%aa%8c%e8%af%81%e7%a0%81%ef%bc%9a6666%ef%bc%8c%e6%8 9%93%e6%ad%bb%e9%83%bd%e4%b8%8d%e8%a6%81%e5%91%8a%e8%af%89%e5%88%ab%e 4%ba%ba%e5%93%a6%ef%bc%81</content>
<svrtype>SMS002</svrtype>
<exno>0007</exno>
<custid>b3d0a2783d31b21b8573</custid>
<exdata>exdata000002</exdata>
</mt>
</multimt>
</mtreq>
```

### GET 请求示例

URLENCODE

用户名密码鉴权方式
```
userid=J10003&pwd=26dad7f364507df18f3841cc9c4ff94d&timestamp=08031920 20&multimt=%5b%7b%22mobile%22%3a%2213800000000%22%2c%22content%22%3a%22%e9%aa%8c%e8%af%81%e7%a0%81%ef%bc%9a6666%ef%bc%8c%e6%89%93%e6%ad%bb %e9%83%bd%e4%b8%8d%e8%a6%81%e5%91%8a%e8%af%89%e5%88%ab%e4%ba%ba%e5%93 %a6%ef%bc%81%22%2c%22svrtype%22%3a%22SMS001%22%2c%22exno%22%3a%220006 %22%2c%22custid%22%3a%22b3d0a2783d31b21b8573%22%2c%22exdata%22%3a%22e xdata000001%22%7d%2c%7b%22mobile%22%3a%2213100000000%22%2c%22content%22%3a%22%e9%aa%8c%e8%af%81%e7%a0%81%ef%bc%9a6666%ef%bc%8c%e6%89%93%e6 %ad%bb%e9%83%bd%e4%b8%8d%e8%a6%81%e5%91%8a%e8%af%89%e5%88%ab%e4%ba%ba %e5%93%a6%ef%bc%81%22%2c%22svrtype%22%3a%22SMS002%22%2c%22exno%22%3a%220007%22%2c%22custid%22%3a%22b3d0a2783d31b21b8573%22%2c%22exdata%22%3a%22exdata000002%22%7d%5d ```

apikey 鉴权方式
```
apikey=abade5589e2798f82f006bbc36d269ce&multimt=%5b%7b%22mobile%22%3a %2213800000000%22%2c%22content%22%3a%22%e9%aa%8c%e8%af%81%e7%a0%81%ef %bc%9a6666%ef%bc%8c%e6%89%93%e6%ad%bb%e9%83%bd%e4%b8%8d%e8%a6%81%e5%9 1%8a%e8%af%89%e5%88%ab%e4%ba%ba%e5%93%a6%ef%bc%81%22%2c%22svrtype%22%3a%22SMS001%22%2c%22exno%22%3a%220006%22%2c%22custid%22%3a%22b3d0a278 3d31b21b8573%22%2c%22exdata%22%3a%22exdata000001%22%7d%7b%22mobile%22 %3a%2213100000000%22%2c%22content%22%3a%22%e9%aa%8c%e8%af%81%e7%a0%81 %ef%bc%9a6666%ef%bc%8c%e6%89%93%e6%ad%bb%e9%83%bd%e4%b8%8d%e8%a6%81%e 5%91%8a%e8%af%89%e5%88%ab%e4%ba%ba%e5%93%a6%ef%bc%81%22%2c%22svrtype%22%3a%22SMS002%22%2c%22exno%22%3a%220007%22%2c%22custid%22%3a%22b3d0a
2783d31b21b8573%22%2c%22exdata%22%3a%22exdata000002%22%7d%5d
```

### 发送成功返回示例

JSON (urlencode 与 JSON 请求都以 JSON 返回数据)
```json
{
"result":0, "desc":%e6%88%90%e5%8a%9f, "msgid":9223372036854775808,
"custid":b3d0a2783d31b21b8573

}
```

XML
```xml
<?xml version=1.0 encoding=utf-8?>
<mtrsp>
<result>0</result>
<desc>%e6%88%90%e5%8a%9f</desc>
<msgid>9223372036854775808</msgid>
<custid>b3d0a2783d31b21b8573</custid>
</mtrsp>
```

### 发送失败返回示例

JSON (urlencode 与 JSON 请求都以 JSON 返回数据)
```json
{
"result":-100999, "desc":%e7%b3%bb%e7%bb%9f%e5%86%85%e9%83%a8%e9%94%99%e8%af%af, "msgid":0,
"custid":""

}
```

XML
```xml
<?xml version=1.0 encoding=utf-8?>
<mtrsp>
<result>-100999</result>
<desc>

%e7%b3%bb%e7%bb%9f%e5%86%85%e9%83%a8%e9%94%99%e8%af%af

</desc>
<msgid>0</msgid>
<custid></custid>
</mtrsp>
```

## 4. 个性化群发接口 send_mixed

### 请求URL

<https://ip:port/sms/v2/std/send_mixed> <http://ip:port/sms/v2/std/send_mixed>

### 请求方式

POST/GET

### 请求参数说明

| 参数 | 类型 | 是否必须 | 描述 | 示例 |
| --- | --- | --- | --- | --- |
| userid | string | 是   | 用户账号:长度最大 6 个字符,统一大写,如提交参数中包含 apikey,则可以不用填写该参数及 pwd,两种鉴权方式中只<br><br>能选择一种方式来进行鉴权 | 示例:J10003 |
| pwd | string | 是   | 用户密码:定长小写 32 位字符,如提交参数中包含 apikey,则可以不用填写该参数及 userid,两种鉴权方式中只能选择一种方式来进行鉴权。密码规则详见“[3.1 鉴权规则]” | 示例:<br><br>密码明文模式:111111密码加密模式:<br><br>账号:J10003密码:111111<br><br>固定字符串:00000000时间戳:0803192020<br><br>MD5 加密之前的对应字符串: J1000300000000111111080 3192020<br><br>MD5 加密之后的密码字符串: 26dad7f364507df18f3841c<br><br>c9c4ff94d |
| timestamp | string | 否   | 时间戳:24 小时制格式:<br><br>MMDDHHMMSS,即月日时<br><br>分秒,定长 10 位,月、日、时、分、秒每段不足 2 位时左补 0,密码选择 MD5 加密方式时必填该参数,密码选择明文方式时则不用填写 | 示例:0803192020 |
| mobile | string | 是   | 短信接收的手机号:多个手机号请用英文逗号分隔,最大 500 个号码。号码规则详见“[3.2]<br><br>手机号码规则” | 示例:<br><br>138xxxxxxxx,130xxxxxxxx |
| content | string | 是   | 短信内容:最大支持 1000 个字(含签名),发送时请预留至少 10 个字的签名长度,一个字母或一个汉字都视为一个 字。多个内容以英文逗号分隔,信息内容与手机号顺序一一对应。<br><br>如果信息内容数量与手机号个数不一致的情况将返回错误。<br><br>本字段支持加密传输,若不使用加密功能编码方式: UrlEncode("验证码:6666,打死 都 不 要 告 诉 别 人 哦 ！ ","UTF-8")+","+UrlEncode ("<br><br>验证码:8888,打死都不要告诉别人哦！","UTF-8") 使用 x-www-form-urlencoded方式提交时,请对结果再进行一次 UrlEncode 编码。<br><br>例如: UrlEncode(UrlEncode("验证 码:6666,打死都不要告诉别人哦！<br><br>","UTF-8")+","+UrlEncode ("<br><br>验证码:8888,打死都不要告诉别人哦！<br><br>","UTF-8"),"UTF-8")<br><br>若使用加密功能详见“[3.4 内]容加密规则” | 示例:<br><br>短信内容:“验证码:6666,打死都不要告诉别人哦！,验证码:8888,打死都不要告诉别人哦！”<br><br>短信内容进行urlencode 编码后:<br><br>%e9%aa%8c%e8%af%81<br><br>%e7%a0%81%ef%bc%9a6 666%ef%bc%8c%e6%89%93%e6%ad%bb%e9%83% bd%e4%b8%8d%e8%a6%81%e5%91%8a%e8%af%8<br><br>9%e5%88%ab%e4%ba%ba<br><br>%e5%93%a6%ef%bc%81,<br><br>%e9%aa%8c%e8%af%81<br><br>%e7%a0%81%ef%bc%9a8 888%ef%bc%8c%e6%89%93%e6%ad%bb%e9%83% bd%e4%b8%8d%e8%a6%81%e5%91%8a%e8%af%8<br><br>9%e5%88%ab%e4%ba%ba<br><br>%e5%93%a6%ef%bc%81 |
| timestamp | string | 否   | 时间戳:24 小时制格式:<br><br>MMDDHHMMSS,即月日时<br><br>分秒,定长 10 位,月、日、时、分、秒每段不足 2 位时左补 0,密码选择 MD5 加密方式时必填该参数,密码选择明文方式时则不用填写 | 示例:0803192020 |
| svrtype | string | 否   | 业务类型:最大可支持 32 个长度的英文字母、数字组合的<br><br>字符串 | 示例:SMS001 |
| exno | string | 否   | 扩展号:长度不能超过 6 位,注意通道号+扩展号的总长度不能超过20 位,若超出则exno无效,如不需要扩展号则不用<br><br>提交此字段或填空 | 示例:0006 |
| custid | string | 否   | 用户自定义流水号:该条短信在您业务系统内的 ID,比如订单号或者短信发送记录的流水号。填写后发送状态返回值内将包含用户自定义流水号。最大可支持 64 位的 ASCII 字符串:字母、数字、下划线、减号, 如不需要则不用提交<br><br>此字段或填空 | 示例:<br><br>b3d0a2783d31b21b8573 |
| exdata | string | 否   | 自定义扩展数据:额外提供的最大 64 个长度的 ASCII 字符串:字母、数字、下划线、减号,作为自定义扩展数据,填写后,状态报告返回时将会包<br><br>含这部分数据,如不需要则不用提交此字段或填空 |     |


### 返回参数说明

| 参数 | 类型 | 描述 | 示例 |
| --- | --- | --- | --- |
| result | int | 个性化群发请求处理结果:<br><br>0:成功<br><br>非 0:失败,详见 [4 错误代码表] | 示例:0 |
| desc | string | 应答结果描述,当 result 非 0 时,为错误描述<br><br>编码方式:urlencode(UTF-8) | 示例:<br><br>“鉴权失败”<br><br>urlencode(UTF-8)编码:<br><br>%e9%89%b4%e6%9d%83<br><br>%e5%a4%b1%e8%b4%a5 |
| msgid | long (64 位) | 平台流水号:返回个性化群发第一条记录中的 msgid,非 0,64 位整型, 对应 Java 和 C#的 long,不可用 int 解析。<br><br>result 非 0 时,msgid 为 0 | 示例:<br><br>9223372036854775808<br><br>注意:msgid 允许出现负数 |
| custid | string | 用户自定义流水号:默认与请求报文中的 custid 保持一致,若请求报文中没有 custid 参数或值为空,则返回由梦网生成的代表本批短信的唯一编号<br><br>result 非 0 时,custid 为空 |     |


### Post 发送请求示例

urlencode
```
userid=J10003&pwd=26dad7f364507df18f3841cc9c4ff94d&mobile=138xxxxxxxx ,130xxxxxxxx&content=%e9%aa%8c%e8%af%81%e7%a0%81%ef%bc%9a6666%ef%bc%8 c%e6%89%93%e6%ad%bb%e9%83%bd%e4%b8%8d%e8%a6%81%e5%91%8a%e8%af%89%e5%8 8%ab%e4%ba%ba%e5%93%a6%ef%bc%81,%e9%aa%8c%e8%af%81%e7%a0%81%ef%bc%9a8 888%ef%bc%8c%e6%89%93%e6%ad%bb%e9%83%bd%e4%b8%8d%e8%a6%81%e5%91%8a%e8 %af%89%e5%88%ab%e4%ba%ba%e5%93%a6%ef%bc%81&timestamp=0803192020&svrty
pe=SMS001&exno=0006&custid=b3d0a2783d31b21b8573&exdata=exdata000002
```

JSON
```json
{"userid":"J10003","pwd":"26dad7f364507df18f3841cc9c4ff94d","mobile": "138xxxxxxxx,130xxxxxxxx","content":"%e9%aa%8c%e8%af%81%e7%a0%81%ef%b c%9a6666%ef%bc%8c%e6%89%93%e6%ad%bb%e9%83%bd%e4%b8%8d%e8%a6%81%e5%91%8a%e8%af%89%e5%88%ab%e4%ba%ba%e5%93%a6%ef%bc%81,%e9%aa%8c%e8%af%81%e7 %a0%81%ef%bc%9a8888%ef%bc%8c%e6%89%93%e6%ad%bb%e9%83%bd%e4%b8%8d%e8%a 6%81%e5%91%8a%e8%af%89%e5%88%ab%e4%ba%ba%e5%93%a6%ef%bc%81","timestam p":"0803192020","svrtype":"SMS001","exno":"0006","custid":"b3d0a2783d3 1b21b8573","exdata":"exdata000002"} 
```

XML
```xml
<?xml version=1.0 encoding=utf-8?>
<mtreq>
<userid>J10003</userid>
<pwd>26dad7f364507df18f3841cc9c4ff94d</pwd>
<mobile>138xxxxxxxx,130xxxxxxxx</mobile>
<content>%e9%aa%8c%e8%af%81%e7%a0%81%ef%bc%9a6666%ef%bc%8c%e6%89%93%e6%ad%bb%e9%83%bd%e4%b8%8d%e8%a6%81%e5%91%8a%e8%af%89%e5%88%ab%e4% ba%ba%e5%93%a6%ef%bc%81,%e9%aa%8c%e8%af%81%e7%a0%81%ef%bc%9a8888%ef%b c%8c%e6%89%93%e6%ad%bb%e9%83%bd%e4%b8%8d%e8%a6%81%e5%91%8a%e8%af%89%e 5%88%ab%e4%ba%ba%e5%93%a6%ef%bc%81</content>
<timestamp>0803192020</timestamp>
<svrtype>SMS001</svrtype>
<exno>0006</exno>
<custid>b3d0a2783d31b21b8573</custid>
<exdata>exdata000002</exdata>
</mtreq>
```

### GET 请求示例

URLENCODE

用户名密码鉴权方式
```
userid=J10003&pwd=26dad7f364507df18f3841cc9c4ff94d&mobile=138xxxxxxxx, 130xxxxxxxx&content=%e9%aa%8c%e8%af%81%e7%a0%81%ef%bc%9a6666%ef%bc%8c %e6%89%93%e6%ad%bb%e9%83%bd%e4%b8%8d%e8%a6%81%e5%91%8a%e8%af%89%e5%88 %ab%e4%ba%ba%e5%93%a6%ef%bc%81,%e9%aa%8c%e8%af%81%e7%a0%81%ef%bc%9a88 88%ef%bc%8c%e6%89%93%e6%ad%bb%e9%83%bd%e4%b8%8d%e8%a6%81%e5%91%8a%e8% af%89%e5%88%ab%e4%ba%ba%e5%93%a6%ef%bc%81&timestamp=0803192020&svrtyp e=SMS001&exno=0006&custid=b3d0a2783d31b21b8573&exdata=exdata000002 ```
```

apikey 鉴权方式
```
apikey=abade5589e2798f82f006bbc36d269ce&mobile=138xxxxxxxx,130xxxxxxx x&content=%e9%aa%8c%e8%af%81%e7%a0%81%ef%bc%9a6666%ef%bc%8c%e6%89%93% e6%ad%bb%e9%83%bd%e4%b8%8d%e8%a6%81%e5%91%8a%e8%af%89%e5%88%ab%e4%ba% ba%e5%93%a6%ef%bc%81,%e9%aa%8c%e8%af%81%e7%a0%81%ef%bc%9a8888%ef%bc%8 c%e6%89%93%e6%ad%bb%e9%83%bd%e4%b8%8d%e8%a6%81%e5%91%8a%e8%af%89%e5%8 8%ab%e4%ba%ba%e5%93%a6%ef%bc%81&timestamp=0803192020&svrtype=SMS001&exno=0006&custid=b3d0a2783d31b21b8573&exdata=exdata000002
```

### 发送成功返回示例

JSON(urlencode 与 JSON 请求都以 JSON 返回数据)
```json
{
"result":0, "desc":%e6%88%90%e5%8a%9f, "msgid":9223372036854775808,
"custid":"b3d0a2783d31b21b8573"

}
```

XML
```xml
<?xml version=1.0 encoding=utf-8?>
<mtrsp>
<result>0</result>
<desc>%e6%88%90%e5%8a%9f</desc>
<msgid>9223372036854775808</msgid>
<custid>b3d0a2783d31b21b8573</custid>
</mtrsp>
```

### 发送失败返回示例

JSON(urlencode 与 JSON 请求都以 JSON 返回数据)
```json
{
"result":-100999, "desc":%e7%b3%bb%e7%bb%9f%e5%86%85%e9%83%a8%e9%94%99%e8%af%af, "msgid":0,
"custid":""

}
```

XML
```xml
<?xml version=1.0 encoding=utf-8?>
<mtrsp>
<result>-100999</result>
<desc>

%e7%b3%bb%e7%bb%9f%e5%86%85%e9%83%a8%e9%94%99%e8%af%af

</desc>
<msgid>0</msgid>
<custid></custid>
</mtrsp>
```

## 5. 模板发送接口 send_template

### 请求URL

<https://ip:port/sms/v2/std/send_template> <http://ip:port/sms/v2/std/send_template>

### 请求方式

POST/GET

### 请求参数说明

| 参数 | 类型 | 是否必须 | 描述 | 示例 |
| --- | --- | --- | --- | --- |
| userid | string | 是   | 用户账号:长度最大 6 个字符,统一大写,如提交参数中包含 apikey,则可以不用填写该参数及 pwd,两种鉴权方式中只<br><br>能选择一种方式来进行鉴权 | 示例:J10003 |
| pwd | string | 是   | 用户密码:定长小写 32 位字符,如提交参数中包含 apikey,则可以不用填写该参数及 userid,两种鉴权方式中只能选择一种方式来进行鉴权。 | 示例:<br><br>密码明文模式:111111密码加密模式:<br><br>账号:J10003密码:111111<br><br>固定字符串:00000000时间戳:0803192020<br><br>MD5 加密之前的对应字符串: J1000300000000111111080 3192020<br><br>MD5 加密之后的密码字符串: 26dad7f364507df18f3841c<br><br>c9c4ff94d |
| tmplid | string | 是   | 短信模版编号:长度最大20位<br><br>字符 | 示例:200170 |
| mobile | string | 是   | 短信接收的手机号:多个手机号请用英文逗号分隔,最大 1000 个号码。号码规则详见<br><br>“[3.2 手机号码规则]”。 | 示 例 : 138xxxxxxxx, 130xxxxxxxx |
| content | string | 是   | 变量名和变量值:一个模板变量名对应一个变量值,多个变量使用key=value的方式进行拼接,需要进行两次urlenco de 编码,第一次只将变量名<br><br>(key)及变量值(value)进行编码,第二次将整体内容进行编码,本字段支持加密传输 若不使用加密功能编码方式: UrlEncode(P1=UrlEncode("梦网 科 技 ","UTF-8")&P2=UrlEncode("2<br><br>01701","UTF-8"),"UTF-8")<br><br>若使用加密功能详见“[3.4 内]容加密规则”  | 示例:<br><br>P1=梦网科技&P2=201701<br><br>变量名(key)与变量值 (value)进行 urlencode 编码后:P1=%e6%a2%a6%e7%bd<br><br>%91%e7%a7%91%e6%8a<br><br>%801&P2=201701<br><br>整体内容再次进行 urlencode 编码后: P1%3d%25e6%25a2%25a6<br><br>%25e7%25bd%2591%25e7<br><br>%25a7%2591%25e6%258a<br><br>%25801%26P2%3d201701 |
| timestamp | string | 否   | 时间戳:24 小时制格式:<br><br>MMDDHHMMSS,即月日时<br><br>分秒,定长 10 位,月、日、时、分、秒每段不足 2 位时左补 0,密码选择 MD5 加密方式时必填该参数,密码选择明文方式时则不用填写 | 示例:0803192020 |
| svrtype | string | 否   | 业务类型:最大可支持 32 个长度的英文字母、数字组合的<br><br>字符串 | 示例:SMS001 |
| exno | string | 否   | 扩展号:长度不能超过 6 位,注意通道号+扩展号的总长度不能超过20 位,若超出则exno无效,如不需要扩展号则不用<br><br>提交此字段或填空 | 示例:0006 |
| custid | string | 否   | 用户自定义流水号:该条短信在您业务系统内的 ID,比如订单号或者短信发送记录的流水号。填写后发送状态返回值内将包含用户自定义流水号。最大可支持 64 位的 ASCII 字符串:字母、数字、下划线、减号,如不需要则不用提交此<br><br>字段或填空 | 示例:<br><br>b3d0a2783d31b21b8573 |
| exdata | string | 否   | 自定义扩展数据:额外提供的最大 64 个长度的 ASCII 字符串:字母、数字、下划线、减号,作为自定义扩展数据,填写后,状态报告返回时将会包含这部分数据,如不需要则不<br><br>用提交此字段或填空 |     |


### 返回参数说明

| 参数 | 类型 | 描述 | 示例 |
| --- | --- | --- | --- |
| result | int | 短信发送请求处理结果:<br><br>0:成功 非 0:失败 | 示例:0 |
| desc | string | 应答结果描述,当 result 非 0 时,为错误描述<br><br>编码方式:urlencode(UTF-8) | 示例:<br><br>“鉴权失败”<br><br>urlencode(UTF-8)编码:<br><br>%e9%89%b4%e6%9d%83<br><br>%e5%a4%b1%e8%b4%a5 |
| msgid | long (64 位) | 平台流水号:非 0,64 位整型,对应<br><br>Java 和 C#的 long,不可用 int 解析。<br><br>result 非 0 时,msgid 为 0 | 示例:<br><br>9223372036854775808<br><br>注意:msgid 允许出现负数 |
| custid | string | 用户自定义流水号:默认与请求报文中的 custid 保持一致,若请求报文中没有 custid 参数或值为空,则返回网关生成的平台流水号<br><br>result 非 0 时,custid 为空 | 示例:<br><br>b3d0a2783d31b21b8573 |


### Post 请求示例

urlencode

userid=J10003&pwd=26dad7f364507df18f3841cc9c4ff94d&tmplid=20170&mobi le=138xxxxxxxx&content=P1%3d%25e6%25a2%25a6%25e7%25bd%2591%25e7%25a7%2591%25e6%258a%2580%26P2%3d201701&timestamp=0803192020&svrtype=SMS001
&exno=0006&custid=b3d0a2783d31b21b8573&exdata=exdata000002

JSON
```json
{"userid":"J10003","pwd":"26dad7f364507df18f3841cc9c4ff94d","tmplid ": "20170","mobile":"138xxxxxxxx","content":"P1%3d%25e6%25a2%25a6%25e7%2 5bd%2591%25e7%25a7%2591%25e6%258a%2580%26P2%3d201701","timestamp":"08 03192020","svrtype":"SMS001","exno":"0006","custid":"b3d0a2783d31b21b8 573","exdata":"exdata000002"}
```

XML
```xml
<?xml version=1.0 encoding=utf-8?>
<mtreq>
<userid>J10003</userid>
<pwd>26dad7f364507df18f3841cc9c4ff94d </pwd>
<tmplid >20170</tmplid >
<mobile>138xxxxxxxx</mobile>
<content>P1%3d%25e6%25a2%25a6%25e7%25bd%2591%25e7%25a7%2591%25e6%

258a%2580%26P2%3d201701</content>
<timestamp>0803192020</timestamp>
<svrtype>SMS001</svrtype>
<exno>0006</exno>
<custid>b3d0a2783d31b21b8573</custid>
<exdata>exdata000002</exdata>
</mtreq>
```

### GET 请求示例

URLENCODE

用户名密码鉴权方式

userid=J10003&pwd=26dad7f364507df18f3841cc9c4ff94d&mobile=138xxxxxxxx, 130xxxxxxxx&content=P1%3d%25e6%25a2%25a6%25e7%25bd%2591%25e7%25a7%259 1%25e6%258a%2580%26P2%3d201701&timestamp=0803192020&svrtype=SMS001&ex no=0006&custid=b3d0a2783d31b21b8573&exdata=exdata000002

apikey 鉴权方式

apikey=abade5589e2798f82f006bbc36d269ce&mobile=138xxxxxxxx,139xxxxxxx x&content=P1%3d%25e6%25a2%25a6%25e7%25bd%2591%25e7%25a7%2591%25e6%258

a%2580%26P2%3d201701&svrtype=SMS001&exno=0006&custid=b3d0a2783d31b21b 8573&exdata=exdata000002

### 发送成功返回示例

JSON (urlencode 与 JSON 请求都以 JSON 返回数据)
```json
{
"result":0, "desc":%e6%88%90%e5%8a%9f, "msgid":9223372036854775808,
"custid":"b3d0a2783d31b21b8573"

}
```

XML

```
<?xml version=1.0 encoding=utf-8?>
<mtrsp>
<result>0</result>
<desc>%e6%88%90%e5%8a%9f</desc>
<msgid>9223372036854775808</msgid>
<custid>b3d0a2783d31b21b8573</custid>
</mtrsp>
```

### 发送失败返回示例

JSON (urlencode 与 JSON 请求都以 JSON 返回数据)
```json
{
"result":-100999, "desc":%e7%b3%bb%e7%bb%9f%e5%86%85%e9%83%a8%e9%94%99%e8%af%af, "msgid":0,
"custid":""

}
```

XML
```xml
<?xml version=1.0 encoding=utf-8?>
<mtrsp>
<result>-100999</result>
<desc>

%e7%b3%bb%e7%bb%9f%e5%86%85%e9%83%a8%e9%94%99%e8%af%af

</desc>
<msgid>0</msgid>
<custid></custid>
</mtrsp>
```

## 6. 获取上行接口 get_mo

### 请求URL

<https://ip:port/sms/v2/std/get_mo> <http://ip:port/sms/v2/std/get_mo>

### 请求方式

POST/GET

### 请求参数说明

| 参数 | 类型 | 是否必须 | 描述 | 示例 |
| --- | --- | --- | --- | --- |
| userid | string | 是   | 用户账号:长度最大 6 个字符,统一大写,如提交参数中包含 apikey,则可以不用填写该参<br><br>数及 pwd,两种鉴权方式中只能选择一种方式来进行鉴权 | 示例:J10003 |
| pwd | string | 是   | 用户密码:定长小写 32 位字符,如提交参数中包含 apikey,则可以不用填写该参数及 userid,两种鉴权方式中只能选择一种方式来进行鉴权。密码规则详见“[3.1 鉴权规则]” | 示例:<br><br>密码明文模式:111111密码加密模式:<br><br>账号:J10003密码:111111<br><br>固定字符串:00000000时间戳:0803192020<br><br>MD5 加密之前的对应字符串: J1000300000000111111080 3192020<br><br>MD5 加密之后的密码字符串: 26dad7f364507df18f3841c<br><br>c9c4ff94d |
| timestamp | string | 否   | 时间戳:24 小时制格式:<br><br>MMDDHHMMSS,即月日时<br><br>分秒,定长 10 位,月、日、时、分、秒每段不足 2 位时左补 0,密码选择 MD5 加密方式时必填该参数,密码选择明文方式时则不用填写 | 示例:0803192020 |
| retsize | int | 否   | 每次请求想要获取的上行最大条数。<br><br>最大 200,超过 200 按 200 返<br><br>回。小于等于 0 或不填时,系<br><br>统返回默认条数,默认 200 条 | 示例:200 |


### 请求方式

POST/GET

### 返回参数说明

| 参数 | 类型 | 是否必须 | 描述 | 示例 |
| --- | --- | --- | --- | --- |
| result | int | 是   | 获取上行请求处理结果:<br><br>0:成功<br><br>非 0:失败,详见 [4] [错误代码]| desc | string | 否   | 应答结果描述,当 result 非 0时,为错误描述<br><br>编码方式:urlencode(UTF-8) | 示例:<br><br>“鉴权失败”<br><br>urlencode(UTF-8)编码:<br><br>%e9%89%b4%e6%9d%83<br><br>%e5%a4%b1%e8%b4%a5 |
| mos | string | 是   | result 非 0 时 mos 为空<br><br>格式见下

mos包结构参数说明:

| 参数 | 类型 | 是否必须 | 描述 | 示例 |
| --- | --- | --- | --- | --- |
| msgid | long(64位) | 是   | 平台流水号:上行在梦网云通信平台中的唯一编号 | 示例:<br><br>9223372045854775808<br><br>注意:msgid允许出现负数 |
| mobile | string | 是   | 手机号:号码规则详见“[3.2 手]机号码规则” |     |
| countrycode | int | 是   | 手机号的国际区号:-1 表示无效号码 | 示例:86 |
| pknum | int | 是   | 当前条数 | 示例:1 |
| pktotal | int | 是   | 总条数 | 示例:2 |
| spno | string | 是   | 完整的通道号 | 示例:955337890 |
| exno | string | 是   | 下行时填写的exno | 示例:7890 |
| rtime | string | 是   | 上行返回的时间:<br><br>YYYY-MM-DD HH:MM:SS | 示例:<br><br>2016-08-04 17:38:59 |
| content | string | 是   | 短信内容:最大支持 70 个字 (含签名),一个字母或一个汉字都视为一个字。本字段支持加密传输,若不使用加密功能编码方式:<br><br>UrlEncode("验证码:6666,打死 都 不 要 告 诉 别 人 哦 ！ ","UTF-8")<br><br>若使用加密功能详见“[3.4 内]容加密规则” | 示例:<br><br>短信内容:“验证码:6666,打死都不要告诉别人哦！”短信内容进行urlencode 编码后:<br><br>%e9%aa%8c%e8%af%81<br><br>%e7%a0%81%ef%bc%9a6 666%ef%bc%8c%e6%89%93%e6%ad%bb%e9%83% bd%e4%b8%8d%e8%a6%81%e5%91%8a%e8%af%8<br><br>9%e5%88%ab%e4%ba%ba<br><br>%e5%93%a6%ef%bc%81 |


### Post 请求示例

urlencode

userid=J10003&pwd=26dad7f364507df18f3841cc9c4ff94d&timestamp=0803192020&retsize=300

JSON
```json
{"userid":"J10003","pwd":"26dad7f364507df18f3841cc9c4ff94d","timestam p":"0803192020","retsize":300}
```

XML
```xml
<?xml version=1.0 encoding=utf-8?>
<moreq>
<userid>J10003</userid>
<pwd>26dad7f364507df18f3841cc9c4ff94d </pwd>
<timestamp>0803192020</timestamp>
<retsize>300</retsize>
</moreq>
```

### GET 请求示例

URLENCODE

用户名密码鉴权方式

userid=J10003&pwd=26dad7f364507df18f3841cc9c4ff94d&timestamp=0803192020&retsize=300

apikey 鉴权方式

apikey=abade5589e2798f82f006bbc36d269ce&retsize=300

### 获取成功返回示例

JSON (urlencode 与 JSON 请求都以 JSON 返回数据)
```json
{
"result":0, "desc":%e6%88%90%e5%8a%9f, "mos": [

{
"msgid":9223372045854775808,
"mobile":"138xxxxxxxx", "countrycode":86, "spno":"1000457890006", "exno":"0006",
"rtime":"2016-08-04 17:38:59",
"content":"%e9%aa%8c%e8%af%81%e7%a0%81%ef%bc%9a6666%ef%bc%8c%e6%89%93%e6%ad%bb%e9%83%bd%e4%b8%8d%e8%a6%81%e5%91%8a%e8%af%89%e5%88%ab% e4%ba%ba%e5%93%a6%ef%bc%81"

},

{
"msgid":9223372045854895809,
"mobile":"131xxxxxxxx", "countrycode":86, "spno":"1000357890006", "exno":"0006",
"rtime":"2016-08-04 17:39:50",
"content":"%e9%aa%8c%e8%af%81%e7%a0%81%ef%bc%9a6666%ef%bc%8c%e6%89%93%e6%ad%bb%e9%83%bd%e4%b8%8d%e8%a6%81%e5%91%8a%e8%af%89%e5%88%ab% e4%ba%ba%e5%93%a6%ef%bc%81"

}

]

}
```

XML
```xml
<?xml version=1.0 encoding=utf-8?>
<morsp>
<result>0</result>
<desc>%e6%88%90%e5%8a%9f</desc>
<mos>
<mo>
<msgid>9223372045854775808</msgid>
<mobile>138xxxxxxxx</mobile>
<countrycode>86</countrycode>
<spno>1000457890006</spno>
<exno>0006</exno>
<rtime>2016-08-04 17:38:59</rtime>
<content>%e9%aa%8c%e8%af%81%e7%a0%81%ef%bc%9a6666%ef%bc%8c%e6%89%93%e6%ad%bb%e9%83%bd%e4%b8%8d%e8%a6%81%e5%91%8a%e8%af%89%e5%88%ab% e4%ba%ba%e5%93%a6%ef%bc%81</content>
</mo>
<mo>
<msgid>9223372045854895809</msgid>
<mobile>131xxxxxxxx</mobile>
<countrycode>86</countrycode>
<spno>1000357890006</spno>
<exno>0006</exno>
<rtime>2016-08-04 17:39:50</rtime>
<content>%e9%aa%8c%e8%af%81%e7%a0%81%ef%bc%9a6666%ef%bc%8c%e6%89%93%e6%ad%bb%e9%83%bd%e4%b8%8d%e8%a6%81%e5%91%8a%e8%af%89%e5%88%ab% e4%ba%ba%e5%93%a6%ef%bc%81</content>
</mo>
</mos>
</morsp>
```

### 获取失败返回示例

JSON (urlencode 与 JSON 请求都以 JSON 返回数据)
```json
{
"result":-100999, "desc":%e7%b3%bb%e7%bb%9f%e5%86%85%e9%83%a8%e9%94%99%e8%af%af, "mos":""

}
``

XML

```
<?xml version=1.0 encoding=utf-8?>
<morsp>
<result>-100999</result>
<desc>

%e7%b3%bb%e7%bb%9f%e5%86%85%e9%83%a8%e9%94%99%e8%af%af

</desc>
<mos></mos>
</morsp>
```

## 7. 获取状态报告接口 get_rpt

### 请求URL

<https://ip:port/sms/v2/std/get_rpt> <http://ip:port/sms/v2/std/get_rpt>

### 请求方式

POST/GET

### 请求参数说明

| 参数 | 类型 | 是否必须 | 描述 | 示例 |
| --- | --- | --- | --- | --- |
| userid | string | 是   | 用户账号:长度最大 6 个字符,统一大写,如提交参数中包含 apikey,则可以不用填写该参数及 pwd,两种鉴权方式中只<br><br>能选择一种方式来进行鉴权 | 示例:J10003 |
| pwd | string | 是   | 用户密码:定长小写 32 位字符,如提交参数中包含 apikey,则可以不用填写该参数及 userid,两种鉴权方式中只能选择一种方式来进行鉴权。密码规则详见“[3.1 鉴权规则]” | 示例:密码明文模式:111111密码加密模式:<br><br>账号:J10003密码:111111<br><br>固定字符串:00000000时间戳:0803192020<br><br>MD5 加密之前的对应字符串: J1000300000000111111080 3192020<br><br>MD5 加密之后的密码字符串: 26dad7f364507df18f3841c<br><br>c9c4ff94d  |
| timestamp | string | 否   | 时间戳:24 小时制格式:<br><br>MMDDHHMMSS,即月日时<br><br>分秒,定长 10 位,月、日、时、分、秒每段不足 2 位时左补 0,密码选择 MD5 加密方式时必填该参数,密码选择明文方式时则不用填写 | 示例:0803192020 |
| apiver | string | 否   | 调用的 API 接口版本。 | V5.7.2 |
| retsize | int | 否   | 本次请求想要获取的状态报告最大条数。<br><br>最大 500,超过 500 按 500 返<br><br>回。小于等于 0 或不填时,系<br><br>统返回默认条数,默认 500 条 | 示例:200 |


### 返回参数说明

| 参数 | 类型 | 描述 | 示例 |
| --- | --- | --- | --- |
| result | int | 获取状态报告请求处理结果:<br><br>0:成功<br><br>非 0:失败,详见 [4 错误代码表] | 示例:0 |
| desc | string | 应答结果描述,当 result 非 0 时,为错误描述<br><br>编码方式:urlencode(UTF-8) | 示例:<br><br>“鉴权失败”<br><br>urlencode(UTF-8)编码:<br><br>%e9%89%b4%e6%9d%83<br><br>%e5%a4%b1%e8%b4%a5 |
| rpts | string | result 非 0 时 rpts 为空<br><br>格式见下

rpts包结构参数说明:

| 参数 | 类型 | 描述 | 示例 |
| --- | --- | --- | --- |
| msgid | long(64 位) | 平台流水号:对应下行请求返回结果中的 msgid | 示例:<br><br>9223372036854775808<br><br>注意:msgid 允许出现负数 |
| custid | string | 用户自定义流水号:对应下行请求时填写的 custid | 示例:<br><br>b3d0a2783d31b21b8573 |
| pknum | int | 当前条数 | 示例:1 |
| pktotal | int | 总条数 | 示例:2 |
| mobile | string | 手机号:号码规则详见“[3.2 手机号码]规则” |     |
| countrycode | int | 手机号的国际区号:-1 表示无效号码 | 示例:86 |
| spno | string | 完整的通道号 | 示例:955337890 |
| exno | string | 下行时填写的exno | 示例:7890 |
| stime | string | 状态报告对应的下行发送时间:<br><br>YYYY-MM-DD HH:MM:SS | 示例:2016-08-04 17:38:59 |
| rtime | string | 状态报告返回时间:<br><br>YYYY-MM-DD HH:MM:SS | 示例:2016-08-04 17:39:01 |
| status | int | 接收状态:<br><br>0:成功 非0:失败 | 示例:0 |
| errcode | string | 状态报告错误代码 | 示例:DELIVRD |
| errdesc | string | 状态报告错误代码的描述编码方式:<br><br>若鉴权成功,编码方式与下发请求内容编码方式相同,否则为urlencode（GBK明文） | 示例:<br><br>递交成功:success递交失败:欠费 |
| exdata | string | 下行时填写的exdata |     |
| smstype | string | 短信发送类型 0:短信<br><br>8:短转AIM | 示例:0 |
| rpttype | string | smstype为0时,该值固定为1 smstype为8时,<br><br>1:短信通知状态报告<br><br>2:AIM解析状态报告 | 示例:1 |


### Post 请求示例

urlencode

userid=J10003&pwd=26dad7f364507df18f3841cc9c4ff94d&timestamp=0803192020&retsize=300

JSON

{"userid":"J10003","pwd":"26dad7f364507df18f3841cc9c4ff94d","timestam p":"0803192020","retsize":300}

XML
```xml
<?xml version=1.0 encoding=utf-8?>
<rptreq>
<userid>J10003</userid>
<pwd>26dad7f364507df18f3841cc9c4ff94d </pwd>
<timestamp>0803192020</timestamp>
<retsize>300</retsize>
</rptreq>
```

### GET 请求示例

URLENCODE

用户名密码鉴权方式

userid=J10003&pwd=26dad7f364507df18f3841cc9c4ff94d&timestamp=0803192020&retsize=300

apikey 鉴权方式

apikey=abade5589e2798f82f006bbc36d269ce&retsize=300

### 获取成功返回示例

JSON (urlencode 与 JSON 请求都以 JSON 返回数据)
```json
{
"result":0, "desc":%e6%88%90%e5%8a%9f, "rpts": [

{
"msgid":9223372045854775808,
"custid":"b3d0a2783d31b21b8573", "pknum":1,
"pktotal":2, "mobile":"138xxxxxxxx",
"countrycode":86,
"spno":"1000457890006", "exno":"0006",
"stime":"2016-08-04 17:38:55",
"rtime":"2016-08-04 17:38:59",
"status":0, "errcode":"DELIVRD", "errdesc":"success", "exdata":"exdata0002", "smstype":"4",
"rpttype":"2"

},

{
"msgid":9223372045854875809,
"custid":"b3d0a2783d31b21b8574", "pknum":2,
"pktotal":2, "mobile":"138xxxxxxxx", "countrycode":86, "spno":"1000457890006", "exno":"0006",
"stime":"2016-08-04 17:38:55",
"rtime":"2016-08-04 17:38:59",
"status":0, "errcode":"DELIVRD", "errdesc":"success", "exdata":"exdata0002", "smstype":"0",
"rpttype":"1"

}

]

}
```

XML
```xml
<?xml version=1.0 encoding=utf-8?>
<rptrsp>
<result>0</result>
<desc>%e6%88%90%e5%8a%9f</desc>
<rpts>
<rpt>
<msgid>9223372045854775808</msgid>
<custid>b3d0a2783d31b21b8573</custid>
<pknum>1</pknum>
<pktotal>2</pktotal>
<mobile>138xxxxxxxx</mobile>
<countrycode>86</countrycode>
<spno>1000457890006</spno>
<exno>0006</exno>
<stime>2016-08-04 17:38:55</stime>
<rtime>2016-08-04 17:38:59</rtime>
<status>0</status>
<errcode>DELIVRD</errcode>
<errdesc>success</errdesc>
<exdata>exdata0002</exdata>
<smstype>4</smstype>
<rpttype>2</rpttype>
</rpt>
<rpt>
<msgid>9223372045854875809</msgid>
<custid>b3d0a2783d31b21b8574</custid>
<pknum>2</pknum>
<pktotal>2</pktotal>
<mobile>138xxxxxxxx</mobile>
<countrycode>86</countrycode>
<spno>1000457890006</spno>
<exno>0006</exno>
<stime>2016-08-04 17:38:55</stime>
<rtime>2016-08-04 17:38:59</rtime>
<status>0</status>
<errcode>DELIVRD</errcode>
<errdesc>success</errdesc>
<exdata>exdata0002</exdata>
<smstype>0</smstype>
<rpttype>1</rpttype>
</rpt>
</rpts>
</rptrsp>
```

### 获取失败返回示例

JSON (urlencode 与 JSON 请求都以 JSON 返回数据)
```json
{
"result":-100999, "desc":%e7%b3%bb%e7%bb%9f%e5%86%85%e9%83%a8%e9%94%99%e8%af%af, "rpts":""

}
``

XML
```xml
<?xml version=1.0 encoding=utf-8?>
<rptrsp>
<result>100999</result>
<desc>

%e7%b3%bb%e7%bb%9f%e5%86%85%e9%83%a8%e9%94%99%e8%af%af

</desc>
<rpts></rpts>
</rptrsp>
```

## 8. 查询余额接口 get_balance

### 请求URL

<https://ip:port/sms/v2/std/get_balance> <http://ip:port/sms/v2/std/get_balance>

### 请求参数说明

| 参数 | 类型 | 是否必须 | 描述 | 示例 |
| --- | --- | --- | --- | --- |
| userid | string | 是   | 用户账号:长度最大 6 个字符,统一大写,如提交参数中包含 apikey,则可以不用填写该参数及 pwd,两种鉴权方式中只<br><br>能选择一种方式来进行鉴权 | 示例:J10003 |
| pwd | string | 是   | 用户密码:定长小写 32 位字符,如提交参数中包含 apikey,则可以不用填写该参数及 userid,两种鉴权方式中只能选择一种方式来进行鉴权。密码规则详见“[3.1 鉴权规则]” | 示例:<br><br>密码明文模式:111111密码加密模式:<br><br>账号:J10003密码:111111<br><br>固定字符串:00000000时间戳:0803192020<br><br>MD5 加密之前的对应字符串: J1000300000000111111080 3192020<br><br>MD5 加密之后的密码字符串: 26dad7f364507df18f3841c<br><br>c9c4ff94d |
| timestamp | string | 否   | 时间戳:24 小时制格式:MMDDHHMMSS,即月日时<br><br>分秒,定长 10 位,月、日、时、分、秒每段不足 2 位时左补 0,密码选择 MD5 加密方式时必填该参数,密码选择明文方式时则不用填写 | 示例:0803192020 |


### 返回参数说明

| 参数 | 类型 | 是否必须 | 描述 | 示例 |
| --- | --- | --- | --- | --- |
| result | int | 是   | 查询余额请求处理结果:<br><br>0:成功<br><br>非 0:失败,详见 [4] [错误代码]| desc | string | 否   | 应答结果描述,当 result 非 0时,为错误描述<br><br>编码方式:urlencode(UTF-8) | 示例:<br><br>“鉴权失败”<br><br>urlencode(UTF-8)编码:<br><br>%e9%89%b4%e6%9d%83<br><br>%e5%a4%b1%e8%b4%a5 |
| balance | int | 否   | 短信余额条数 |     |


### Post 请求示例

urlencode

userid=J10003&pwd=26dad7f364507df18f3841cc9c4ff94d&timestamp=0803192020

JSON
```json
{"userid":"J10003","pwd":"26dad7f364507df18f3841cc9c4ff94d","timestam p":"0803192020"}
```

XML
```xml
<?xml version=1.0 encoding=utf-8?>
<feereq>
<userid>J10003</userid>
<pwd>26dad7f364507df18f3841cc9c4ff94d</pwd>
<timestamp>0803192020</timestamp>
</feereq>
```

### GET 请求示例

URLENCODE

用户名密码鉴权方式

userid=J10003&pwd=26dad7f364507df18f3841cc9c4ff94d&timestamp=0803192020

apikey 鉴权方式

apikey=abade5589e2798f82f006bbc36d269ce

### 查询成功返回示例

JSON (urlencode 与 JSON 请求都以 JSON 返回数据)
```json
{
"result":0, "desc":%e6%88%90%e5%8a%9f, "balance":99885514,

}
```

XML
```xml
<?xml version=1.0 encoding=utf-8?>
<feersp>
<result>0</result>
<desc>%e6%88%90%e5%8a%9f</desc>
<balance>99885514</balance>
</feersp>
```
### 查询失败返回示例

JSON (urlencode 与 JSON 请求都以 JSON 返回数据)
```json
{
"result":-100999, "desc":%e7%b3%bb%e7%bb%9f%e5%86%85%e9%83%a8%e9%94%99%e8%af%af, "balance":0,

}
```
XML
```xml
<?xml version=1.0 encoding=utf-8?>
<feersp>
<result>-100999</result>
<desc>

%e7%b3%bb%e7%bb%9f%e5%86%85%e9%83%a8%e9%94%99%e8%af%af

</desc>
<balance>0</balance>
</feersp>
```

## 推送上行接口

### 功能说明

开通此接口后,我们将为您实时推送上行。您需要提供一个 url 地址,接收 http post 请求。 该接口一次最多推送 100 个上行,为了不影响推送速度,建议接收到数据后立即回应,用另外线程异步处理业务逻辑。

### 推送方式

POST

### 请求参数说明

| 参数 | 类型 | 是否必须 | 描述 | 示例 |
| --- | --- | --- | --- | --- |
| userid | string | 是   | 用户账号:长度最大6个字节,<br><br>统一大写 | 示例:J10003 |
| pwd | string | 是   | 用户密码:定长小写 32 位字符串。密码规则详见“[3.1 鉴]权规则” | 示例:<br><br>密码明文模式:111111密码加密模式:<br><br>账号:J10003密码:111111<br><br>固定字符串:00000000时间戳:0803192020<br><br>MD5 加密之前的对应字符串: J1000300000000111111080 3192020<br><br>MD5 加密之后的密码字符串: 26dad7f364507df18f3841c<br><br>c9c4ff94d |
| timestamp | string | 否   | 时间戳:24 小时制格式:<br><br>MMDDHHMMSS,即月日时<br><br>分秒,定长 10 位,月、日、时、分、秒每段不足 2 位时左补 0,密码选择 MD5 加密方式时必填该参数,密码选择明文方式<br><br>时则不用填写 | 示例:0803192020 |
| cmd | string | 是   | 推送上行请求命令: 必须填<br><br>MO_REQ | 示例:MO_REQ |
| seqid | int | 是   | 请求消息流水号:匹配回应请<br><br>求的短信包,每次网络请求加 1 | 示例:1003 |
| mos | string | 是   | 上行信息<br><br>格式见下

mos包结构参数说明:

| 参数 | 类型 | 是否必须 | 描述 | 示例 |
| --- | --- | --- | --- | --- |
| msgid | long (64位) | 是   | 平台流水号:上行在梦网云通信平台中的唯一编号 | 示例:<br><br>9223372045854775808<br><br>注意:msgid允许出现负数 |
| mobile | string | 是   | 手机号:号码规则详见“[3.2 手]机号码规则” |     |
| countrycode | int | 是   | 手机号的国际区号:-1 表示无效号码 | 示例:86 |
| pknum | int | 是   | 当前条数 | 示例:1 |
| pktotal | int | 是   | 总条数 | 示例:2 |
| spno | string | 是   | 完整的通道号 | 示例:955337890 |
| exno | string | 是   | 下行时填写的exno | 示例:7890 |
| rtime | string | 是   | 上行返回的时间:<br><br>YYYY-MM-DD HH:MM:SS | 示例:<br><br>2016-08-04 17:38:59 |
| content | string | 是   | 短信内容:最大支持 70 个字 (含签名),一个字母或一个汉字都视为一个字。本字段支持加密传输,若不使用加密功能编码方式:<br><br>UrlEncode("验证码:6666,打死 都 不 要 告 诉 别 人 哦 ！ ","UTF-8")<br><br>若使用加密功能详见“[3.4 内]容加密规则” | 示例:<br><br>短信内容:“验证码:6666,打死都不要告诉别人哦！”短信内容进行urlencode 编码后:<br><br>%e9%aa%8c%e8%af%81<br><br>%e7%a0%81%ef%bc%9a6 666%ef%bc%8c%e6%89%93%e6%ad%bb%e9%83% bd%e4%b8%8d%e8%a6%81%e5%91%8a%e8%af%8<br><br>9%e5%88%ab%e4%ba%ba<br><br>%e5%93%a6%ef%bc%81 |


### 返回参数说明

| 参数 | 类型 | 是否必须 | 描述 | 示例 |
| --- | --- | --- | --- | --- |
| cmd | string | 是   | 必须填MO_RESP |     |
| seqid | int | 是   | 与请求中的 seqid 保持一致 |     |
| result | int | 是   | 上行短消息请求处理结果:<br><br>0:成功<br><br>非 0:失败,详见 [4] [错误代码]

## 9. Post 请求示例

userid=J10003&pwd=26dad7f364507df18f3841cc9c4ff94d&timestamp=08031920 20&cmd=MO_REQ&seqid=1003&mos=[{"msgid":9223372045854775808,"mobile":" 138xxxxxxxx","countrycode":86,"spno":"1000457890006","exno":"0006","r time":"2016-08-04 17:38:59","content":"%e9%aa%8c%e8%af%81%e7%a0%81%ef%bc%9a6666%ef%bc%8 c%e6%89%93%e6%ad%bb%e9%83%bd%e4%b8%8d%e8%a6%81%e5%91%8a%e8%af%89%e5%8 8%ab%e4%ba%ba%e5%93%a6%ef%bc%81"},{"msgid":9223372045854895808,"mobil e":"131xxxxxxxx","countrycode":86,"spno":"1000357890006","exno":"0006 ","rtime":"2016-08-04 17:39:50","content":"%e9%aa%8c%e8%af%81%e7%a0%81%ef%bc%9a6666%ef %bc%8c%e6%89%93%e6%ad%bb%e9%83%bd%e4%b8%8d%e8%a6%81%e5%91%8a%e8%af%89%e5%88%ab%e4%ba%ba%e5%93%a6%ef%bc%81"}]

## 10. 推送成功返回示例
```json
{
"cmd":"MO_RESP",
"seqid":1003, "result":0

}
```

## 11. 推送失败返回示例

```json
{
"cmd":"MO_RESP",
"seqid":1003, "result":-100999

}
```

## 推送状态报告接口

### 功能说明

开通此接口后,我们将为您实时推送状态报告。您需要提供一个 url 地址,接受 http post请求。 该接口一次最多推送 500 个状态报告,为了不影响推送速度,建议接收到数据后立即回应,另外线程异步处理业务逻辑。

### 推送方式

POST

### 请求参数说明

| 参数 | 类型 | 是否必须 | 描述 | 示例 |
| --- | --- | --- | --- | --- |
| userid | string | 是   | 用户账号:长度最大6个字节,<br><br>统一大写 | 示例:J10003 |
| pwd | string | 是   | 用户密码:定长小写 32 位字符串。密码规则详见“[3.1 鉴]权规则” | 示例:<br><br>密码明文模式:111111密码加密模式:<br><br>账号:J10003密码:111111<br><br>固定字符串:00000000时间戳:0803192020<br><br>MD5 加密之前的对应字符串: J1000300000000111111080 3192020<br><br>MD5 加密之后的密码字符串: 26dad7f364507df18f3841c<br><br>c9c4ff94d |
| timestamp | string | 否   | 时间戳:24 小时制格式:<br><br>MMDDHHMMSS,即月日时<br><br>分秒,定长 10 位,月、日、时、分、秒每段不足 2 位时左补 0,密码选择 MD5 加密方式时必填该参数,密码选择明文方式<br><br>时则不用填写 | 示例:0803192020 |
| cmd | string | 是   | 推送状态报告请求命令:必须<br><br>填 RPT_REQ | 示例:RPT_REQ |
| seqid | int | 是   | 请求消息流水号:匹配回应请<br><br>求的短信包,每次网络请求加 1 | 示例:1004 |
| rpts | string | 是   | 状态报告。<br><br>格式见下

rpts包结构参数说明:

| 参数 | 类型 | 描述 | 示例 |
| --- | --- | --- | --- |
| msgid | long (64 位) | 平台流水号:对应下行请求返回结果中的 msgid | 示例:<br><br>9223372036854775808<br><br>注意:msgid 允许出现负数 |
| custid | string | 用户自定义流水号:对应下行请求时<br><br>填写的 custid | 示例:<br><br>b3d0a2783d31b21b8573 |
| pknum | int | 当前条数 | 示例:1 |
| pktotal | int | 总条数 | 示例:2 |
| mobile | string | 手机号:号码规则详见“[3.2 手机号码]<br><br>规则” |     |
| countrycode | int | 手机号的国际区号:-1 表示无效号码 | 示例:86 |
| spno | string | 完整的通道号 | 示例:955337890 |
| exno | string | 下行时填写的exno | 示例:7890 |
| stime | string | 状态报告对应的下行发送时间:<br><br>YYYY-MM-DD HH:MM:SS | 示例:2016-08-04 17:38:59 |
| rtime | string | 状态报告返回时间:<br><br>YYYY-MM-DD HH:MM:SS | 示例:2016-08-04 17:39:01 |
| status | int | 接收状态:<br><br>0:成功 非0:失败 | 示例:0 |
| errcode | string | 状态报告错误代码 | 示例:DELIVRD |
| errdesc | string | 状态报告错误代码的描述编码方式:<br><br>根据不同报文格式编码: XML/x-www-form-urlencoded: URELENCODE(UTF-8)<br><br>JSON:UTF-8 | 示例:<br><br>递交成功:success递交失败:欠费 |
| exdata | string | 下行时填写的exdata |     |
| smstype | string | 短信类型 0:短信<br><br>4:短转富信<br><br>8:短转AIM | 示例:4 |
| rpttype | string | 状态报告类型 1:通知状态报告<br><br>2:下载状态报告<br><br>当短信类型为4(短转富信)或者8(短转AIM)才会有下载状态报告 | 示例:2 |


### 返回参数说明

| 参数 | 类型 | 是否必须 | 描述 | 示例 |
| --- | --- | --- | --- | --- |
| cmd | string | 是   | 必须填RPT_RESP |     |
| seqid | int | 是   | 与请求中的 seqid 保持一致 |     |
| result | int | 是   | 状态报告请求处理结果:<br><br>0:成功<br><br>非 0:失败,详见 [4] [错误代码]

## 12. Post 请求示例

userid=J10003&pwd=26dad7f364507df18f3841cc9c4ff94d&timestamp=08031920 20&cmd=RPT_REQ&seqid=1004&rpts=[{"msgid":9223372045854775808,"custid": "b3d0a2783d31b21b8573","pknum":1,"pktotal":2,"mobile":"138xxxxxxxx"," countrycode":86,"spno":"1000457890006","exno":"0006","stime":"2016-08 04 17:38:55","rtime":"2016-08-04 17:38:59","status":0,"errcode":"DELIVRD","errdesc":"success","exdata": "exdata0002","smstype":"4","rpttype":"2"},{"msgid":9223372045854875808," custid":"b3d0a2783d31b21b8573","pknum":2,"pktotal":2,"mobile":"138xxx xxxxx","spno":"1000457890006","exno":"0006","stime":"2016-08-04 17:38:55","rtime":"2016-08-04 17:38:59","status":0, "errcode":"DELIVRD","errdesc":"success","exdata":"exdata0002","smstype":"4","rpttype":"2"}]

## 13. 推送成功返回示例

```json
{
"cmd":"RPT_RESP",
"seqid":1004, "result":0

}
```

## 14. 推送失败返回示例

```json
{
"cmd":"RPT_RESP",
"seqid":1004, "result": 100999

}
```

# 规则说明

## 鉴权规则

1. 账号+密码的密文值进行进行用户鉴权 (账号默认鉴权模式):

userid 填写账号明文,pwd 的密码加密方式如下:将 userid 值大写、固定字符串 00000000、明文 pwd、timestamp 依次拼接成字符串后,再进行 MD5 加密,userid 和明文的 pwd 在梦网开户时提供,timestamp 为时间戳,24 小时制,格式为: MMDDHHMMSS,即月日时分秒,定长 10 位,月、日、时、分、秒每段不足 2 位时左补 0。

备注:时间戳请填写发送时动态的准确时间,不要填写固定值,否则时间误差太大服务器有可能拒绝您的请求。

1. 账号+密码的明文值进行进行用户鉴权(SHA256 方式,需申请):

userid 填写账号明文,pwd 的密码加密方式如下:将 userid 值大写、固定字符串 00000000、明文 pwd、timestamp 依次拼接成字符串后,再进行 SHA256 加密后转换成十六进制小写字符串,userid 和明文的 pwd 在梦网开户时提供,timestamp 为时间戳, 24 小时制,格式为:MMDDHHMMSS,即月日时分秒,定长 10 位,月、日、时、分、秒每段不足 2 位时左补 0。

备注:时间戳请填写发送时动态的准确时间,不要填写固定值,否则时间误差太大服务器有可能拒绝您的请求。

1. 账号+密码的明文值进行进行用户鉴权(SM3 方式,需申请):

userid 填写账号明文,pwd 的密码加密方式如下:将 userid 值大写、固定字符串

00000000、明文 pwd、timestamp 依次拼接成字符串后,再进行 SM3 加密后转换成十六进制小写字符串,userid 和明文的 pwd 在梦网开户时提供,timestamp 为时间戳,24 小时制,格式为:MMDDHHMMSS,即月日时分秒,定长 10 位,月、日、时、分、秒每段不足 2 位时左补 0。

备注:时间戳请填写发送时动态的准确时间,不要填写固定值,否则时间误差太大服务器有可能拒绝您的请求。

1. 账号+密码的明文值进行进行用户鉴权(需申请):

userid 及 pwd 都填写明文,如需使用此种鉴权方式,可联系我司技术支撑人员申请密码明文鉴权。

1. apikey 进行用户鉴权(需申请):

apikey 为 32 位长度,使用此种鉴权方式时,则 userid 及 pwd 无效不可用,如需使用此种鉴权方式,可联系我司技术支撑人员申请 apikey。

apikey 取值示例:abade5589e2798f82f006bbc36d269ce

## 手机号码规则

1. 每次请求所提交的号码段类型不受限制,但系统会对每个手机号码做合法性验证,若有 不合法的手机号码,以失败状态报告的形式返回。号码段类型分为:移动、联通、电信手机。
2. 若只发送国外的号码,那么号码规则必须是:00+国际电话区号+手机号码。
3. 若只发送中国的号码,号码前面无需加国际电话区号(0086)。
4. 发送的号码当中既有中国也有国外,那么号码规则必须是:00+国家电话区号+手机号码,如:香港号码564xxxxx,中国号码132xxxxxxxx,那么发送时应该填00852564xxxxx, 0086132xxxxxxxx

## 匹配状态报告规则

1. 单条发送接口

custid 匹配方法:

短短信时使用下行请求包里的 custid+mobile 进行匹配；

长短信时使用下行请求包里的 custid+mobile+pknum 进行匹配； msgid 匹配方法:

短短信时使用返回包里的 msgid 进行匹配；

长短信时使用下述公式计算手机号对应的长短信分条的流水号: msgid+(当前分条数-1)\*17179869184；

1. 相同内容群发接口

custid 匹配方法:

短短信时使用下行请求包里的 custid+mobile 进行匹配；

长短信时使用下行请求包里的 custid+mobile+pknum 进行匹配； msgid 匹配方法:

短短信时使用返回包里的 msgid 进行匹配,但每个手机号码对应的 msgid 需要按下面公式进行计算:msgid+手机号码位置-1；

长短信时使用下述公式计算每个手机号对应的长短信分条的流水号: (msgid+(手机号位置-1))+(当前分条数-1)\*17179869184；

1. 个性化群发接口

custid 匹配方法:

短短信时使用请求包里的 multimt 中的 custid+mobile 进行匹配；

长短信时使用请求包里的 multimt 中的 custid+mobile+pknum 进行匹配；

msgid 匹配方法:

短短信时使用返回包里的 msgid 进行匹配,但每个手机号码对应的 msgid 需要按下面公式进行计算:msgid+手机号码位置-1;

长短信时使用下述公式计算每个手机号对应的长短信分条的流水号: (msgid+(手机号位置-1))+(当前分条数-1)\*17179869184；

## 内容加密规则

1\. SM4-ECB 加密算法规则

加密前内容编码方式:send_template 接口为 P1=UrlEncode("梦网科技 ","UTF-8")&P2=UrlEncode("201701","UTF-8"),"UTF-8")；send_mixed 接口为

UrlEncode("验证码:6666,打死都不要告诉别人哦！","UTF-8")+","+UrlEncode ("验证码:8888,打死都不要告诉别人哦！","UTF-8")；其余接口均为 UTF-8。

待加密内容字节数不能被 16 整除时,需要对内容进行填充,支持两种填充方式: 1）ZEROPADDING:在末尾填充（16-内容字节数%16）个字符 0x00,若 16-内容字

节数%16 等于 0,即内容长度为 16 字节的整数倍,则无需填充。

2）PKCS7PADDING:在末尾填充 16-内容字节数%16 个（16-内容字节数%16）ASCII码值对应的字符。

支持两种密钥值:

1. 固定值:由双方共同约定好一个固定不变长度为 16 字节的密钥,客户端使用该密钥加密,服务端使用该密钥解密。
2. SM3 哈希值:SM3(pwd+timestamp)后的哈希值前 16 字节,其中 pwd 为账号对应的明文密码,timestamp 为请求报文中时间戳字段。

加密后需再对加密结果进行十六进制转换。

# 错误代码表

| 返回错误代码 | 错误说明 |
| --- | --- |
| 100001 | 鉴权不通过,请检查账号,密码,时间戳,固定串,以及MD5<br><br>算法是否按照文档要求进行设置 |
| 100002 | 用户多次鉴权不通过,请检查帐号,密码,时间戳,固定<br><br>串,以及MD5算法是否按照文档要求进行设置 |
| 100003 | 用户欠费 |
| 100004 | custid或者exdata字段填写不合法 |
| 100011 | 短信内容超长 |
| 100012 | 手机号码不合法 |
| 100014 | 手机号码超过最大支持数量（1000） |
| 100029 | 端口绑定失败 |
| 100056 | 用户账号登录的连接数超限 |
| 100057 | 用户账号登录的 IP 错误 |
| 100058 | 模板 ID 不存在 |
| 100059 | 模板参数个数不匹配 |
| 100060 | 手机号与信息内容个数不一致 |
| 100070 | 没有发送该接口的权限 |
| 100091 | XML 报文解析异常(不符合 XML 格式) |
| 100092 | 报文解析异常 |
| 100093 | Json 报文解析异常 |
| 100094 | XML 节点解析失败(节点不存在) |
| 100095 | Json 报文解析异常 |
| 100096 | Json 数组对象解析错误 |
| 100098 | Json 报文解析异常 |
| 100100 | 不支持的 Content-type 格式 |
| 100101 | sign 不合法 |
| 100102 | 无有效查询账号或查询的账号数量超限 |
| 100103 | 查询账号数量超出限制 |
| 100106 | 短信内容字段不存在或为空 |
| 100252 | SVRTYPE 参数含非法字符 |
| 100253 | 不满足模板修改条件(仅审核不通过的模板才可以修改) |
| 100254 | 没有权限操作该模板(模板不属于该企业) |
| 100255 | 不满足模板启用条件(内容安全原因被禁用、审核中、审<br><br>核不通过、已过期的模板不允许启用) |
| 100256 | 企业模板数量已超出限制 |
| 100257 | 模板管理数量超出限制(一次最多可管理 1000 个模板) |
| 100259 | 模板有效期不合法 |
| 100260 | 模板内容超长 |
| 100261 | 模板名称不合法 |
| 100262 | 模板类型不合法 |
| 100264 | 模板操作类型不合法 |
| 100265 | 模板参数不合法 |
| 100266 | 当天模板启禁用次数超过 100 次 |
| 100267 | 下行内容解密失败 |
| 100502 | 短信模板被禁用 |
| 100800 | 短地址相关接口失败 |
| 100801 | 短地址长地址错误 |
| 100802 | 长地址已被禁用 |
| 100803 | 短地址短域名错误 |
| 100804 | 短地址短域名已经被禁用 |
| 100805 | 短地址手机号码超过 10000 个 |
| 100806 | 短地址中心资源不够 |
| 100807 | 短地址服务器忙 |
| 100900 | 号码查询号码手机个数错误 |
| 100901 | 号码查询获取手机号码信息失败 |
| 100902 | 号码查询扣费请求失败 |
| 100907 | 号码查询服务器忙 |
| 100999 | 平台数据库内部错误 |

# 注意事项

1. 请求报文支持 JSON,XML,x-www-form-urlencoded,用 Content-type 识别请求包报文的格式,如使用 x-www-form-urlencoded 时,格式如下:

POST /sms/v2/std/single_send HTTP/1.1 Host: 192.168.1.1

Content-Type: application/x-www-form-urlencoded Content-Length: length

1. 为确保能够快速及时获取到上行或状态报告,请在调用接口后判断接口是否有上行或状

态报告返回,若有返回,则需要一直获取,直到接口返回无数据时,延时 5 秒,然后再次重复前面的获取和判断操作。

1. 单条发送接口每次只能发送一个号码,否则接口返回失败。
2. 单个发送帐号默认并发链接数为 100 个,如需调整增加并发连接数,请联系技术支撑人员。
3. “推送上行”及“推送状态报告”默认不开启,如果需开启,请联系技术支撑人员。
4. 在“状态报告”的包结构“rpts”中,若是短短信,“pknum”与“pktotal”两个字段的值都为 1。
5. 对应本接口说明文档的的示例代码,详见同级目录下的示例代码文件。
6. 本接口基于 http 协议, 适用于任何支持 http 协议的编程语言进行开发,如: C++, JAVA,C#,PHP,Python,NodeJS,R,JavaScript,Ruby,Go,Swift 等其他的编程语言。
7. 所有接口均支持 HTTP 和 HTTPS 方式调用,如对信息安全要求较高的用户建议使用 HTTPS 加密方式接入,增加数据在网络传输过程中的安全性。