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

package client

import (
    "github.com/jdcloud-api/jdcloud-sdk-go/core"
    dh "github.com/jdcloud-api/jdcloud-sdk-go/services/dh/apis"
    "encoding/json"
    "errors"
)

type DhClient struct {
    core.JDCloudClient
}

func NewDhClient(credential *core.Credential) *DhClient {
    if credential == nil {
        return nil
    }

    config := core.NewConfig()
    config.SetEndpoint("dh.jdcloud-api.com")

    return &DhClient{
        core.JDCloudClient{
            Credential:  *credential,
            Config:      *config,
            ServiceName: "dh",
            Revision:    "1.1.3",
            Logger:      core.NewDefaultLogger(core.LogInfo),
        }}
}

func (c *DhClient) SetConfig(config *core.Config) {
    c.Config = *config
}

func (c *DhClient) SetLogger(logger core.Logger) {
    c.Logger = logger
}

func (c *DhClient) DisableLogger() {
    c.Logger = core.NewDummyLogger()
}

/* 查询专有宿主机列表。<br>
此接口支持分页查询，默认每页20条。
 */
func (c *DhClient) DescribeDedicatedHosts(request *dh.DescribeDedicatedHostsRequest) (*dh.DescribeDedicatedHostsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &dh.DescribeDedicatedHostsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询专有宿主机机型。
 */
func (c *DhClient) DescribeDedicatedHostType(request *dh.DescribeDedicatedHostTypeRequest) (*dh.DescribeDedicatedHostTypeResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &dh.DescribeDedicatedHostTypeResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 申请专有宿主机池。<br>
专有宿主机池为专有宿主机的集合，为用户提供独享且跨机架的宿主机资源池，基于资源独享及故障隔离实现业务安全、高可用部署。<br>
申请宿主机池需要指定机器类型，目前每一个专有宿主机池只能指定一种机器类型。可供选用的机器类型有通用性，内存性，高频计算型及GPU型。<br>
deployPolicy表示是否强制池中专有宿主机在每个AZ内数量均衡。可选值为prefered和required，默认为prefered。
 */
func (c *DhClient) AllocDedicatedPool(request *dh.AllocDedicatedPoolRequest) (*dh.AllocDedicatedPoolResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &dh.AllocDedicatedPoolResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 创建一台或多台指定机型的专有宿主机。<br>
创建专有宿主机必须指定专有宿主机池。
 */
func (c *DhClient) AllocDedicatedHosts(request *dh.AllocDedicatedHostsRequest) (*dh.AllocDedicatedHostsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &dh.AllocDedicatedHostsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 释放按配置计费、或包年包月已到期的单个专有宿主机。不能释放没有计费信息的专有宿主机。<br>
专有宿主机状态必须为可用<b>available</b>、不可用<b>unavailable</b>、维护中<b>under-assessment</b>，同时专有宿主机上必须没有云主机实例才可删除。<br>
 [MFA enabled] */
func (c *DhClient) ReleaseDedicatedHost(request *dh.ReleaseDedicatedHostRequest) (*dh.ReleaseDedicatedHostResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &dh.ReleaseDedicatedHostResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 修改专有宿主机池属性，包括名称、描述和可用区。
 */
func (c *DhClient) ModifyDedicatedPoolAttribute(request *dh.ModifyDedicatedPoolAttributeRequest) (*dh.ModifyDedicatedPoolAttributeResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &dh.ModifyDedicatedPoolAttributeResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询配额，支持的类型：专有宿主机、专有宿主机池。
 */
func (c *DhClient) DescribeQuotas(request *dh.DescribeQuotasRequest) (*dh.DescribeQuotasResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &dh.DescribeQuotasResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 修改专有宿主机部分属性，包括名称、描述。
 */
func (c *DhClient) ModifyDedicatedHostAttribute(request *dh.ModifyDedicatedHostAttributeRequest) (*dh.ModifyDedicatedHostAttributeResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &dh.ModifyDedicatedHostAttributeResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 释放单个专有宿主机池。<br>
专有宿主机池中必须没有专有宿主机时才可释放。
 [MFA enabled] */
func (c *DhClient) ReleaseDedicatedPool(request *dh.ReleaseDedicatedPoolRequest) (*dh.ReleaseDedicatedPoolResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &dh.ReleaseDedicatedPoolResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询专有宿主机池列表
 */
func (c *DhClient) DescribeDedicatedPools(request *dh.DescribeDedicatedPoolsRequest) (*dh.DescribeDedicatedPoolsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &dh.DescribeDedicatedPoolsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

