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


type DistributorContractPolicyDTO struct {

    /* ID (Optional) */
    Id int `json:"id"`

    /* 服务商pin (Optional) */
    DistributorPin string `json:"distributorPin"`

    /* 服务商主协议PIN(1是0否) (Optional) */
    MainFlag int `json:"mainFlag"`

    /* 服务商政策ID (Optional) */
    ReturnPolicyId string `json:"returnPolicyId"`

    /* 服务商政策名称 (Optional) */
    ReturnPolicyName string `json:"returnPolicyName"`
}
