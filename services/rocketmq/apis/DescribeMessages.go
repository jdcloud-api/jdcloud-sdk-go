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
    "github.com/jdcloud-api/jdcloud-sdk-go/core"
    rocketmq "github.com/jdcloud-api/jdcloud-sdk-go/services/rocketmq/models"
)

type DescribeMessagesRequest struct {

    core.JDCloudRequest

    /* 区域ID  */
    RegionId string `json:"regionId"`

    /* 实例id  */
    InstanceId string `json:"instanceId"`

    /* topic 名称  */
    Topic string `json:"topic"`

    /* 如果填写了就按messageId查询 (Optional) */
    MessageId *string `json:"messageId"`

    /* 如果填写了就按key查询 (Optional) */
    Key *string `json:"key"`

    /* 开始时间 (Optional) */
    Begin *string `json:"begin"`

    /* 结束时间 (Optional) */
    End *string `json:"end"`

    /* 分页大小；  */
    PageSize int `json:"pageSize"`

    /* 页码  */
    PageNumber int `json:"pageNumber"`
}

/*
 * param regionId: 区域ID (Required)
 * param instanceId: 实例id (Required)
 * param topic: topic 名称 (Required)
 * param pageSize: 分页大小； (Required)
 * param pageNumber: 页码 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeMessagesRequest(
    regionId string,
    instanceId string,
    topic string,
    pageSize int,
    pageNumber int,
) *DescribeMessagesRequest {

	return &DescribeMessagesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/topics/{topic}/messages",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        Topic: topic,
        PageSize: pageSize,
        PageNumber: pageNumber,
	}
}

/*
 * param regionId: 区域ID (Required)
 * param instanceId: 实例id (Required)
 * param topic: topic 名称 (Required)
 * param messageId: 如果填写了就按messageId查询 (Optional)
 * param key: 如果填写了就按key查询 (Optional)
 * param begin: 开始时间 (Optional)
 * param end: 结束时间 (Optional)
 * param pageSize: 分页大小； (Required)
 * param pageNumber: 页码 (Required)
 */
func NewDescribeMessagesRequestWithAllParams(
    regionId string,
    instanceId string,
    topic string,
    messageId *string,
    key *string,
    begin *string,
    end *string,
    pageSize int,
    pageNumber int,
) *DescribeMessagesRequest {

    return &DescribeMessagesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/topics/{topic}/messages",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        Topic: topic,
        MessageId: messageId,
        Key: key,
        Begin: begin,
        End: end,
        PageSize: pageSize,
        PageNumber: pageNumber,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeMessagesRequestWithoutParam() *DescribeMessagesRequest {

    return &DescribeMessagesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/topics/{topic}/messages",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 区域ID(Required) */
func (r *DescribeMessagesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param instanceId: 实例id(Required) */
func (r *DescribeMessagesRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}
/* param topic: topic 名称(Required) */
func (r *DescribeMessagesRequest) SetTopic(topic string) {
    r.Topic = topic
}
/* param messageId: 如果填写了就按messageId查询(Optional) */
func (r *DescribeMessagesRequest) SetMessageId(messageId string) {
    r.MessageId = &messageId
}
/* param key: 如果填写了就按key查询(Optional) */
func (r *DescribeMessagesRequest) SetKey(key string) {
    r.Key = &key
}
/* param begin: 开始时间(Optional) */
func (r *DescribeMessagesRequest) SetBegin(begin string) {
    r.Begin = &begin
}
/* param end: 结束时间(Optional) */
func (r *DescribeMessagesRequest) SetEnd(end string) {
    r.End = &end
}
/* param pageSize: 分页大小；(Required) */
func (r *DescribeMessagesRequest) SetPageSize(pageSize int) {
    r.PageSize = pageSize
}
/* param pageNumber: 页码(Required) */
func (r *DescribeMessagesRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = pageNumber
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeMessagesRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeMessagesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeMessagesResult `json:"result"`
}

type DescribeMessagesResult struct {
    MessageList []rocketmq.Message `json:"messageList"`
    TotalCount int `json:"totalCount"`
}