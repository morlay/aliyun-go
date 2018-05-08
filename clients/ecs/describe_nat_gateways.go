package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeNatGatewaysRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	VpcId                string `position:"Query" name:"VpcId"`
	PageSize             int    `position:"Query" name:"PageSize"`
	NatGatewayId         string `position:"Query" name:"NatGatewayId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
}

func (req *DescribeNatGatewaysRequest) Invoke(client *sdk.Client) (resp *DescribeNatGatewaysResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeNatGateways", "ecs", "")
	resp = &DescribeNatGatewaysResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeNatGatewaysResponse struct {
	responses.BaseResponse
	RequestId   common.String
	TotalCount  common.Integer
	PageNumber  common.Integer
	PageSize    common.Integer
	NatGateways DescribeNatGatewaysNatGatewayList
}

type DescribeNatGatewaysNatGateway struct {
	NatGatewayId        common.String
	RegionId            common.String
	Name                common.String
	Description         common.String
	VpcId               common.String
	Spec                common.String
	InstanceChargeType  common.String
	BusinessStatus      common.String
	CreationTime        common.String
	Status              common.String
	ForwardTableIds     DescribeNatGatewaysForwardTableIdList
	BandwidthPackageIds DescribeNatGatewaysBandwidthPackageIdList
}

type DescribeNatGatewaysNatGatewayList []DescribeNatGatewaysNatGateway

func (list *DescribeNatGatewaysNatGatewayList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeNatGatewaysNatGateway)
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

type DescribeNatGatewaysForwardTableIdList []common.String

func (list *DescribeNatGatewaysForwardTableIdList) UnmarshalJSON(data []byte) error {
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

type DescribeNatGatewaysBandwidthPackageIdList []common.String

func (list *DescribeNatGatewaysBandwidthPackageIdList) UnmarshalJSON(data []byte) error {
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
