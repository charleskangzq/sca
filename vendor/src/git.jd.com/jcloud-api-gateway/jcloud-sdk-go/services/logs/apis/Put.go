// Copyright 2018 JDCLOUD.COM
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// NOTE: This class is auto generated by the jdcloud code generator program.

package apis

import (
    "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/core"
    logs "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/logs/models"
)

type PutRequest struct {

    core.JDCloudRequest

    /* 日志主题uid  */
    LogtopicUID string `json:"logtopicUID"`

    /* 全局 strean 日志流标识符（建议起能唯一界定一个文件的名字，如 /i-iqnvqpinkjiq/app.log），不传则写入default日志流中（会导致很多文件混合在一起，不推荐） (Optional) */
    Stream *string `json:"stream"`

    /* 全局时间戳，UTC格式，最多支持到纳秒级别，不传入则取服务器时间。如 2019-04-08T03:08:04.437670934Z、2019-04-08T03:08:04Z、2019-04-08T03:08:04.123Z (Optional) */
    Timestamp *string `json:"timestamp"`

    /* 日志数据  */
    Entries []logs.Entry `json:"entries"`
}

/*
 * param logtopicUID: 日志主题uid (Required)
 * param entries: 日志数据 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewPutRequest(
    logtopicUID string,
    entries []logs.Entry,
) *PutRequest {

	return &PutRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/logtopics/{logtopicUID}:put",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        LogtopicUID: logtopicUID,
        Entries: entries,
	}
}

/*
 * param logtopicUID: 日志主题uid (Required)
 * param stream: 全局 strean 日志流标识符（建议起能唯一界定一个文件的名字，如 /i-iqnvqpinkjiq/app.log），不传则写入default日志流中（会导致很多文件混合在一起，不推荐） (Optional)
 * param timestamp: 全局时间戳，UTC格式，最多支持到纳秒级别，不传入则取服务器时间。如 2019-04-08T03:08:04.437670934Z、2019-04-08T03:08:04Z、2019-04-08T03:08:04.123Z (Optional)
 * param entries: 日志数据 (Required)
 */
func NewPutRequestWithAllParams(
    logtopicUID string,
    stream *string,
    timestamp *string,
    entries []logs.Entry,
) *PutRequest {

    return &PutRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/logtopics/{logtopicUID}:put",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        LogtopicUID: logtopicUID,
        Stream: stream,
        Timestamp: timestamp,
        Entries: entries,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewPutRequestWithoutParam() *PutRequest {

    return &PutRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/logtopics/{logtopicUID}:put",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param logtopicUID: 日志主题uid(Required) */
func (r *PutRequest) SetLogtopicUID(logtopicUID string) {
    r.LogtopicUID = logtopicUID
}

/* param stream: 全局 strean 日志流标识符（建议起能唯一界定一个文件的名字，如 /i-iqnvqpinkjiq/app.log），不传则写入default日志流中（会导致很多文件混合在一起，不推荐）(Optional) */
func (r *PutRequest) SetStream(stream string) {
    r.Stream = &stream
}

/* param timestamp: 全局时间戳，UTC格式，最多支持到纳秒级别，不传入则取服务器时间。如 2019-04-08T03:08:04.437670934Z、2019-04-08T03:08:04Z、2019-04-08T03:08:04.123Z(Optional) */
func (r *PutRequest) SetTimestamp(timestamp string) {
    r.Timestamp = &timestamp
}

/* param entries: 日志数据(Required) */
func (r *PutRequest) SetEntries(entries []logs.Entry) {
    r.Entries = entries
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r PutRequest) GetRegionId() string {
    return ""
}

type PutResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result PutResult `json:"result"`
}

type PutResult struct {
}