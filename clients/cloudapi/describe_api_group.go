package cloudapi

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeApiGroupRequest struct {
	requests.RpcRequest
	GroupId string `position:"Query" name:"GroupId"`
}

func (req *DescribeApiGroupRequest) Invoke(client *sdk.Client) (resp *DescribeApiGroupResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeApiGroup", "apigateway", "")
	resp = &DescribeApiGroupResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeApiGroupResponse struct {
	responses.BaseResponse
	RequestId     common.String
	GroupId       common.String
	GroupName     common.String
	SubDomain     common.String
	Description   common.String
	CreatedTime   common.String
	ModifiedTime  common.String
	RegionId      common.String
	Status        common.String
	BillingStatus common.String
	IllegalStatus common.String
	TrafficLimit  common.Integer
	CustomDomains DescribeApiGroupDomainItemList
	StageItems    DescribeApiGroupStageInfoList
}

type DescribeApiGroupDomainItem struct {
	DomainName          common.String
	CertificateId       common.String
	CertificateName     common.String
	DomainCNAMEStatus   common.String
	DomainBindingStatus common.String
	DomainLegalStatus   common.String
	DomainRemark        common.String
}

type DescribeApiGroupStageInfo struct {
	StageId     common.String
	StageName   common.String
	Description common.String
}

type DescribeApiGroupDomainItemList []DescribeApiGroupDomainItem

func (list *DescribeApiGroupDomainItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApiGroupDomainItem)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}

type DescribeApiGroupStageInfoList []DescribeApiGroupStageInfo

func (list *DescribeApiGroupStageInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApiGroupStageInfo)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
