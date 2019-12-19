package common

import (
	"git.jd.com/jcloud-api-gateway/jcloud-sdk-go/core"
	functionClient "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/function/client"
	logsClient "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/logs/client"
)

type NoLogger struct{}

func NewNoLogger() *NoLogger {
	return &NoLogger{}
}

func (logger NoLogger) Log(level int, message ...interface{}) {}

func NewFunctionClient(user *User) *functionClient.FunctionClient {
	credential := core.NewCredential(user.AccessKey, user.SecretKey)
	config := core.NewConfig()
	config.SetScheme(core.SchemeHttp)
	functionClient := functionClient.NewFunctionClient(credential)
	functionClient.SetConfig(config)
	functionClient.SetLogger(logger.NewNoLogger())
	return functionClient
}

func NewLogClient(user *User) *logsClient.LogsClient {
	credential := core.NewCredential(user.AccessKey, user.SecretKey)
	config := core.NewConfig()
	config.SetScheme(core.SchemeHttp)
	client := logsClient.NewLogsClient(credential)
	client.SetConfig(config)
	client.SetLogger(logger.NewNoLogger())
	return client
}