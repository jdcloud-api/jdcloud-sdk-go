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


type LegalPersonSpec struct {

    /* 企业类型：
0: 企业(ET_PE_QiYe)
1: 个体工商户(ET_SE_GeTiGongShangHu)
2: 政府机构/事业单位(ET_OU_ZhengFu_ShiYeDanWei)
 (Optional) */
    CompanyType *int `json:"companyType"`

    /* 企业名称  */
    CompanyName string `json:"companyName"`

    /* 统一社会信用代码或营业执照注册号  */
    IdCode string `json:"idCode"`

    /* 法定代表人姓名  */
    LegalPersonName string `json:"legalPersonName"`

    /* 法定代表人身份证号  */
    LegalPersonId string `json:"legalPersonId"`
}
