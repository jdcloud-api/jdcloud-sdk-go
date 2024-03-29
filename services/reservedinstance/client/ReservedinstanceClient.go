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
    reservedinstance "github.com/jdcloud-api/jdcloud-sdk-go/services/reservedinstance/apis"
    "encoding/json"
    "errors"
)

type ReservedinstanceClient struct {
    core.JDCloudClient
}

func NewReservedinstanceClient(credential *core.Credential) *ReservedinstanceClient {
    if credential == nil {
        return nil
    }

    config := core.NewConfig()
    config.SetEndpoint("reservedinstance.jdcloud-api.com")

    return &ReservedinstanceClient{
        core.JDCloudClient{
            Credential:  *credential,
            Config:      *config,
            ServiceName: "reservedinstance",
            Revision:    "0.1.2",
            Logger:      core.NewDefaultLogger(core.LogInfo),
        }}
}

func (c *ReservedinstanceClient) SetConfig(config *core.Config) {
    c.Config = *config
}

func (c *ReservedinstanceClient) SetLogger(logger core.Logger) {
    c.Logger = logger
}

func (c *ReservedinstanceClient) DisableLogger() {
    c.Logger = core.NewDummyLogger()
}

/* 查询实例抵扣券列表 */
func (c *ReservedinstanceClient) DescribeConsoleInstanceVouchers(request *reservedinstance.DescribeConsoleInstanceVouchersRequest) (*reservedinstance.DescribeConsoleInstanceVouchersResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &reservedinstance.DescribeConsoleInstanceVouchersResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询实例券使用明细列表 */
func (c *ReservedinstanceClient) DescribeDeductionDetails(request *reservedinstance.DescribeDeductionDetailsRequest) (*reservedinstance.DescribeDeductionDetailsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &reservedinstance.DescribeDeductionDetailsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

