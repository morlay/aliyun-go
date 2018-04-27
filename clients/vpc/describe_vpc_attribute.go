package vpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeVpcAttributeRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	VpcId                string `position:"Query" name:"VpcId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	IsDefault            string `position:"Query" name:"IsDefault"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeVpcAttributeRequest) Invoke(client *sdk.Client) (resp *DescribeVpcAttributeResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "DescribeVpcAttribute", "vpc", "")
	resp = &DescribeVpcAttributeResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeVpcAttributeResponse struct {
	responses.BaseResponse
	RequestId          string
	VpcId              string
	RegionId           string
	Status             string
	VpcName            string
	CreationTime       string
	CidrBlock          string
	VRouterId          string
	Description        string
	IsDefault          bool
	ClassicLinkEnabled bool
	ResourceGroupId    string
	AssociatedCens     DescribeVpcAttributeAssociatedCenList
	CloudResources     DescribeVpcAttributeCloudResourceSetTypeList
	VSwitchIds         DescribeVpcAttributeVSwitchIdList
	UserCidrs          DescribeVpcAttributeUserCidrList
}

type DescribeVpcAttributeAssociatedCen struct {
	CenId      string
	CenOwnerId int64
	CenStatus  string
}

type DescribeVpcAttributeCloudResourceSetType struct {
	ResourceType  string
	ResourceCount int
}

type DescribeVpcAttributeAssociatedCenList []DescribeVpcAttributeAssociatedCen

func (list *DescribeVpcAttributeAssociatedCenList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeVpcAttributeAssociatedCen)
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

type DescribeVpcAttributeCloudResourceSetTypeList []DescribeVpcAttributeCloudResourceSetType

func (list *DescribeVpcAttributeCloudResourceSetTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeVpcAttributeCloudResourceSetType)
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

type DescribeVpcAttributeVSwitchIdList []string

func (list *DescribeVpcAttributeVSwitchIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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

type DescribeVpcAttributeUserCidrList []string

func (list *DescribeVpcAttributeUserCidrList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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
