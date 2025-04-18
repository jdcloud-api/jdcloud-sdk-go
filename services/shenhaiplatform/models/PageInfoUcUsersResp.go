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


type PageInfoUcUsersResp struct {

    /* 总数量 (Optional) */
    Total int `json:"total"`

    /* 分页页数 (Optional) */
    PageNum int `json:"pageNum"`

    /* 页面容量 (Optional) */
    PageSize int `json:"pageSize"`

    /* 项目成员列表 (Optional) */
    List []UcUsersResp `json:"list"`
}
