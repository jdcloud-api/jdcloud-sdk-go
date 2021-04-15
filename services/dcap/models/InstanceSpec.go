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


type InstanceSpec struct {

    /* 实例名称,长度限制32字节,允许英文字母,数字,下划线,中划线和中文  */
    InstanceName string `json:"instanceName"`

    /* 实例类型-规格 1->网关版本, 2->插件版本  */
    InstanceType int `json:"instanceType"`

    /* 实例副本数量  */
    Replicas int `json:"replicas"`

    /* 数据源ID  */
    DataSourceId string `json:"dataSourceId"`

    /* AccessKey  */
    AccessKey string `json:"accessKey"`

    /* SecretKey  */
    SecretKey string `json:"secretKey"`
}
