package vpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeVpcAttributeRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	VpcId                string `position:"Query" name:"VpcId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	IsDefault            string `position:"Query" name:"IsDefault"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeVpcAttributeRequest) Invoke(client *sdk.Client) (response *DescribeVpcAttributeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeVpcAttributeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Vpc", "2016-04-28", "DescribeVpcAttribute", "vpc", "")

	resp := struct {
		*responses.BaseResponse
		DescribeVpcAttributeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeVpcAttributeResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeVpcAttributeResponse struct {
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
	CloudResources     DescribeVpcAttributeCloudResourceSetTypeList
	VSwitchIds         DescribeVpcAttributeVSwitchIdList
	UserCidrs          DescribeVpcAttributeUserCidrList
}

type DescribeVpcAttributeCloudResourceSetType struct {
	ResourceType  string
	ResourceCount int
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
