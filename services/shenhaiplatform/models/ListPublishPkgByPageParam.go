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


type ListPublishPkgByPageParam struct {

    /* 页面容量 (Optional) */
    PageSize int `json:"pageSize"`

    /* 页号 (Optional) */
    PageNum int `json:"pageNum"`

    /* 发布包名称 (Optional) */
    PkgName string `json:"pkgName"`

    /* 发布包ID (Optional) */
    PkgId int64 `json:"pkgId"`

    /* 发布人 (Optional) */
    Publisher string `json:"publisher"`

    /* 发布包状态 (Optional) */
    PkgStatus string `json:"pkgStatus"`

    /* 发布时间开始 (Optional) */
    PkgPublishTimeStart int64 `json:"pkgPublishTimeStart"`

    /* 发布时间结束 (Optional) */
    PkgPublishTimeEnd int64 `json:"pkgPublishTimeEnd"`

    /* 打包人 (Optional) */
    PkgCreator string `json:"pkgCreator"`

    /* 打包时间开始 (Optional) */
    PkgCreateTimeStart int64 `json:"pkgCreateTimeStart"`

    /* 打包时间结束 (Optional) */
    PkgCreateTimeEnd int64 `json:"pkgCreateTimeEnd"`
}
