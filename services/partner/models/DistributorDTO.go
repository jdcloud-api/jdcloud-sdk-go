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


type DistributorDTO struct {

    /* 渠道商ID (Optional) */
    DistributorId string `json:"distributorId"`

    /* 渠道商名称 (Optional) */
    DistributorName string `json:"distributorName"`

    /* pin (Optional) */
    Pin string `json:"pin"`

    /* 合同编号 (Optional) */
    ContractNo string `json:"contractNo"`

    /* 营业执照号 (Optional) */
    BusinessLicense string `json:"businessLicense"`

    /* 法定代表人 (Optional) */
    LegalRepresentative string `json:"legalRepresentative"`

    /* 营业执照图片 (Optional) */
    BusinessLicensePic string `json:"businessLicensePic"`

    /* 主营业务描述 (Optional) */
    BusinessDesc string `json:"businessDesc"`

    /* 办公地址 (Optional) */
    WorkAddress string `json:"workAddress"`

    /* 联系人姓名 (Optional) */
    Contracter string `json:"contracter"`

    /* 联系人电话 (Optional) */
    Tel string `json:"tel"`

    /* 邮箱 (Optional) */
    Email string `json:"email"`

    /* 所属地域 (Optional) */
    Region string `json:"region"`

    /* 入驻日期(一级渠道商手工录入、二级渠道商审批通过日期) (Optional) */
    SettleTime string `json:"settleTime"`

    /* 状态(0 审批中、2驳回、1 已入驻、3已停止合作) (Optional) */
    Status int `json:"status"`

    /* 驳回原因 (Optional) */
    Reason string `json:"reason"`

    /* 级次(0一级、1 二级) (Optional) */
    DistributorLevel int `json:"distributorLevel"`

    /* 渠道商类型(0合作伙伴、1 渠道代理) (Optional) */
    DistributorType int `json:"distributorType"`

    /* 上级渠道商ID (Optional) */
    ParentDistributorId string `json:"parentDistributorId"`

    /* 银行开户名 (Optional) */
    BankCompanyName string `json:"bankCompanyName"`

    /* 银行账户 (Optional) */
    BankCardNo string `json:"bankCardNo"`

    /* 开户行支行名称 (Optional) */
    BankBranchName string `json:"bankBranchName"`

    /* 开户行支行联行号 (Optional) */
    BankBranchNo string `json:"bankBranchNo"`

    /* 税务识别号 (Optional) */
    TaxNumber string `json:"taxNumber"`

    /* 发票抬头 (Optional) */
    InvoiceTitle string `json:"invoiceTitle"`

    /* 1 全强代付,2 全自付, 3 自选 (Optional) */
    PayType int `json:"payType"`

    /* 合同主体 (Optional) */
    ContractSubject string `json:"contractSubject"`

    /* 所属部门(0企业线、1政府线) (Optional) */
    Dept int `json:"dept"`

    /* 是否需要返还（0需要1不需要） (Optional) */
    ReturnFlag int `json:"returnFlag"`

    /* 返还政策ID (Optional) */
    ReturnPolicyId string `json:"returnPolicyId"`

    /*  (Optional) */
    DistributorPolicyList []DistributorPolicyDTO `json:"distributorPolicyList"`

    /*  (Optional) */
    SubDistributorPolicyList []DistributorPolicyDTO `json:"subDistributorPolicyList"`

    /*  (Optional) */
    DistributorProductList []DistributorProductDTO `json:"distributorProductList"`

    /* 结算周期类型（1月、2季度、3年、4天、5周） (Optional) */
    CircleType int `json:"circleType"`

    /* 服务商返还方式（1现金2代金券） (Optional) */
    ReturnMode int `json:"returnMode"`

    /* 是否有下级服务商（0有1不没有） (Optional) */
    SubFlag int `json:"subFlag"`

    /* 下级服务商是否需要返还（0需要1不需要） (Optional) */
    SubReturnFlag int `json:"subReturnFlag"`

    /* 下级服务商返还政策ID (Optional) */
    SubReturnPolicyId string `json:"subReturnPolicyId"`

    /* 结算周期类型（1月、2季度、3年、4天、5周） (Optional) */
    SubCircleType int `json:"subCircleType"`

    /* 下级服务商返还方式（1现金2代金券） (Optional) */
    SubReturnMode int `json:"subReturnMode"`

    /* 京东云负责人(京东云人员erp或名称) (Optional) */
    Erp string `json:"erp"`

    /* 京东云负责人姓名 (Optional) */
    ErpName string `json:"erpName"`

    /* 京东云负责人部门 (Optional) */
    ErpDept string `json:"erpDept"`

    /* 服务商级别 (Optional) */
    Grade int `json:"grade"`

    /* 服务期限（开始时间） (Optional) */
    EffectiveDate string `json:"effectiveDate"`

    /* 服务期限（结束时间） (Optional) */
    ExpirationDate string `json:"expirationDate"`

    /* 电子合同编号 (Optional) */
    EContractNo string `json:"eContractNo"`

    /* EBS合同编号 (Optional) */
    EbsContractNo string `json:"ebsContractNo"`

    /* 是否享受折扣(1享受2不享受) (Optional) */
    DiscountFlag int `json:"discountFlag"`
}
