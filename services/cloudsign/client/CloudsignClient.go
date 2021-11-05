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
    cloudsign "github.com/jdcloud-api/jdcloud-sdk-go/services/cloudsign/apis"
    "encoding/json"
    "errors"
)

type CloudsignClient struct {
    core.JDCloudClient
}

func NewCloudsignClient(credential *core.Credential) *CloudsignClient {
    if credential == nil {
        return nil
    }

    config := core.NewConfig()
    config.SetEndpoint("cloudsign.jdcloud-api.com")

    return &CloudsignClient{
        core.JDCloudClient{
            Credential:  *credential,
            Config:      *config,
            ServiceName: "cloudsign",
            Revision:    "1.1.2",
            Logger:      core.NewDefaultLogger(core.LogInfo),
        }}
}

func (c *CloudsignClient) SetConfig(config *core.Config) {
    c.Config = *config
}

func (c *CloudsignClient) SetLogger(logger core.Logger) {
    c.Logger = logger
}

func (c *CloudsignClient) DisableLogger() {
    c.Logger = core.NewDummyLogger()
}

/* 查询服务开通状态 */
func (c *CloudsignClient) DescribeApplyStatus(request *cloudsign.DescribeApplyStatusRequest) (*cloudsign.DescribeApplyStatusResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &cloudsign.DescribeApplyStatusResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 启用合同存管 */
func (c *CloudsignClient) EnableContractSave(request *cloudsign.EnableContractSaveRequest) (*cloudsign.EnableContractSaveResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &cloudsign.EnableContractSaveResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 获取合同模板列表 */
func (c *CloudsignClient) DescribeTemplateList(request *cloudsign.DescribeTemplateListRequest) (*cloudsign.DescribeTemplateListResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &cloudsign.DescribeTemplateListResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 获取已签章合同列表 */
func (c *CloudsignClient) DescribeContractList(request *cloudsign.DescribeContractListRequest) (*cloudsign.DescribeContractListResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &cloudsign.DescribeContractListResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 上传印章 */
func (c *CloudsignClient) UploadStamp(request *cloudsign.UploadStampRequest) (*cloudsign.UploadStampResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &cloudsign.UploadStampResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 1. 下载印章
2. 多个印章id用逗号分隔
 [MFA enabled] */
func (c *CloudsignClient) DownloadStamps(request *cloudsign.DownloadStampsRequest) (*cloudsign.DownloadStampsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &cloudsign.DownloadStampsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 填充合同模板 */
func (c *CloudsignClient) PaddingTemplate(request *cloudsign.PaddingTemplateRequest) (*cloudsign.PaddingTemplateResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &cloudsign.PaddingTemplateResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 合同签章四种方式：
1. 合同文件 + 印章文件：contractContent + stampContent
2. 合同文件 + 印章ID：contractContent + stampId
3. 模板ID + 印章文件：templateId + stampContent
4. 模板ID + 印章ID：templateId + stampId
 */
func (c *CloudsignClient) SignContract(request *cloudsign.SignContractRequest) (*cloudsign.SignContractResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &cloudsign.SignContractResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 1. 下载已签章的合同
2. 多个合同id用逗号分隔
 [MFA enabled] */
func (c *CloudsignClient) DownloadContracts(request *cloudsign.DownloadContractsRequest) (*cloudsign.DownloadContractsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &cloudsign.DownloadContractsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 1. 下载合同模板
2. 多个合同id用逗号分隔
 [MFA enabled] */
func (c *CloudsignClient) DownloadTemplates(request *cloudsign.DownloadTemplatesRequest) (*cloudsign.DownloadTemplatesResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &cloudsign.DownloadTemplatesResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 验签已签章合同 */
func (c *CloudsignClient) VerifyContract(request *cloudsign.VerifyContractRequest) (*cloudsign.VerifyContractResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &cloudsign.VerifyContractResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 上传合同模板 */
func (c *CloudsignClient) UploadTemplate(request *cloudsign.UploadTemplateRequest) (*cloudsign.UploadTemplateResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &cloudsign.UploadTemplateResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 删除合同模板 [MFA enabled] */
func (c *CloudsignClient) DeleteTemplate(request *cloudsign.DeleteTemplateRequest) (*cloudsign.DeleteTemplateResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &cloudsign.DeleteTemplateResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 删除已签章的合同 [MFA enabled] */
func (c *CloudsignClient) DeleteContract(request *cloudsign.DeleteContractRequest) (*cloudsign.DeleteContractResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &cloudsign.DeleteContractResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 禁用合同存管 */
func (c *CloudsignClient) DisableContractSave(request *cloudsign.DisableContractSaveRequest) (*cloudsign.DisableContractSaveResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &cloudsign.DisableContractSaveResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 签章系统加密密钥 [MFA enabled] */
func (c *CloudsignClient) SetKmsKeyId(request *cloudsign.SetKmsKeyIdRequest) (*cloudsign.SetKmsKeyIdResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &cloudsign.SetKmsKeyIdResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 删除印章 [MFA enabled] */
func (c *CloudsignClient) DeleteStamp(request *cloudsign.DeleteStampRequest) (*cloudsign.DeleteStampResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &cloudsign.DeleteStampResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 获取印章列表 */
func (c *CloudsignClient) DescribeStampList(request *cloudsign.DescribeStampListRequest) (*cloudsign.DescribeStampListResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &cloudsign.DescribeStampListResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

