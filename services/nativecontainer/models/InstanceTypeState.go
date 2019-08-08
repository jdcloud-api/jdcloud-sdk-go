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


type InstanceTypeState struct {

    /* 可用区 (Optional) */
    Az string `json:"az"`

    /* 可售卖情况，true:可售卖、false:已售罄不可用 (Optional) */
    InStock bool `json:"inStock"`

    /* 在线情况，true:在线、false:已下线不可用 (Optional) */
    Online bool `json:"online"`
}
