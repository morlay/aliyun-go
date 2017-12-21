package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeNatGatewaysRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	VpcId                string `position:"Query" name:"VpcId"`
	PageSize             int    `position:"Query" name:"PageSize"`
	NatGatewayId         string `position:"Query" name:"NatGatewayId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
}

func (r DescribeNatGatewaysRequest) Invoke(client *sdk.Client) (response *DescribeNatGatewaysResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeNatGatewaysRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeNatGateways", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		DescribeNatGatewaysResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeNatGatewaysResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeNatGatewaysResponse struct {
	RequestId   string
	TotalCount  int
	PageNumber  int
	PageSize    int
	NatGateways DescribeNatGatewaysNatGatewayList
}

type DescribeNatGatewaysNatGateway struct {
	NatGatewayId        string
	RegionId            string
	Name                string
	Description         string
	VpcId               string
	Spec                string
	InstanceChargeType  string
	BusinessStatus      string
	CreationTime        string
	Status              string
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

type DescribeNatGatewaysForwardTableIdList []string

func (list *DescribeNatGatewaysForwardTableIdList) UnmarshalJSON(data []byte) error {
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

type DescribeNatGatewaysBandwidthPackageIdList []string

func (list *DescribeNatGatewaysBandwidthPackageIdList) UnmarshalJSON(data []byte) error {
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
