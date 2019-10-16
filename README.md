# 京东云Golang SDK

[![Build Status](https://travis-ci.org/jdcloud-api/jdcloud-sdk-go.svg?branch=master)](https://travis-ci.org/jdcloud-api/jdcloud-sdk-go)
[![GoDoc](https://godoc.org/github.com/jdcloud-api/jdcloud-sdk-go?status.svg)](https://godoc.org/github.com/jdcloud-api/jdcloud-sdk-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/jdcloud-api/jdcloud-sdk-go)](https://goreportcard.com/report/github.com/jdcloud-api/jdcloud-sdk-go)

欢迎使用京东云开发者Golang工具套件（Go SDK）。使用京东云Go SDK，您无需复杂编程就可以访问京东云提供的各种服务。
为了方便您理解SDK中的一些概念和参数的含义，使用SDK前建议您先查看OpenAPI使用入门。

## 环境准备
1.	京东云Go SDK适用于Go 1.6及以上版本。
2.	在开始调用京东云open API之前，需提前在京东云用户中心账户管理下的[AccessKey管理页面](https://uc.jdcloud.com/accesskey/index)申请accesskey和secretKey密钥对（简称AK/SK）。AK/SK信息请妥善保管，如果遗失可能会造成非法用户使用此信息操作您在云上的资源，给你造成数据和财产损失。

## 下载和安装
1.	京东云Go SDK的下载地址：https://github.com/jdcloud-api/jdcloud-sdk-go 。
2.	您也可以使用以下命令获取安装包，代码会被下载到GOPATH环境变量中第一个路径src目录中。

    `go get github.com/jdcloud-api/jdcloud-sdk-go/core github.com/gofrs/uuid`

## 调用SDK
### 业务侧SDK的调用主要分为4步：
1.	设置accessKey和secretKey
2.	创建业务Client
3.	设置请求参数
4.	执行请求得到响应

### 大致代码如下：
``` go
package main

import (
	"fmt"
  	. "github.com/jdcloud-api/jdcloud-sdk-go/services/vm/apis"
	. "github.com/jdcloud-api/jdcloud-sdk-go/services/vm/client"
	. "github.com/jdcloud-api/jdcloud-sdk-go/core"
)

func main() {
	accessKey := "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	secretKey := "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	credentials := NewCredentials(accessKey, secretKey)
	client := NewVmClient(credentials)

	req := NewDescribeInstancesRequest("cn-north-1")
	resp, err := client.DescribeInstances(req)
	if err != nil {
		return
	}
	fmt.Println(resp.Result.Instances)
	fmt.Println(resp.Result.TotalCount)
	fmt.Println(len(resp.Result.Instances))
}
```

**注意：**
- 每支云产品都有自己的Client，当调用该产品API时，需使用该产品的Client。例如：使用NewVmClient创建的client只能调用云主机（Vm）的接口；使用用NewAgClient创建的client只能调用高可用组（Ag）的接口。

如果需要设置额外的header，例如要调用开启了MFA操作保护的接口，需要传递x-jdcloud-security-token，则按照如下方式：
```
const HeaderSecurityToken = "x-jdcloud-security-token"
req := NewDeleteInstanceRequest("cn-north-1", "i-xxxxx")
req.AddHeader(HeaderSecurityToken, "xxx")
resp, err := client.DeleteInstance(req)
```

如果需要设置访问点，配置超时等，请参考如下更复杂的例子：
```
	config := NewConfig()
	config.SetEndpoint("vm.internal.cn-north-1.jdcloud-api.com") //指定非默认访问点，如供VPC专用调用的域名
	config.SetScheme(SchemeHttp) //设置使用HTTP而不是HTTPS，vpc专用域名不支持HTTPS
	config.SetTimeout(20 * time.Second) //设置请求超时
	client.SetConfig(config)
```

如果需要关闭日志输出，则按照如下方式：
```
	client.DisableLogger()
```

