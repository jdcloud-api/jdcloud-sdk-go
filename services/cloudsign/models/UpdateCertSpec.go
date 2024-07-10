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


type UpdateCertSpec struct {

    /* 个人用户姓名或企业名 (Optional) */
    UserName *string `json:"userName"`

    /* 0：个人用户，1：企业用户 (Optional) */
    UserType *int `json:"userType"`

    /* 邮箱 (Optional) */
    Mail *string `json:"mail"`

    /* 手机号 (Optional) */
    Phone *string `json:"phone"`

    /* 证件号类型。个人用户：身份证（idCardNum）,护照（passportNum）,港澳通行证（mtpNum）,自定义（customUserId）。企业用户：统一社会信用代码（usci）,组织机构代码（orgCode）,工商营业执照（businessNum）,自定义（customUserId） (Optional) */
    IdentifyField *string `json:"identifyField"`

    /* 证件号 (Optional) */
    IdentifyValue *string `json:"identifyValue"`

    /* P10 (Optional) */
    CertReqBuf *string `json:"certReqBuf"`

    /* 证书序列号 (Optional) */
    CertSerialNumber *string `json:"certSerialNumber"`
}
