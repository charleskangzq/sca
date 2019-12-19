# 京东云Golang SDK
欢迎使用京东云开发者Golang工具套件（Go SDK）。使用京东云Go SDK，您无需复杂编程就可以访问京东云提供的各种服务。
为了方便您理解SDK中的一些概念和参数的含义，使用SDK前建议您先查看OpenAPI使用入门。

## 环境准备
1.	京东云Go SDK适用于Go 1.6及以上版本。
2.	在开始调用京东云open API之前，需提前在京东云用户中心账户管理下的[AccessKey管理页面](https://uc.jdcloud.com/accesskey/index)申请accesskey和secretKey密钥对（简称AK/SK）。AK/SK信息请妥善保管，如果遗失可能会造成非法用户使用此信息操作您在云上的资源，给你造成数据和财产损失。

## 下载和安装
在项目路径下(GOPATH)执行以下命令：
```
GOPATH=$(PWD)
mkdir -p src/git.jd.com/jcloud-api-gateway/jcloud-sdk-go
cd src/git.jd.com/jcloud-api-gateway/jcloud-sdk-go
go get -insecure git.jd.com/jcloud-api-gateway/jcloud-sdk-go/core github.com/satori/go.uuid
```

或手工下载压缩包，最终目录结构为：
`$GOPATH/src/git.jd.com/jcloud-api-gateway/jcloud-sdk-go`

## 调用SDK
### 业务侧SDK的调用主要分为4步：
1.	设置accessKey和secretKey
2.	创建业务Client
3.	设置请求参数
4.	执行请求得到响应

### 大致代码如下：
``` go
accessKey := "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
secretKey := "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
credentials := NewCredentials(accessKey, secretKey)
client := NewVmClient(credentials)

req := NewDescribeInstancesRequest("cn-north-1")
resp, err := client.DescribeInstances(req)
fmt.Println(resp.Result.Instances)
fmt.Println(resp.Result.TotalCount)
fmt.Println(len(resp.Result.Instances))
```
请参考demo中的测试用例，访问京东云各业务线接口。