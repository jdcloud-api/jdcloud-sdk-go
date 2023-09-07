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
)

type DescribeBackupDownloadURLRequest struct {

    core.JDCloudRequest

    /* 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)  */
    RegionId string `json:"regionId"`

    /* 备份ID  */
    BackupId string `json:"backupId"`

    /* 文件名称<br>- MySQL：不支持该参数<br>- SQL Server：必须输入该参数，指定该备份中需要获取下载链接的文件名称。备份中的文件名（不包括后缀）即为备份的数据库名。例如文件名为my_test_db.bak，表示该文件是my_test_db数据库的备份 (Optional) */
    FileName *string `json:"fileName"`

    /* 指定下载链接的过期时间，单位秒, 取值范围为 1 ~ 86400 秒；支持 SQL Server：缺省为 86400 秒。支持 MySQL, Percona, MariaDB：缺省为 300 秒。 (Optional) */
    UrlExpirationSecond *string `json:"urlExpirationSecond"`
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param backupId: 备份ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeBackupDownloadURLRequest(
    regionId string,
    backupId string,
) *DescribeBackupDownloadURLRequest {

	return &DescribeBackupDownloadURLRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/backups/{backupId}:describeBackupDownloadURL",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        BackupId: backupId,
	}
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param backupId: 备份ID (Required)
 * param fileName: 文件名称<br>- MySQL：不支持该参数<br>- SQL Server：必须输入该参数，指定该备份中需要获取下载链接的文件名称。备份中的文件名（不包括后缀）即为备份的数据库名。例如文件名为my_test_db.bak，表示该文件是my_test_db数据库的备份 (Optional)
 * param urlExpirationSecond: 指定下载链接的过期时间，单位秒, 取值范围为 1 ~ 86400 秒；支持 SQL Server：缺省为 86400 秒。支持 MySQL, Percona, MariaDB：缺省为 300 秒。 (Optional)
 */
func NewDescribeBackupDownloadURLRequestWithAllParams(
    regionId string,
    backupId string,
    fileName *string,
    urlExpirationSecond *string,
) *DescribeBackupDownloadURLRequest {

    return &DescribeBackupDownloadURLRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/backups/{backupId}:describeBackupDownloadURL",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        BackupId: backupId,
        FileName: fileName,
        UrlExpirationSecond: urlExpirationSecond,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeBackupDownloadURLRequestWithoutParam() *DescribeBackupDownloadURLRequest {

    return &DescribeBackupDownloadURLRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/backups/{backupId}:describeBackupDownloadURL",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)(Required) */
func (r *DescribeBackupDownloadURLRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param backupId: 备份ID(Required) */
func (r *DescribeBackupDownloadURLRequest) SetBackupId(backupId string) {
    r.BackupId = backupId
}
/* param fileName: 文件名称<br>- MySQL：不支持该参数<br>- SQL Server：必须输入该参数，指定该备份中需要获取下载链接的文件名称。备份中的文件名（不包括后缀）即为备份的数据库名。例如文件名为my_test_db.bak，表示该文件是my_test_db数据库的备份(Optional) */
func (r *DescribeBackupDownloadURLRequest) SetFileName(fileName string) {
    r.FileName = &fileName
}
/* param urlExpirationSecond: 指定下载链接的过期时间，单位秒, 取值范围为 1 ~ 86400 秒；支持 SQL Server：缺省为 86400 秒。支持 MySQL, Percona, MariaDB：缺省为 300 秒。(Optional) */
func (r *DescribeBackupDownloadURLRequest) SetUrlExpirationSecond(urlExpirationSecond string) {
    r.UrlExpirationSecond = &urlExpirationSecond
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeBackupDownloadURLRequest) GetRegionId() string {
    return r.RegionId
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