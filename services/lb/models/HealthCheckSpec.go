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


type HealthCheckSpec struct {

    /* 健康检查协议 <br>【alb、nlb】取值为Http, Tcp, Icmp(仅支持alb/nlb的Udp的Backend) <br>【dnlb】取值为Tcp, Icmp(仅支持dnlb的Udp的Backend)  */
    Protocol string `json:"protocol"`

    /* 健康阀值，取值范围为[1,5]，默认为3 (Optional) */
    HealthyThresholdCount *int `json:"healthyThresholdCount"`

    /* 不健康阀值，取值范围为[1,5], 默认为3 (Optional) */
    UnhealthyThresholdCount *int `json:"unhealthyThresholdCount"`

    /* 响应超时时间, 取值范围为[2,60]，默认为3s (Optional) */
    CheckTimeoutSeconds *int `json:"checkTimeoutSeconds"`

    /* 健康检查间隔, 范围为[5,300], 默认为5s (Optional) */
    IntervalSeconds *int `json:"intervalSeconds"`

    /* 健康检查的目标端口, 取值范围为[0,65535], 默认为0，默认端口为每个后端实例接收负载均衡流量的端口，Icmp类型不支持配置端口 (Optional) */
    Port *int `json:"port"`

    /* 健康检查的协议类型 HTTP1.0/HTTP1.1 默认选择HTTP1.0 (Optional) */
    HttpVersion *string `json:"httpVersion"`

    /* 健康检查的目标域名，仅支持HTTP协议。支持输入域名和IP地址。如果输入域名，仅支持大小写字母、数字、英文中划线"-"和点"."，不区分大小写，且不超过255个字符。默认为空，表示健康检查不携带域名 (Optional) */
    HttpDomain *string `json:"httpDomain"`

    /* 健康检查的目标路径，仅支持HTTP协议。必须以"/"开头，支持大小写字母、数字、汉字和英文字符-/.%?#&_;~!()*[]@^:',+=<>{}。%后仅支持输入URL编码后字符，且不超过1000个字符 (Optional) */
    HttpPath *string `json:"httpPath"`

    /* 健康检查的期望状态码，仅支持HTTP协议。检查来自后端目标服务器的成功响应时，要使用的HTTP状态码。您可以指定单个数值（例如："200"，取值范围200-499）、一段连续数值（例如："201-205"，取值范围范围200-499，且前面的参数小于后面）和一类连续数值缩写（例如："3xx"，等价于"300-399"，取值范围2xx、3xx和4xx）。多个数值之间通过","分割（例如："200,202-207,302,4xx"）。目前仅支持2xx、3xx、4xx。仅支持HTTP协议，默认为[2xx、3xx] (Optional) */
    HttpCode []string `json:"httpCode"`
}
