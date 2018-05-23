// Copyright 2018-2025 JDCLOUD.COM
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


type ImportFile struct {

    /* 文件名称 (Optional) */
    Name string `json:"name"`

    /* 如果该文件是共享文件，则有全局ID，如不是共享文件，则为空 (Optional) */
    SharedFileGid string `json:"sharedFileGid"`

    /* 文件大小 (Optional) */
    SizeByte int `json:"sizeByte"`

    /* 文件上传完成时间 (Optional) */
    UploadTime string `json:"uploadTime"`

    /* 是否所属当前实例. “true”或者“false” (Optional) */
    IsLocal string `json:"isLocal"`
}