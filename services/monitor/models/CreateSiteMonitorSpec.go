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


type CreateSiteMonitorSpec struct {

    /*  (Optional) */
    Address string `json:"address"`

    /*  (Optional) */
    AdvanceChecked string `json:"advanceChecked"`

    /*  (Optional) */
    CreatedTime int64 `json:"createdTime"`

    /*  (Optional) */
    Cycle int64 `json:"cycle"`

    /*  (Optional) */
    DefaultSource string `json:"defaultSource"`

    /*  (Optional) */
    Enabled string `json:"enabled"`

    /*  (Optional) */
    HawkeyeId int64 `json:"hawkeyeId"`

    /*  (Optional) */
    HttpOption SiteMonitorHttpOption `json:"httpOption"`

    /*  (Optional) */
    Id string `json:"id"`

    /*  (Optional) */
    IsDeleted string `json:"isDeleted"`

    /*  (Optional) */
    Name string `json:"name"`

    /*  (Optional) */
    Pin string `json:"pin"`

    /*  (Optional) */
    Port string `json:"port"`

    /*  (Optional) */
    Source []SiteMonitorSource `json:"source"`

    /*  (Optional) */
    Stats interface{} `json:"stats"`

    /*  (Optional) */
    TcpOption SiteMonitorTcpOption `json:"tcpOption"`

    /*  (Optional) */
    Type string `json:"type"`

    /*  (Optional) */
    UdpOption SiteMonitorUdpOption `json:"udpOption"`

    /*  (Optional) */
    UpdatedTime int64 `json:"updatedTime"`
}
