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
    ias "github.com/jdcloud-api/jdcloud-sdk-go/services/ias/apis"
    "encoding/json"
    "errors"
)

type IasClient struct {
    core.JDCloudClient
}

func NewIasClient(credential *core.Credential) *IasClient {
    if credential == nil {
        return nil
    }

    config := core.NewConfig()
    config.SetEndpoint("ias.jdcloud-api.com")

    return &IasClient{
        core.JDCloudClient{
            Credential:  *credential,
            Config:      *config,
            ServiceName: "ias",
            Revision:    "0.2.0",
            Logger:      core.NewDefaultLogger(core.LogInfo),
        }}
}

func (c *IasClient) SetConfig(config *core.Config) {
    c.Config = *config
}

func (c *IasClient) SetLogger(logger core.Logger) {
    c.Logger = logger
}

/* 创建app */
func (c *IasClient) CreateApp(request *ias.CreateAppRequest) (*ias.CreateAppResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ias.CreateAppResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 更新app */
func (c *IasClient) UpdateApp(request *ias.UpdateAppRequest) (*ias.UpdateAppResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ias.UpdateAppResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 获取主账号下所有应用 */
func (c *IasClient) GetApps(request *ias.GetAppsRequest) (*ias.GetAppsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ias.GetAppsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 删除应用 */
func (c *IasClient) DeleteApp(request *ias.DeleteAppRequest) (*ias.DeleteAppResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ias.DeleteAppResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 获取应用 */
func (c *IasClient) GetApp(request *ias.GetAppRequest) (*ias.GetAppResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ias.GetAppResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

