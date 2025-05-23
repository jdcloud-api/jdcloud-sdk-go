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


type UranusUploaderChunkReq struct {

    /* 文件类型 (Optional) */
    FileTypeCode int `json:"fileTypeCode"`

    /* 是否spark-jar的管理模块,spark-jar操作hdfs通过字段区分 (Optional) */
    IsJarManagement bool `json:"isJarManagement"`

    /* 任务流Code  */
    FlowCode string `json:"flowCode"`

    /* 文件夹上传的时候文件的相对路径属性 (Optional) */
    RelativePath string `json:"relativePath"`

    /* 文件唯一标识 (Optional) */
    UploadId string `json:"uploadId"`

    /* 当前块的次序，第一个块是 1，注意不是从 0 开始的 (Optional) */
    ChunkNumber int `json:"chunkNumber"`

    /* 文件被分成块的总数 (Optional) */
    TotalChunks int `json:"totalChunks"`

    /* 分块大小，根据 totalSize 和这个值你就可以计算出总共的块数。注意最后一块的大小可能会比这个要大 (Optional) */
    ChunkSize int64 `json:"chunkSize"`

    /* 当前块的大小，实际大小 (Optional) */
    CurrentChunkSize int64 `json:"currentChunkSize"`

    /* 文件总大小 (Optional) */
    TotalSize int64 `json:"totalSize"`

    /* 文件名  */
    Filename string `json:"filename"`

    /*  (Optional) */
    File string `json:"file"`

    /* 文件字节数组 (Optional) */
    FileBytes []string `json:"fileBytes"`

    /* partETagList (Optional) */
    PartETagList []UranusUploaderPartETag `json:"partETagList"`
}
