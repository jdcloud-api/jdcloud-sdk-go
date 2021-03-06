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

package models


type DownloadCertDesc struct {

    /* 证书Id (Optional) */
    CertId string `json:"certId"`

    /* 证书名称 (Optional) */
    CertName string `json:"certName"`

    /* 私钥 (Optional) */
    KeyFile string `json:"keyFile"`

    /* 证书 (Optional) */
    CertFile string `json:"certFile"`

    /* 对私钥文件使用sha256算法计算的摘要信息 (Optional) */
    Digest string `json:"digest"`

    /* 中间证书 (Optional) */
    CaCertFile string `json:"caCertFile"`

    /* 证书应用服务器类型 (Optional) */
    ServerType string `json:"serverType"`

    /* 证书加密密码 (Optional) */
    CertEncryptePassword string `json:"certEncryptePassword"`

    /* 域名 (Optional) */
    CommonName string `json:"commonName"`
}
