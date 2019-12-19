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

type DescribeShippersRequest struct {

    core.JDCloudRequest

    /* 地域 Id  */
    RegionId string `json:"regionId"`

    /* 日志集 UID  */
    LogsetUID string `json:"logsetUID"`

    /* 日志主题 UID  */
    LogtopicUID string `json:"logtopicUID"`

    /* 当前所在页，默认为1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 页面大小，默认为20；取值范围[1, 100] (Optional) */
    PageSize *int `json:"pageSize"`
}

/*
 * param regionId: 地域 Id (Required)
 * param logsetUID: 日志集 UID (Required)
 * param logtopicUID: 日志主题 UID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeShippersRequest(
    regionId string,
    logsetUID string,
    logtopicUID string,
) *DescribeShippersRequest {

	return &DescribeShippersRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/logsets/{logsetUID}/logtopics/{logtopicUID}/shippers",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        LogsetUID: logsetUID,
        LogtopicUID: logtopicUID,
	}
}

/*
 * param regionId: 地域 Id (Required)
 * param logsetUID: 日志集 UID (Required)
 * param logtopicUID: 日志主题 UID (Required)
 * param pageNumber: 当前所在页，默认为1 (Optional)
 * param pageSize: 页面大小，默认为20；取值范围[1, 100] (Optional)
 */
func NewDescribeShippersRequestWithAllParams(
    regionId string,
    logsetUID string,
    logtopicUID string,
    pageNumber *int,
    pageSize *int,
) *DescribeShippersRequest {

    return &DescribeShippersRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/logsets/{logsetUID}/logtopics/{logtopicUID}/shippers",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        LogsetUID: logsetUID,
        LogtopicUID: logtopicUID,
        PageNumber: pageNumber,
        PageSize: pageSize,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeShippersRequestWithoutParam() *DescribeShippersRequest {

    return &DescribeShippersRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/logsets/{logsetUID}/logtopics/{logtopicUID}/shippers",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 Id(Required) */
func (r *DescribeShippersRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param logsetUID: 日志集 UID(Required) */
func (r *DescribeShippersRequest) SetLogsetUID(logsetUID string) {
    r.LogsetUID = logsetUID
}

/* param logtopicUID: 日志主题 UID(Required) */
func (r *DescribeShippersRequest) SetLogtopicUID(logtopicUID string) {
    r.LogtopicUID = logtopicUID
}

/* param pageNumber: 当前所在页，默认为1(Optional) */
func (r *DescribeShippersRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 页面大小，默认为20；取值范围[1, 100](Optional) */
func (r *DescribeShippersRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeShippersRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeShippersResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeShippersResult `json:"result"`
}

type DescribeShippersResult struct {
    NumberPages int64 `json:"numberPages"`
    NumberRecords int64 `json:"numberRecords"`
    PageNumber int64 `json:"pageNumber"`
    PageSize int64 `json:"pageSize"`
    ShipperResult []logs.ShipperEnd `json:"shipperResult"`
}