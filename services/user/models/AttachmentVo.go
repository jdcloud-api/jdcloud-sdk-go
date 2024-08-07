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


type AttachmentVo struct {

    /* null (Optional) */
    Id int `json:"id"`

    /* pin (Optional) */
    Pin string `json:"pin"`

    /* 运营给用户填写的备注 (Optional) */
    Remark string `json:"remark"`

    /* 行业属性 (Optional) */
    Industry string `json:"industry"`

    /* 子行业属性 (Optional) */
    SubIndustry string `json:"subIndustry"`

    /* 主营业务 (Optional) */
    Business string `json:"business"`

    /* 推荐码 (Optional) */
    RecommendCode string `json:"recommendCode"`

    /* 老用户-邮箱是否修改 (Optional) */
    OldUserMailModified string `json:"oldUserMailModified"`

    /* 老用户-手机是否修改 (Optional) */
    OldUserPhoneModified string `json:"oldUserPhoneModified"`

    /* 联系人 (Optional) */
    CpName string `json:"cpName"`

    /* 网址 (Optional) */
    Website string `json:"website"`

    /* 联系人国家 (Optional) */
    CpState string `json:"cpState"`

    /* 联系人省份 (Optional) */
    CpProvince string `json:"cpProvince"`

    /* 联系人城市 (Optional) */
    CpCity string `json:"cpCity"`

    /* 主营业务 (Optional) */
    CpCountry string `json:"cpCountry"`

    /* 联系人国家 (Optional) */
    CpAddress string `json:"cpAddress"`

    /* 联系人联系电话 (Optional) */
    CpTelphone string `json:"cpTelphone"`

    /* 商户号 (Optional) */
    MerchantId string `json:"merchantId"`
}
