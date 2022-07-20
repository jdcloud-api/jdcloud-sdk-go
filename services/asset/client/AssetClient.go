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
    "github.com/lshuining/jdcloud-sdk-go/core"
    asset "github.com/lshuining/jdcloud-sdk-go/services/asset/apis"
    "encoding/json"
    "errors"
)

type AssetClient struct {
    core.JDCloudClient
}

func NewAssetClient(credential *core.Credential) *AssetClient {
    if credential == nil {
        return nil
    }

    config := core.NewConfig()
    config.SetEndpoint("asset.jdcloud-api.com")

    return &AssetClient{
        core.JDCloudClient{
            Credential:  *credential,
            Config:      *config,
            ServiceName: "asset",
            Revision:    "0.0.3",
            Logger:      core.NewDefaultLogger(core.LogInfo),
        }}
}

func (c *AssetClient) SetConfig(config *core.Config) {
    c.Config = *config
}

func (c *AssetClient) SetLogger(logger core.Logger) {
    c.Logger = logger
}

func (c *AssetClient) DisableLogger() {
    c.Logger = core.NewDummyLogger()
}

/* 查询账户金额（总金额、可用金额、冻结金额、可提现金额、提现中金额） */
func (c *AssetClient) DescribeAccountAmount(request *asset.DescribeAccountAmountRequest) (*asset.DescribeAccountAmountResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &asset.DescribeAccountAmountResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 设置余额预警信息 */
func (c *AssetClient) ModifyBalanceWarningInfo(request *asset.ModifyBalanceWarningInfoRequest) (*asset.ModifyBalanceWarningInfoResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &asset.ModifyBalanceWarningInfoResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

