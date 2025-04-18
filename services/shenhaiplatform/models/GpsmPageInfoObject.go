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


type GpsmPageInfoObject struct {

    /*  (Optional) */
    Total int64 `json:"total"`

    /*  (Optional) */
    List []interface{} `json:"list"`

    /*  (Optional) */
    PageNum int `json:"pageNum"`

    /*  (Optional) */
    PageSize int `json:"pageSize"`

    /*  (Optional) */
    Size int `json:"size"`

    /*  (Optional) */
    StartRow int `json:"startRow"`

    /*  (Optional) */
    EndRow int `json:"endRow"`

    /*  (Optional) */
    Pages int `json:"pages"`

    /*  (Optional) */
    PrePage int `json:"prePage"`

    /*  (Optional) */
    NextPage int `json:"nextPage"`

    /*  (Optional) */
    IsFirstPage bool `json:"isFirstPage"`

    /*  (Optional) */
    IsLastPage bool `json:"isLastPage"`

    /*  (Optional) */
    HasPreviousPage bool `json:"hasPreviousPage"`

    /*  (Optional) */
    HasNextPage bool `json:"hasNextPage"`

    /*  (Optional) */
    NavigatePages int `json:"navigatePages"`

    /*  (Optional) */
    NavigatepageNums []int `json:"navigatepageNums"`

    /*  (Optional) */
    NavigateFirstPage int `json:"navigateFirstPage"`

    /*  (Optional) */
    NavigateLastPage int `json:"navigateLastPage"`
}
