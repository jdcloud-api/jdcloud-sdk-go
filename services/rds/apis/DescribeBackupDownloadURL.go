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

package apis

import (
    "github.com/jdcloud-api/jdcloud-sdk-go/core"
    "reflect"
)

type DescribeBackupDownloadURLRequest struct {

    core.JDCloudRequest

    /* 地域代码  */
    RegionId string `json:"regionId"`

    /* 备份ID  */
    BackupId string `json:"backupId"`

    /* MySQL：无需此参数；SQL Server：指定该备份中需要获取下载链接的文件名称，SQL Server必须输入该参数 (Optional) */
    FileName *string `json:"fileName"`

    /* 指定下载链接的有效时间，单位秒,缺省为86400秒（即24小时） 取值范围：1-864000 (Optional) */
    UrlExpirationSecond *string `json:"urlExpirationSecond"`
}

/*
 * param regionId: 地域代码 
 * param backupId: 备份ID 
 * param fileName: MySQL：无需此参数；SQL Server：指定该备份中需要获取下载链接的文件名称，SQL Server必须输入该参数 (Optional)
 * param urlExpirationSecond: 指定下载链接的有效时间，单位秒,缺省为86400秒（即24小时） 取值范围：1-864000 (Optional)
 */
func NewDescribeBackupDownloadURLRequest(
    regionId string,
    backupId string,
) *DescribeBackupDownloadURLRequest {

	return &DescribeBackupDownloadURLRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/backups/{backupId}/downloadURLs",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        BackupId: backupId,
	}
}

func (r *DescribeBackupDownloadURLRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

func (r *DescribeBackupDownloadURLRequest) SetBackupId(backupId string) {
    r.BackupId = backupId
}

func (r *DescribeBackupDownloadURLRequest) SetFileName(fileName string) {
    r.FileName = &fileName
}

func (r *DescribeBackupDownloadURLRequest) SetUrlExpirationSecond(urlExpirationSecond string) {
    r.UrlExpirationSecond = &urlExpirationSecond
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeBackupDownloadURLRequest) GetRegionId() string {
    fieldName := "RegionId"
    reqType := reflect.TypeOf(r)
    value := reflect.ValueOf(r)
    _, ok := reqType.FieldByName(fieldName)
    if ok {
        return value.FieldByName(fieldName).String()
    }

    return ""
}

type DescribeBackupDownloadURLResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeBackupDownloadURLResult `json:"result"`
}

type DescribeBackupDownloadURLResult struct {
    PublicURL string `json:"publicURL"`
    InternalURL string `json:"internalURL"`
}