package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeVpcsRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	VpcId                string `position:"Query" name:"VpcId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int    `position:"Query" name:"PageSize"`
	IsDefault            string `position:"Query" name:"IsDefault"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
}

func (req *DescribeVpcsRequest) Invoke(client *sdk.Client) (resp *DescribeVpcsResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeVpcs", "ecs", "")
	resp = &DescribeVpcsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeVpcsResponse struct {
	responses.BaseResponse
	RequestId  common.String
	TotalCount common.Integer
	PageNumber common.Integer
	PageSize   common.Integer
	Vpcs       DescribeVpcsVpcList
}

type DescribeVpcsVpc struct {
	VpcId        common.String
	RegionId     common.String
	Status       common.String
	VpcName      common.String
	CreationTime common.String
	CidrBlock    common.String
	VRouterId    common.String
	Description  common.String
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

type DescribeVpcsVSwitchIdList []common.String

func (list *DescribeVpcsVSwitchIdList) UnmarshalJSON(data []byte) error {
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

type DescribeVpcsUserCidrList []common.String

func (list *DescribeVpcsUserCidrList) UnmarshalJSON(data []byte) error {
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
