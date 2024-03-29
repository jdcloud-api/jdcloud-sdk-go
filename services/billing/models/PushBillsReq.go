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


type PushBillsReq struct {

    /* 操作类型： 1、新增账单 2、更新账单数据  */
    OpType int `json:"opType"`

    /* 数据来源：10、物联网 11、视频云 12、CDN 13、PCDN 14、IDC 15、通信云  */
    DataSource int `json:"dataSource"`

    /* 全局唯一ID  */
    GlobalId string `json:"globalId"`

    /* 用户pin，不可修改  */
    Pin string `json:"pin"`

    /* 需要推送的账单列表，一次最多传10条  */
    BillList []BillVo `json:"billList"`
}
