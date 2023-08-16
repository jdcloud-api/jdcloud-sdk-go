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


type AutoRenewSpec struct {

    /* 是否开启自动续费,取值范围"OPEN","CLOSE" (Optional) */
    AutoRenewStatus string `json:"autoRenewStatus"`

    /* 自动续费周期,取值为MONTH，YEAR；autoRenewStatus为OPEN时默认为MONTH。 (Optional) */
    RenewTimeUnit string `json:"renewTimeUnit"`

    /* 端口租用费付费时长，renewTimeUnit为MONTH时取值范围1，2，3，4，5，6，7，8，9；renewTimeUnit为YEAR时取值范围1,2,3。autoRenewStatus为OPEN时默认为1 (Optional) */
    RenewTimeSpan int `json:"renewTimeSpan"`
}
