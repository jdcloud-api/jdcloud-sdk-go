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

type CreateAddressRequest struct {

    core.JDCloudRequest

    /*   */
    RegionId string `json:"regionId"`

    /* id (Optional) */
    Id *int `json:"id"`

    /* 收件人姓名 (Optional) */
    Name *string `json:"name"`

    /* 收件人电话 (Optional) */
    Phone *string `json:"phone"`

    /* 省 (Optional) */
    Province *string `json:"province"`

    /* 市 (Optional) */
    City *string `json:"city"`

    /* 区 (Optional) */
    County *string `json:"county"`

    /* 县 (Optional) */
    Town *string `json:"town"`

    /* 邮编 (Optional) */
    Address *string `json:"address"`

    /* 是否默认收货地址 (Optional) */
    IsDefault *int `json:"isDefault"`

    /* 更新时间 (Optional) */
    UpdateTime *string `json:"updateTime"`

    /* 用户pin (Optional) */
    Pin *string `json:"pin"`

    /* 创建时间 (Optional) */
    CreateTime *string `json:"createTime"`

    /* 邮编 (Optional) */
    ZipCode *string `json:"zipCode"`

    /* 省id (Optional) */
    ProvinceId *int `json:"provinceId"`

    /* 市id (Optional) */
    CityId *int `json:"cityId"`

    /* 区县id (Optional) */
    CountyId *int `json:"countyId"`

    /* 乡镇id (Optional) */
    TownId *int `json:"townId"`
}

/*
 * param regionId:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateAddressRequest(
    regionId string,
) *CreateAddressRequest {

	return &CreateAddressRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/address:add",
			Method:  "POST",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId:  (Required)
 * param id: id (Optional)
 * param name: 收件人姓名 (Optional)
 * param phone: 收件人电话 (Optional)
 * param province: 省 (Optional)
 * param city: 市 (Optional)
 * param county: 区 (Optional)
 * param town: 县 (Optional)
 * param address: 邮编 (Optional)
 * param isDefault: 是否默认收货地址 (Optional)
 * param updateTime: 更新时间 (Optional)
 * param pin: 用户pin (Optional)
 * param createTime: 创建时间 (Optional)
 * param zipCode: 邮编 (Optional)
 * param provinceId: 省id (Optional)
 * param cityId: 市id (Optional)
 * param countyId: 区县id (Optional)
 * param townId: 乡镇id (Optional)
 */
func NewCreateAddressRequestWithAllParams(
    regionId string,
    id *int,
    name *string,
    phone *string,
    province *string,
    city *string,
    county *string,
    town *string,
    address *string,
    isDefault *int,
    updateTime *string,
    pin *string,
    createTime *string,
    zipCode *string,
    provinceId *int,
    cityId *int,
    countyId *int,
    townId *int,
) *CreateAddressRequest {

    return &CreateAddressRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/address:add",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        Id: id,
        Name: name,
        Phone: phone,
        Province: province,
        City: city,
        County: county,
        Town: town,
        Address: address,
        IsDefault: isDefault,
        UpdateTime: updateTime,
        Pin: pin,
        CreateTime: createTime,
        ZipCode: zipCode,
        ProvinceId: provinceId,
        CityId: cityId,
        CountyId: countyId,
        TownId: townId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateAddressRequestWithoutParam() *CreateAddressRequest {

    return &CreateAddressRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/address:add",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: (Required) */
func (r *CreateAddressRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param id: id(Optional) */
func (r *CreateAddressRequest) SetId(id int) {
    r.Id = &id
}
/* param name: 收件人姓名(Optional) */
func (r *CreateAddressRequest) SetName(name string) {
    r.Name = &name
}
/* param phone: 收件人电话(Optional) */
func (r *CreateAddressRequest) SetPhone(phone string) {
    r.Phone = &phone
}
/* param province: 省(Optional) */
func (r *CreateAddressRequest) SetProvince(province string) {
    r.Province = &province
}
/* param city: 市(Optional) */
func (r *CreateAddressRequest) SetCity(city string) {
    r.City = &city
}
/* param county: 区(Optional) */
func (r *CreateAddressRequest) SetCounty(county string) {
    r.County = &county
}
/* param town: 县(Optional) */
func (r *CreateAddressRequest) SetTown(town string) {
    r.Town = &town
}
/* param address: 邮编(Optional) */
func (r *CreateAddressRequest) SetAddress(address string) {
    r.Address = &address
}
/* param isDefault: 是否默认收货地址(Optional) */
func (r *CreateAddressRequest) SetIsDefault(isDefault int) {
    r.IsDefault = &isDefault
}
/* param updateTime: 更新时间(Optional) */
func (r *CreateAddressRequest) SetUpdateTime(updateTime string) {
    r.UpdateTime = &updateTime
}
/* param pin: 用户pin(Optional) */
func (r *CreateAddressRequest) SetPin(pin string) {
    r.Pin = &pin
}
/* param createTime: 创建时间(Optional) */
func (r *CreateAddressRequest) SetCreateTime(createTime string) {
    r.CreateTime = &createTime
}
/* param zipCode: 邮编(Optional) */
func (r *CreateAddressRequest) SetZipCode(zipCode string) {
    r.ZipCode = &zipCode
}
/* param provinceId: 省id(Optional) */
func (r *CreateAddressRequest) SetProvinceId(provinceId int) {
    r.ProvinceId = &provinceId
}
/* param cityId: 市id(Optional) */
func (r *CreateAddressRequest) SetCityId(cityId int) {
    r.CityId = &cityId
}
/* param countyId: 区县id(Optional) */
func (r *CreateAddressRequest) SetCountyId(countyId int) {
    r.CountyId = &countyId
}
/* param townId: 乡镇id(Optional) */
func (r *CreateAddressRequest) SetTownId(townId int) {
    r.TownId = &townId
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateAddressRequest) GetRegionId() string {
    return r.RegionId
}

type CreateAddressResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateAddressResult `json:"result"`
}

type CreateAddressResult struct {
    Code int `json:"code"`
    Msg string `json:"msg"`
}