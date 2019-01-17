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
)

type DeleteDeadLettersRequest struct {

    core.JDCloudRequest

    /* 所在区域的Region ID  */
    RegionId string `json:"regionId"`

    /* topic 名称  */
    TopicName string `json:"topicName"`

    /* consumerGroupId  */
    ConsumerGroupId string `json:"consumerGroupId"`

    /* messageIds,多个逗号隔开，不传该值就是删除所有的死信 (Optional) */
    MessageIds *string `json:"messageIds"`
}

/*
 * param regionId: 所在区域的Region ID (Required)
 * param topicName: topic 名称 (Required)
 * param consumerGroupId: consumerGroupId (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteDeadLettersRequest(
    regionId string,
    topicName string,
    consumerGroupId string,
) *DeleteDeadLettersRequest {

	return &DeleteDeadLettersRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/topics/{topicName}/subscriptions/{consumerGroupId}:deleteDeadLetters",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        TopicName: topicName,
        ConsumerGroupId: consumerGroupId,
	}
}

/*
 * param regionId: 所在区域的Region ID (Required)
 * param topicName: topic 名称 (Required)
 * param consumerGroupId: consumerGroupId (Required)
 * param messageIds: messageIds,多个逗号隔开，不传该值就是删除所有的死信 (Optional)
 */
func NewDeleteDeadLettersRequestWithAllParams(
    regionId string,
    topicName string,
    consumerGroupId string,
    messageIds *string,
) *DeleteDeadLettersRequest {

    return &DeleteDeadLettersRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/topics/{topicName}/subscriptions/{consumerGroupId}:deleteDeadLetters",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        TopicName: topicName,
        ConsumerGroupId: consumerGroupId,
        MessageIds: messageIds,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteDeadLettersRequestWithoutParam() *DeleteDeadLettersRequest {

    return &DeleteDeadLettersRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/topics/{topicName}/subscriptions/{consumerGroupId}:deleteDeadLetters",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 所在区域的Region ID(Required) */
func (r *DeleteDeadLettersRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param topicName: topic 名称(Required) */
func (r *DeleteDeadLettersRequest) SetTopicName(topicName string) {
    r.TopicName = topicName
}

/* param consumerGroupId: consumerGroupId(Required) */
func (r *DeleteDeadLettersRequest) SetConsumerGroupId(consumerGroupId string) {
    r.ConsumerGroupId = consumerGroupId
}

/* param messageIds: messageIds,多个逗号隔开，不传该值就是删除所有的死信(Optional) */
func (r *DeleteDeadLettersRequest) SetMessageIds(messageIds string) {
    r.MessageIds = &messageIds
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteDeadLettersRequest) GetRegionId() string {
    return r.RegionId
}

type DeleteDeadLettersResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteDeadLettersResult `json:"result"`
}

type DeleteDeadLettersResult struct {
    MessageIds []string `json:"messageIds"`
}