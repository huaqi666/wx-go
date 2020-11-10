## wx-go - 微信开发 Golang SDK（开发工具包） 

[![travis-image]][travis-url]
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

本项目从 [WxJava](https://github.com/Wechat-Group/WxJava) 转化而来。

#### 支持包括微信支付、公众号、小程序等微信功能的后端开发。

### 使用方式
注意：最新版本（包括测试版）为 [![GitHub release](https://img.shields.io/badge/github-releases-blue)](https://github.com/cliod/wx-go/releases)

```shell script
go get -u github.com/cliod/wx-go
```

  - 微信小程序：`ma`   
  - 微信支付：`pay`
  - 公众号（包括订阅号和服务号）：`mp`    

### 使用Demo
公众号使用例子
```go
conf := mp.NewWxMpConfig("<appId>", "<secret>")
service := mp.NewWxMpService(conf)

da, err := service.CreateJsapiSignature("https://www.xxx.com")
if err == nil {
    // todo ...
} else {
    fmt.Println(err.Error())
}
```

小程序使用例子
```go
conf := mp.NewWxMaConfig("<appId>", "<secret>")
service := ma.NewWxMaService(conf)

uc := service.GetWxMaUserService()

res, err := uc.GetSessionInfo("<js_code>")
if err == nil {
    // todo ...
} else {
    fmt.Println(err.Error())
}
```

微信支付使用例子
```go
conf := mp.NewWxPayV2Config("<appId>", "<secret>", "<mchId>", "mchKey", "<notifyUrl>", "keyPath")
service := ma.NewWxPayService(conf)

res, err := service.UnifyPay(&pay.WxPayUnifiedOrderRequest{
    TotalFee:   100,
    Openid:     c.Openid,
    OutTradeNo: s,
    Body:       "测试数据",
})
if err != nil {
    fmt.Println(err.Error())
} else {
    // todo...
}
```

[travis-image]: https://api.travis-ci.com/cliod/wx-go.svg?branch=main
[travis-url]: https://travis-ci.com/cliod/wx-go
