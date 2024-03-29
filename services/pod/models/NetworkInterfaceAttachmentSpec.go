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

import vpc "github.com/jdcloud-api/jdcloud-sdk-go/services/vpc/models"

type NetworkInterfaceAttachmentSpec struct {

    /* 指明删除pod时是否删除网卡，主网卡默认是true,辅助网卡默认false (Optional) */
    AutoDelete *bool `json:"autoDelete"`

    /* 设备Index，目前pod只支持一个网卡，所以只能设置为1 (Optional) */
    DeviceIndex *int `json:"deviceIndex"`

    /* 网卡接口规范  */
    NetworkInterface *vpc.NetworkInterfaceSpec `json:"networkInterface"`
}
