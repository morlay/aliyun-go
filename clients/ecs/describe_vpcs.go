package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeVpcsRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	VpcId                string `position:"Query" name:"VpcId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int    `position:"Query" name:"PageSize"`
	IsDefault            string `position:"Query" name:"IsDefault"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
}

func (r DescribeVpcsRequest) Invoke(client *sdk.Client) (response *DescribeVpcsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeVpcsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeVpcs", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		DescribeVpcsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeVpcsResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeVpcsResponse struct {
	RequestId  string
	TotalCount int
	PageNumber int
	PageSize   int
	Vpcs       DescribeVpcsVpcList
}

type DescribeVpcsVpc struct {
	VpcId        string
	RegionId     string
	Status       string
	VpcName      string
	CreationTime string
	CidrBlock    string
	VRouterId    string
	Description  string
	IsDefault    bool
	VSwitchIds   DescribeVpcsVSwitchIdList
	UserCidrs    DescribeVpcsUserCidrList
}

type DescribeVpcsVpcList []DescribeVpcsVpc

func (list *DescribeVpcsVpcList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeVpcsVpc)
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

type DescribeVpcsVSwitchIdList []string

func (list *DescribeVpcsVSwitchIdList) UnmarshalJSON(data []byte) error {
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

type DescribeVpcsUserCidrList []string

func (list *DescribeVpcsUserCidrList) UnmarshalJSON(data []byte) error {
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
