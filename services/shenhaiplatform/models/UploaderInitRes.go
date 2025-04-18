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


type UploaderInitRes struct {

    /* bucket名称 (Optional) */
    Bucket string `json:"bucket"`

    /* 文件Code (Optional) */
    FileCode string `json:"fileCode"`

    /* 文件名称 (Optional) */
    FileName string `json:"fileName"`

    /* 任务流Code (Optional) */
    FlowCode string `json:"flowCode"`

    /* 文件夹上传的时候文件的路径 (Optional) */
    Prefix string `json:"prefix"`

    /* 区域 (Optional) */
    Region string `json:"region"`
}
