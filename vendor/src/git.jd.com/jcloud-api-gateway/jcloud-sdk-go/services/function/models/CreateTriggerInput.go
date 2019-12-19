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

package models


type CreateTriggerInput struct {

    /* 触发器configurationId，对应jqs的jrn (Optional) */
    ConfigurationId string `json:"configurationId"`

    /* 事件源，目前支持apigateway、oss和jqs (Optional) */
    EventSource string `json:"eventSource"`

    /* 事件源Id，分别对应apiId、BucketId和jqsId (Optional) */
    EventSourceId string `json:"eventSourceId"`

    /* jqs queue名称 (Optional) */
    JqsName string `json:"jqsName"`

    /* jqs触发器批处理大小 (Optional) */
    JqsBatchSize int `json:"jqsBatchSize"`
}