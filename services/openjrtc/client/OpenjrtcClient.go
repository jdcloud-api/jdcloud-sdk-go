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
    openjrtc "github.com/jdcloud-api/jdcloud-sdk-go/services/openjrtc/apis"
    "encoding/json"
    "errors"
)

type OpenjrtcClient struct {
    core.JDCloudClient
}

func NewOpenjrtcClient(credential *core.Credential) *OpenjrtcClient {
    if credential == nil {
        return nil
    }

    config := core.NewConfig()
    config.SetEndpoint("openjrtc.jdcloud-api.com")

    return &OpenjrtcClient{
        core.JDCloudClient{
            Credential:  *credential,
            Config:      *config,
            ServiceName: "openjrtc",
            Revision:    "1.1.0",
            Logger:      core.NewDefaultLogger(core.LogInfo),
        }}
}

func (c *OpenjrtcClient) SetConfig(config *core.Config) {
    c.Config = *config
}

func (c *OpenjrtcClient) SetLogger(logger core.Logger) {
    c.Logger = logger
}

func (c *OpenjrtcClient) DisableLogger() {
    c.Logger = core.NewDummyLogger()
}

/* 查询应用下的房间列表
允许通过条件过滤查询，支持的过滤字段如下：
           - appId[eq] 按应用ID查询
 */
func (c *OpenjrtcClient) DescribeRooms(request *openjrtc.DescribeRoomsRequest) (*openjrtc.DescribeRoomsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &openjrtc.DescribeRoomsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 获取房间信息
 */
func (c *OpenjrtcClient) DescribeRoomInfo(request *openjrtc.DescribeRoomInfoRequest) (*openjrtc.DescribeRoomInfoResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &openjrtc.DescribeRoomInfoResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询应用appKey
 */
func (c *OpenjrtcClient) DescribeAppKey(request *openjrtc.DescribeAppKeyRequest) (*openjrtc.DescribeAppKeyResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &openjrtc.DescribeAppKeyResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询应用信息:
 */
func (c *OpenjrtcClient) DescribeApp(request *openjrtc.DescribeAppRequest) (*openjrtc.DescribeAppResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &openjrtc.DescribeAppResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询房间实时在线人数
 */
func (c *OpenjrtcClient) DescribeRoomOnlineUserNum(request *openjrtc.DescribeRoomOnlineUserNumRequest) (*openjrtc.DescribeRoomOnlineUserNumResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &openjrtc.DescribeRoomOnlineUserNumResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 修改房间
 */
func (c *OpenjrtcClient) UpdateRoom(request *openjrtc.UpdateRoomRequest) (*openjrtc.UpdateRoomResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &openjrtc.UpdateRoomResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询用户应用列表信息
 */
func (c *OpenjrtcClient) DescribeApps(request *openjrtc.DescribeAppsRequest) (*openjrtc.DescribeAppsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &openjrtc.DescribeAppsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 删除房间
 */
func (c *OpenjrtcClient) DeleteRoom(request *openjrtc.DeleteRoomRequest) (*openjrtc.DeleteRoomResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &openjrtc.DeleteRoomResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 创建房间
 */
func (c *OpenjrtcClient) CreateRoom(request *openjrtc.CreateRoomRequest) (*openjrtc.CreateRoomResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &openjrtc.CreateRoomResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 创建用户
 */
func (c *OpenjrtcClient) CreateUser(request *openjrtc.CreateUserRequest) (*openjrtc.CreateUserResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &openjrtc.CreateUserResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

