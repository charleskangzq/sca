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

type DescribeUserPermissionsRequest struct {

    core.JDCloudRequest

    /* 当前所在页，默认为1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 页面大小，默认为20；取值范围[1, 100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* erp (Optional) */
    Erp *string `json:"erp"`

    /* serviceCode (Optional) */
    ServiceCode *string `json:"serviceCode"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeUserPermissionsRequest(
) *DescribeUserPermissionsRequest {

	return &DescribeUserPermissionsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/permissionsx",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param pageNumber: 当前所在页，默认为1 (Optional)
 * param pageSize: 页面大小，默认为20；取值范围[1, 100] (Optional)
 * param erp: erp (Optional)
 * param serviceCode: serviceCode (Optional)
 */
func NewDescribeUserPermissionsRequestWithAllParams(
    pageNumber *int,
    pageSize *int,
    erp *string,
    serviceCode *string,
) *DescribeUserPermissionsRequest {

    return &DescribeUserPermissionsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/permissionsx",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        PageNumber: pageNumber,
        PageSize: pageSize,
        Erp: erp,
        ServiceCode: serviceCode,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeUserPermissionsRequestWithoutParam() *DescribeUserPermissionsRequest {

    return &DescribeUserPermissionsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/permissionsx",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param pageNumber: 当前所在页，默认为1(Optional) */
func (r *DescribeUserPermissionsRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 页面大小，默认为20；取值范围[1, 100](Optional) */
func (r *DescribeUserPermissionsRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param erp: erp(Optional) */
func (r *DescribeUserPermissionsRequest) SetErp(erp string) {
    r.Erp = &erp
}

/* param serviceCode: serviceCode(Optional) */
func (r *DescribeUserPermissionsRequest) SetServiceCode(serviceCode string) {
    r.ServiceCode = &serviceCode
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeUserPermissionsRequest) GetRegionId() string {
    return ""
}

type DescribeUserPermissionsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeUserPermissionsResult `json:"result"`
}

type DescribeUserPermissionsResult struct {
    NumberPages int64 `json:"numberPages"`
    NumberRecords int64 `json:"numberRecords"`
    PageNumber int64 `json:"pageNumber"`
    PageSize int64 `json:"pageSize"`
    Users []logs.Erp `json:"users"`
}