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
    jdccs "github.com/jdcloud-api/jdcloud-sdk-go/services/jdccs/apis"
    "encoding/json"
    "errors"
)

type JdccsClient struct {
    core.JDCloudClient
}

func NewJdccsClient(credential *core.Credential) *JdccsClient {
    if credential == nil {
        return nil
    }

    config := core.NewConfig()
    config.SetEndpoint("jdccs.jdcloud-api.com")

    return &JdccsClient{
        core.JDCloudClient{
            Credential:  *credential,
            Config:      *config,
            ServiceName: "jdccs",
            Revision:    "1.1.2",
            Logger:      core.NewDefaultLogger(core.LogInfo),
        }}
}

func (c *JdccsClient) SetConfig(config *core.Config) {
    c.Config = *config
}

func (c *JdccsClient) SetLogger(logger core.Logger) {
    c.Logger = logger
}

func (c *JdccsClient) DisableLogger() {
    c.Logger = core.NewDummyLogger()
}

/* 创建报警 */
func (c *JdccsClient) CreateAlarm(request *jdccs.CreateAlarmRequest) (*jdccs.CreateAlarmResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdccs.CreateAlarmResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询工单列表 */
func (c *JdccsClient) DescribeTickets(request *jdccs.DescribeTicketsRequest) (*jdccs.DescribeTicketsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdccs.DescribeTicketsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 报警历史列表 */
func (c *JdccsClient) DescribeAlarmHistory(request *jdccs.DescribeAlarmHistoryRequest) (*jdccs.DescribeAlarmHistoryResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdccs.DescribeAlarmHistoryResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询IDC机房列表 */
func (c *JdccsClient) DescribeIdcs(request *jdccs.DescribeIdcsRequest) (*jdccs.DescribeIdcsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdccs.DescribeIdcsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 启用、禁用报警 */
func (c *JdccsClient) SwitchAlarm(request *jdccs.SwitchAlarmRequest) (*jdccs.SwitchAlarmResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdccs.SwitchAlarmResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询开放设备数据信息接口 */
func (c *JdccsClient) DescribeOpenDevicesData(request *jdccs.DescribeOpenDevicesDataRequest) (*jdccs.DescribeOpenDevicesDataResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdccs.DescribeOpenDevicesDataResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询带宽（出口）详情 */
func (c *JdccsClient) DescribeBandwidth(request *jdccs.DescribeBandwidthRequest) (*jdccs.DescribeBandwidthResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdccs.DescribeBandwidthResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 提交通用服务工单 */
func (c *JdccsClient) CreateGeneralServicesTicket(request *jdccs.CreateGeneralServicesTicketRequest) (*jdccs.CreateGeneralServicesTicketResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdccs.CreateGeneralServicesTicketResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询机柜列表 */
func (c *JdccsClient) DescribeCabinets(request *jdccs.DescribeCabinetsRequest) (*jdccs.DescribeCabinetsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdccs.DescribeCabinetsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询开放设备列表 */
func (c *JdccsClient) DescribeOpenDevices(request *jdccs.DescribeOpenDevicesRequest) (*jdccs.DescribeOpenDevicesResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdccs.DescribeOpenDevicesResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查看某资源单个监控项数据 */
func (c *JdccsClient) DescribeMetricData(request *jdccs.DescribeMetricDataRequest) (*jdccs.DescribeMetricDataResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdccs.DescribeMetricDataResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询工单详情 */
func (c *JdccsClient) DescribeTicket(request *jdccs.DescribeTicketRequest) (*jdccs.DescribeTicketResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdccs.DescribeTicketResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查看某资源的最后一个监控数据点（目前只支持机柜电流） */
func (c *JdccsClient) LastDownsample(request *jdccs.LastDownsampleRequest) (*jdccs.LastDownsampleResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdccs.LastDownsampleResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询机房房间号列表 */
func (c *JdccsClient) DescribeRooms(request *jdccs.DescribeRoomsRequest) (*jdccs.DescribeRoomsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdccs.DescribeRoomsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询带宽（出口）流量（资源）详情 */
func (c *JdccsClient) DescribeBandwidthTraffic(request *jdccs.DescribeBandwidthTrafficRequest) (*jdccs.DescribeBandwidthTrafficResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdccs.DescribeBandwidthTrafficResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询设备详情 */
func (c *JdccsClient) DescribeDevice(request *jdccs.DescribeDeviceRequest) (*jdccs.DescribeDeviceResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdccs.DescribeDeviceResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询带宽（出口）流量列表 */
func (c *JdccsClient) DescribeBandwidthTraffics(request *jdccs.DescribeBandwidthTrafficsRequest) (*jdccs.DescribeBandwidthTrafficsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdccs.DescribeBandwidthTrafficsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询报警规则列表 */
func (c *JdccsClient) DescribeAlarms(request *jdccs.DescribeAlarmsRequest) (*jdccs.DescribeAlarmsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdccs.DescribeAlarmsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询带宽（出口）列表 */
func (c *JdccsClient) DescribeBandwidths(request *jdccs.DescribeBandwidthsRequest) (*jdccs.DescribeBandwidthsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdccs.DescribeBandwidthsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询设备列表 */
func (c *JdccsClient) DescribeDevices(request *jdccs.DescribeDevicesRequest) (*jdccs.DescribeDevicesResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdccs.DescribeDevicesResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询公网IP列表 */
func (c *JdccsClient) DescribeIps(request *jdccs.DescribeIpsRequest) (*jdccs.DescribeIpsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdccs.DescribeIpsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 修改报警规则 */
func (c *JdccsClient) UpdateAlarm(request *jdccs.UpdateAlarmRequest) (*jdccs.UpdateAlarmResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdccs.UpdateAlarmResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询可用监控项列表 */
func (c *JdccsClient) DescribeMetrics(request *jdccs.DescribeMetricsRequest) (*jdccs.DescribeMetricsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdccs.DescribeMetricsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 根据IP网段查询流量采样数据 */
func (c *JdccsClient) DescribeTrafficSampling(request *jdccs.DescribeTrafficSamplingRequest) (*jdccs.DescribeTrafficSamplingResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdccs.DescribeTrafficSamplingResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询机柜详情 */
func (c *JdccsClient) DescribeCabinet(request *jdccs.DescribeCabinetRequest) (*jdccs.DescribeCabinetResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdccs.DescribeCabinetResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询机房资源概览 */
func (c *JdccsClient) DescribeIdcOverview(request *jdccs.DescribeIdcOverviewRequest) (*jdccs.DescribeIdcOverviewResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdccs.DescribeIdcOverviewResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询报警规则详情 */
func (c *JdccsClient) DescribeAlarm(request *jdccs.DescribeAlarmRequest) (*jdccs.DescribeAlarmResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdccs.DescribeAlarmResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 删除报警 */
func (c *JdccsClient) DeleteAlarm(request *jdccs.DeleteAlarmRequest) (*jdccs.DeleteAlarmResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdccs.DeleteAlarmResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

