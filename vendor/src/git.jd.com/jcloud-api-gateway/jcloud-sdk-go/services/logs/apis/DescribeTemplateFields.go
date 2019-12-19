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

type DescribeTemplateFieldsRequest struct {

    core.JDCloudRequest

    /* 地域 Id  */
    RegionId string `json:"regionId"`

    /* 日志主题 UID  */
    LogtopicUID string `json:"logtopicUID"`
}

/*
 * param regionId: 地域 Id (Required)
 * param logtopicUID: 日志主题 UID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeTemplateFieldsRequest(
    regionId string,
    logtopicUID string,
) *DescribeTemplateFieldsRequest {

	return &DescribeTemplateFieldsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/logtopics/{logtopicUID}/fields",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        LogtopicUID: logtopicUID,
	}
}

/*
 * param regionId: 地域 Id (Required)
 * param logtopicUID: 日志主题 UID (Required)
 */
func NewDescribeTemplateFieldsRequestWithAllParams(
    regionId string,
    logtopicUID string,
) *DescribeTemplateFieldsRequest {

    return &DescribeTemplateFieldsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/logtopics/{logtopicUID}/fields",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        LogtopicUID: logtopicUID,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeTemplateFieldsRequestWithoutParam() *DescribeTemplateFieldsRequest {

    return &DescribeTemplateFieldsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/logtopics/{logtopicUID}/fields",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 Id(Required) */
func (r *DescribeTemplateFieldsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param logtopicUID: 日志主题 UID(Required) */
func (r *DescribeTemplateFieldsRequest) SetLogtopicUID(logtopicUID string) {
    r.LogtopicUID = logtopicUID
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeTemplateFieldsRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeTemplateFieldsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeTemplateFieldsResult `json:"result"`
}

type DescribeTemplateFieldsResult struct {
    Fields []logs.Field `json:"fields"`
}