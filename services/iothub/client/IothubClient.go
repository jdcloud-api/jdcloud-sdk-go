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
    iothub "github.com/jdcloud-api/jdcloud-sdk-go/services/iothub/apis"
    "encoding/json"
    "errors"
)

type IothubClient struct {
    core.JDCloudClient
}

func NewIothubClient(credential *core.Credential) *IothubClient {
    if credential == nil {
        return nil
    }

    config := core.NewConfig()
    config.SetEndpoint("iothub.jdcloud-api.com")

    return &IothubClient{
        core.JDCloudClient{
            Credential:  *credential,
            Config:      *config,
            ServiceName: "iothub",
            Revision:    "0.5.0",
            Logger:      core.NewDefaultLogger(core.LogInfo),
        }}
}

func (c *IothubClient) SetConfig(config *core.Config) {
    c.Config = *config
}

func (c *IothubClient) SetLogger(logger core.Logger) {
    c.Logger = logger
}

/* 物模型注册接口
 */
func (c *IothubClient) OmEnroll(request *iothub.OmEnrollRequest) (*iothub.OmEnrollResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &iothub.OmEnrollResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 激活一个设备，包括Edge、Dragon和普通设备
 */
func (c *IothubClient) DeviceActivate(request *iothub.DeviceActivateRequest) (*iothub.DeviceActivateResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &iothub.DeviceActivateResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 删除设备 */
func (c *IothubClient) DeleteDevice(request *iothub.DeleteDeviceRequest) (*iothub.DeleteDeviceResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &iothub.DeleteDeviceResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 客户用该接口可以登记模块
 */
func (c *IothubClient) ModuleEnroll(request *iothub.ModuleEnrollRequest) (*iothub.ModuleEnrollResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &iothub.ModuleEnrollResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 客户用该接口可以批量登记设备
 */
func (c *IothubClient) DevicesEnroll(request *iothub.DevicesEnrollRequest) (*iothub.DevicesEnrollResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &iothub.DevicesEnrollResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 客户用该接口可以对设备下发命令
 */
func (c *IothubClient) DeviceCommand(request *iothub.DeviceCommandRequest) (*iothub.DeviceCommandResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &iothub.DeviceCommandResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 客户用该接口可以修改模块预期状态
 */
func (c *IothubClient) ModuleState(request *iothub.ModuleStateRequest) (*iothub.ModuleStateResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &iothub.ModuleStateResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 物模型通过文件上传注册接口
 */
func (c *IothubClient) OmEnrollbyFile(request *iothub.OmEnrollbyFileRequest) (*iothub.OmEnrollbyFileResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &iothub.OmEnrollbyFileResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询设备在线信息 */
func (c *IothubClient) QueryDeviceOnlineInfos(request *iothub.QueryDeviceOnlineInfosRequest) (*iothub.QueryDeviceOnlineInfosResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &iothub.QueryDeviceOnlineInfosResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 客户用该接口可以查询设备预期状态
 */
func (c *IothubClient) QueryDeviceStates(request *iothub.QueryDeviceStatesRequest) (*iothub.QueryDeviceStatesResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &iothub.QueryDeviceStatesResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 客户用该接口可以查询设备命令列表
 */
func (c *IothubClient) QueryDeviceCommands(request *iothub.QueryDeviceCommandsRequest) (*iothub.QueryDeviceCommandsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &iothub.QueryDeviceCommandsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 获取下载物模型的外链地址
 */
func (c *IothubClient) GetOMPrivateURL(request *iothub.GetOMPrivateURLRequest) (*iothub.GetOMPrivateURLResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &iothub.GetOMPrivateURLResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 客户用该接口可以修改设备预期状态
 */
func (c *IothubClient) DeviceState(request *iothub.DeviceStateRequest) (*iothub.DeviceStateResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &iothub.DeviceStateResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

