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


type UserQuotaReqVo struct {

    /* 业务线 (Optional) */
    AppCode string `json:"appCode"`

    /* 可用配额 (Optional) */
    AvailableQuota int `json:"availableQuota"`

    /*  (Optional) */
    CountSql bool `json:"countSql"`

    /* 用户该地域该资源下的配额值 (Optional) */
    ExpireTime string `json:"expireTime"`

    /* id (Optional) */
    Id int64 `json:"id"`

    /*  (Optional) */
    OrderBy string `json:"orderBy"`

    /*  (Optional) */
    PageNum int `json:"pageNum"`

    /*  (Optional) */
    PageSize int `json:"pageSize"`

    /*  (Optional) */
    PageSizeZero bool `json:"pageSizeZero"`

    /* 父层id (Optional) */
    ResourceId string `json:"resourceId"`

    /* 用户pin (Optional) */
    Pin string `json:"pin"`

    /* 产品名称 (Optional) */
    ProductName string `json:"productName"`

    /*  (Optional) */
    Reasonable bool `json:"reasonable"`

    /* 地域 (Optional) */
    Region string `json:"region"`

    /* 区域配额 (Optional) */
    RegionQuota int64 `json:"regionQuota"`

    /* 资源产品线 (Optional) */
    ServiceCode string `json:"serviceCode"`

    /* 站点类型 (Optional) */
    SiteType int `json:"siteType"`

    /* 已用配额 (Optional) */
    UsedQuota int `json:"usedQuota"`

    /* 用户该地域该资源下的配额值 (Optional) */
    UserQuota int64 `json:"userQuota"`

    /* 扩展字段，json格式字符串 (Optional) */
    ExtraInfo string `json:"extraInfo"`
}
