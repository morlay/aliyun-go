package vpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId          common.String
	VpcId              common.String
	RegionId           common.String
	Status             common.String
	VpcName            common.String
	CreationTime       common.String
	CidrBlock          common.String
	VRouterId          common.String
	Description        common.String
	IsDefault          bool
	ClassicLinkEnabled bool
	ResourceGroupId    common.String
	AssociatedCens     DescribeVpcAttributeAssociatedCenList
	CloudResources     DescribeVpcAttributeCloudResourceSetTypeList
	VSwitchIds         DescribeVpcAttributeVSwitchIdList
	UserCidrs          DescribeVpcAttributeUserCidrList
}

type DescribeVpcAttributeAssociatedCen struct {
	CenId      common.String
	CenOwnerId common.Long
	CenStatus  common.String
}

type DescribeVpcAttributeCloudResourceSetType struct {
	ResourceType  common.String
	ResourceCount common.Integer
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

type DescribeVpcAttributeVSwitchIdList []common.String

func (list *DescribeVpcAttributeVSwitchIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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

type DescribeVpcAttributeUserCidrList []common.String

func (list *DescribeVpcAttributeUserCidrList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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
