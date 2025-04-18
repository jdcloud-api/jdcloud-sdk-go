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


type JtlasPageSearchTablePrivilegeReq struct {

    /* 数据库（工作空间编码）  */
    Database string `json:"database"`

    /* 表名称  */
    TableName string `json:"tableName"`

    /*  (Optional) */
    PrivilegeTypes []string `json:"privilegeTypes"`

    /* 分页参数-页码  */
    PageNum int `json:"pageNum"`

    /* 分页参数-页数  */
    PageSize int `json:"pageSize"`
}
