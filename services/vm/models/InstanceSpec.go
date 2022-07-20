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

import vpc "github.com/lshuining/jdcloud-sdk-go/services/vpc/models"
import charge "github.com/lshuining/jdcloud-sdk-go/services/charge/models"
import disk "github.com/lshuining/jdcloud-sdk-go/services/disk/models"

type InstanceSpec struct {

    /* 高可用组ID。指定此参数后，将默认使用高可用组关联的实例模板创建实例，实例模板中的参数不可覆盖替换。实例模板以外的参数（内网IPv4/Ipv6分配方式、名称、描述、标签）可指定。
 (Optional) */
    AgId *string `json:"agId"`

    /* 实例模板ID。指定此参数后，如实例模板中参数不另行指定将默认以模板配置创建实例，如指定则以指定值为准创建。
指定 `agId` 时此参数无效。
 (Optional) */
    InstanceTemplateId *string `json:"instanceTemplateId"`

    /* 实例所属的可用区。
如不指定 `agId` 以使用高可用组设置的可用区，此参数为必选。
 (Optional) */
    Az *string `json:"az"`

    /* 实例规格。可通过 [DescribeInstanceTypes](https://docs.jdcloud.com/virtual-machines/api/describeinstancetypes) 接口查询各地域及可用区下的规格售卖情况。
如不指定 `agId` 或 `instanceTemplateId` 以使用实例模板中配置的规格，此参数为必选。
 (Optional) */
    InstanceType *string `json:"instanceType"`

    /* 镜像ID。可通过 [DescribeImages](https://docs.jdcloud.com/virtual-machines/api/describeimages) 接口获得指定地域的镜像信息。
如不指定 `agId` 或 `instanceTemplateId` 以使用实例模板中配置的镜像，此参数为必选。
 (Optional) */
    ImageId *string `json:"imageId"`

    /* 实例名称。长度为2\~128个字符，只允许中文、数字、大小写字母、英文下划线（\_）、连字符（-）及点（.），不能以（.）作为首尾。
批量创建多台实例时，可在name中非首位位置以\[start_number]格式来设置有序name。start_number为起始序号，其位数代表编号字符位数，范围：\[0,9999]。
例如：name设置为“instance-\[001]-ops”，则第一台主机name为“instance-001o-ps”，第二台主机name为“instance-002-ops”。详情参见[为实例设置有序名称及Hostname]()。
  */
    Name string `json:"name"`

    /* 实例hostname。若不指定hostname，则默认以实例名称 `name` 作为hostname，但是会以RFC 952和RFC 1123命名规范做一定转义。
**Windows系统**：长度为2\~15个字符，允许大小写字母、数字或连字符（-），不能以连字符（-）开头或结尾，不能连续使用连字符（-），也不能全部使用数字。不支持点号（.）。
**Linux系统**：长度为2-64个字符，允许支持多个点号，点之间为一段，每段允许使用大小写字母、数字或连字符（-），但不能连续使用点号（.）或连字符（-），不能以点号（.）或连字符（-）开头或结尾。
批量创建多台实例时，可在hostname中非首位位置以\[start_number]格式来设置有序hostname。start_number为起始序号，其位数代表编号字符位数，范围：\[0,9999]。。例如：hostname设置为“instance-\[000]-ops”，则第一台主机hostname为“instance-000-ops”，第二台主机hostname为“instance-001-ops”。详情参见[为实例设置有序名称及Hostname]()。
批量创建时若不指定序号，则会默认追加从1开始的数字，例如批量创建两台虚拟机，且指定hostname是test，则hostname默认是test1，test2。
 (Optional) */
    Hostname *string `json:"hostname"`

    /* 实例密码。可用于SSH登录和VNC登录。长度为8\~30个字符，必须同时包含大、小写英文字母、数字和特殊符号中的三类字符。特殊符号包括：\(\)\`~!@#$%^&\*\_-+=\|{}\[ ]:";'<>,.?/，更多密码输入要求请参见 [公共参数规范](https://docs.jdcloud.com/virtual-machines/api/general_parameters)。
如指定密钥且 `passwordAuth` 设置为 `true` ，则密码不会生成注入，否则即使不指定密码系统也将默认自动生成随机密码，并以短信和邮件通知。
 (Optional) */
    Password *string `json:"password"`

    /* 密钥对名称。仅Linux系统下该参数生效，当前仅支持输入单个密钥。
 (Optional) */
    KeyNames []string `json:"keyNames"`

    /* 主网卡主IP关联的弹性公网IP配置。
 (Optional) */
    ElasticIp *vpc.ElasticIpSpec `json:"elasticIp"`

    /* 主网卡配置。
 (Optional) */
    PrimaryNetworkInterface *InstanceNetworkInterfaceAttachmentSpec `json:"primaryNetworkInterface"`

    /* 系统盘配置。
 (Optional) */
    SystemDisk *InstanceDiskAttachmentSpec `json:"systemDisk"`

    /* 数据盘配置。单实例最多可挂载云硬盘（系统盘+数据盘）的数量受实例规格的限制。
 (Optional) */
    DataDisks []InstanceDiskAttachmentSpec `json:"dataDisks"`

    /* 计费配置。
云主机不支持按用量方式计费，默认为按配置计费。
打包创建数据盘的情况下，数据盘的计费方式只能与云主机保持一致。
打包创建弹性公网IP的情况下，若公网IP的计费方式没有指定为按用量计费，那么公网IP计费方式只能与云主机保持一致。
 (Optional) */
    Charge *charge.ChargeSpec `json:"charge"`

    /* 用户自定义元数据。以key-value键值对形式指定，可在实例系统内通过元数据服务查询获取。最多支持40对键值对，且key不超过256字符，value不超过16KB，不区分大小写。
注意：key不要以连字符(-)结尾，否则此key不生效。
 (Optional) */
    Metadata []Metadata `json:"metadata"`

    /* 自定义脚本。目前仅支持启动脚本，即 `launch-script`，须 `base64` 编码且编码前数据长度不能超过16KB。
**linux系统**：支持 `bash` 和 `python`，编码前须分别以 `#!/bin/bash` 和 `#!/usr/bin/env python` 作为内容首行。
**Windows系统**：支持 `bat` 和 `powershell`，编码前须分别以 `<cmd></cmd>和<powershell></powershell>` 作为内容首、尾行。
 (Optional) */
    Userdata []Userdata `json:"userdata"`

    /* 实例描述。256字符以内。
 (Optional) */
    Description *string `json:"description"`

    /* 使用实例模板创建实例时，如模板中已设置密码，期望不使用该密码而由系统自动生成时，可通过此参数（`true`）实现。
可选值：
`true`：不使用实例模板中配置的密码。
`false`：使用实例模板中配置的密码。
仅在未指定 `agId` 且指定 `instanceTemplateId`，且 `password` 为空时，此参数(`true`)生效。
 (Optional) */
    NoPassword *bool `json:"noPassword"`

    /* 使用实例模板创建实例时，如模板中已设置密钥，期望不使用该密钥仅使用密码作为登录凭证时，可通过此参数（`true`）实现。
仅在未指定 `agId` 且指定 `instanceTemplateId`，且 `keyNames` 为空时，此参数(`true`)生效。
 (Optional) */
    NoKeyNames *bool `json:"noKeyNames"`

    /* 使用实例模板创建实例时，如模板中已设置弹性公网IP，期望不绑定弹性公网IP时，可通过此参数（`true`）实现。
仅在未指定 `agId` 且指定 `instanceTemplateId`，且 `elasticIp` 为空时，此参数(`true`)生效。
 (Optional) */
    NoElasticIp *bool `json:"noElasticIp"`

    /* 自定义实例标签。以key-value键值对形式指定，最多支持10个标签。key不能以 "jrn:" 或“jdc-”开头，仅支持中文、大/小写英文、数字及如下符号：`\_.,:\/=+-@`。
 (Optional) */
    UserTags []disk.Tag `json:"userTags"`

    /* 停机不计费模式。该参数仅对按配置计费且系统盘为云硬盘的实例生效，并且不是专有宿主机中的实例。配置停机不计费且停机后，实例部分将停止计费，且释放实例自身包含的资源（CPU/内存/GPU/本地数据盘）。
可选值：
`keepCharging`（默认值）：停机后保持计费，不释放资源。
`stopCharging`：停机后停止计费，释放实例资源。
 (Optional) */
    ChargeOnStopped *string `json:"chargeOnStopped"`

    /* 自动任务策略ID。
 (Optional) */
    AutoImagePolicyId *string `json:"autoImagePolicyId"`

    /* 允许SSH密码登录。
可选值：
`yes`（默认值）：允许SSH密码登录。
`no`：禁止SSH密码登录。
仅在指定密钥时此参数有效，指定此参数后密码即使输入也将被忽略，同时会在系统内禁用SSH密码登录。
 (Optional) */
    PasswordAuth *string `json:"passwordAuth"`

    /* 使用镜像中的登录凭证，无须再指定密码或密钥（指定无效）。
可选值：
`yes`：使用镜像登录凭证。
`no`（默认值）：不使用镜像登录凭证。
仅使用私有或共享镜像时此参数有效。
 (Optional) */
    ImageInherit *string `json:"imageInherit"`

    /* 资源组ID (Optional) */
    ResourceGroupId *string `json:"resourceGroupId"`
}
