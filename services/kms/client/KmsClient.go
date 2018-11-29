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
    kms "github.com/jdcloud-api/jdcloud-sdk-go/services/kms/apis"
    "encoding/json"
    "errors"
)

type KmsClient struct {
    core.JDCloudClient
}

func NewKmsClient(credential *core.Credential) *KmsClient {
    if credential == nil {
        return nil
    }

    config := core.NewConfig()
    config.SetEndpoint("kms.jdcloud-api.com")

    return &KmsClient{
        core.JDCloudClient{
            Credential:  *credential,
            Config:      *config,
            ServiceName: "kms",
            Revision:    "0.1.1",
            Logger:      core.NewDefaultLogger(core.LogInfo),
        }}
}

func (c *KmsClient) SetConfig(config *core.Config) {
    c.Config = *config
}

func (c *KmsClient) SetLogger(logger core.Logger) {
    c.Logger = logger
}

/* 创建机密 */
func (c *KmsClient) CreateSecret(request *kms.CreateSecretRequest) (*kms.CreateSecretResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &kms.CreateSecretResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 获取指定机密版本的详细信息 */
func (c *KmsClient) DescribeSecretVersionInfo(request *kms.DescribeSecretVersionInfoRequest) (*kms.DescribeSecretVersionInfoResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &kms.DescribeSecretVersionInfoResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 禁用机密 */
func (c *KmsClient) DisableSecret(request *kms.DisableSecretRequest) (*kms.DisableSecretResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &kms.DisableSecretResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 禁用指定版本机密 */
func (c *KmsClient) DisableSecretVersion(request *kms.DisableSecretVersionRequest) (*kms.DisableSecretVersionResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &kms.DisableSecretVersionResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 获取版本详情 */
func (c *KmsClient) DescribeKeyDetail(request *kms.DescribeKeyDetailRequest) (*kms.DescribeKeyDetailResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &kms.DescribeKeyDetailResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 创建一个CMK（用户主密钥），默认为启用状态 */
func (c *KmsClient) CreateKey(request *kms.CreateKeyRequest) (*kms.CreateKeyResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &kms.CreateKeyResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 使用密钥对数据进行加密 */
func (c *KmsClient) Encrypt(request *kms.EncryptRequest) (*kms.EncryptResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &kms.EncryptResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 启用指定版本密钥 */
func (c *KmsClient) EnableKeyVersion(request *kms.EnableKeyVersionRequest) (*kms.EnableKeyVersionResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &kms.EnableKeyVersionResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 修改机密描述 */
func (c *KmsClient) UpdateSecret(request *kms.UpdateSecretRequest) (*kms.UpdateSecretResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &kms.UpdateSecretResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 取消删除指定版本密钥 */
func (c *KmsClient) CancelKeyVersionDeletion(request *kms.CancelKeyVersionDeletionRequest) (*kms.CancelKeyVersionDeletionResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &kms.CancelKeyVersionDeletionResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 从KMS中获取一对数据密钥的明文/密文 */
func (c *KmsClient) GenerateDataKey(request *kms.GenerateDataKeyRequest) (*kms.GenerateDataKeyResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &kms.GenerateDataKeyResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 修改机密指定版本配置 */
func (c *KmsClient) UpdateSecretVersion(request *kms.UpdateSecretVersionRequest) (*kms.UpdateSecretVersionResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &kms.UpdateSecretVersionResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 修改密钥配置，包括key的名称、用途、是否自动轮换和轮换周期等 */
func (c *KmsClient) UpdateKeyDescription(request *kms.UpdateKeyDescriptionRequest) (*kms.UpdateKeyDescriptionResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &kms.UpdateKeyDescriptionResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 禁用当前状态为`已启用`的密钥 */
func (c *KmsClient) DisableKey(request *kms.DisableKeyRequest) (*kms.DisableKeyResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &kms.DisableKeyResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 获取机密详情 */
func (c *KmsClient) DescribeSecretVersionList(request *kms.DescribeSecretVersionListRequest) (*kms.DescribeSecretVersionListResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &kms.DescribeSecretVersionListResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 计划在以后的是个时间点删除密钥，默认为7天 */
func (c *KmsClient) ScheduleKeyDeletion(request *kms.ScheduleKeyDeletionRequest) (*kms.ScheduleKeyDeletionResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &kms.ScheduleKeyDeletionResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 导入机密 */
func (c *KmsClient) ImportSecret(request *kms.ImportSecretRequest) (*kms.ImportSecretResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &kms.ImportSecretResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 立即轮换密钥，自动轮换周期顺延 */
func (c *KmsClient) KeyRotation(request *kms.KeyRotationRequest) (*kms.KeyRotationResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &kms.KeyRotationResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 获取密钥详情 */
func (c *KmsClient) DescribeKey(request *kms.DescribeKeyRequest) (*kms.DescribeKeyResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &kms.DescribeKeyResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 删除指定版本机密 */
func (c *KmsClient) DeleteSecretVersion(request *kms.DeleteSecretVersionRequest) (*kms.DeleteSecretVersionResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &kms.DeleteSecretVersionResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 获取密钥列表 */
func (c *KmsClient) DescribeKeyList(request *kms.DescribeKeyListRequest) (*kms.DescribeKeyListResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &kms.DescribeKeyListResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 创建机密新的版本，默认为已启用状态 */
func (c *KmsClient) CreateSecretVersion(request *kms.CreateSecretVersionRequest) (*kms.CreateSecretVersionResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &kms.CreateSecretVersionResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 启用当前状态为`已禁用`的密钥 */
func (c *KmsClient) EnableKey(request *kms.EnableKeyRequest) (*kms.EnableKeyResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &kms.EnableKeyResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 使用密钥对数据进行解密 */
func (c *KmsClient) Decrypt(request *kms.DecryptRequest) (*kms.DecryptResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &kms.DecryptResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 计划在以后的是个时间点删除指定版本密钥，默认为7天 */
func (c *KmsClient) ScheduleKeyVersionDeletion(request *kms.ScheduleKeyVersionDeletionRequest) (*kms.ScheduleKeyVersionDeletionResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &kms.ScheduleKeyVersionDeletionResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 取消删除密钥 */
func (c *KmsClient) CancelKeyDeletion(request *kms.CancelKeyDeletionRequest) (*kms.CancelKeyDeletionResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &kms.CancelKeyDeletionResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 启用指定版本机密 */
func (c *KmsClient) EnableSecretVersion(request *kms.EnableSecretVersionRequest) (*kms.EnableSecretVersionResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &kms.EnableSecretVersionResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 删除机密 */
func (c *KmsClient) DeleteSecret(request *kms.DeleteSecretRequest) (*kms.DeleteSecretResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &kms.DeleteSecretResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 禁用指定版本密钥 */
func (c *KmsClient) DisableKeyVersion(request *kms.DisableKeyVersionRequest) (*kms.DisableKeyVersionResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &kms.DisableKeyVersionResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 启用机密 */
func (c *KmsClient) EnableSecret(request *kms.EnableSecretRequest) (*kms.EnableSecretResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &kms.EnableSecretResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 导出机密 */
func (c *KmsClient) ExportSecret(request *kms.ExportSecretRequest) (*kms.ExportSecretResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &kms.ExportSecretResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 获取机密列表 */
func (c *KmsClient) DescribeSecretList(request *kms.DescribeSecretListRequest) (*kms.DescribeSecretListResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &kms.DescribeSecretListResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

