package cloudapi

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId     string
	GroupId       string
	GroupName     string
	SubDomain     string
	Description   string
	CreatedTime   string
	ModifiedTime  string
	RegionId      string
	Status        string
	BillingStatus string
	IllegalStatus string
	TrafficLimit  int
	CustomDomains DescribeApiGroupDomainItemList
	StageItems    DescribeApiGroupStageInfoList
}

type DescribeApiGroupDomainItem struct {
	DomainName          string
	CertificateId       string
	CertificateName     string
	DomainCNAMEStatus   string
	DomainBindingStatus string
	DomainLegalStatus   string
	DomainRemark        string
}

type DescribeApiGroupStageInfo struct {
	StageId     string
	StageName   string
	Description string
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
