package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeRecommendInstanceTypeRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Channel              string `position:"Query" name:"Channel"`
	NetworkType          string `position:"Query" name:"NetworkType"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Operator             string `position:"Query" name:"Operator"`
	Token                string `position:"Query" name:"Token"`
	Scene                string `position:"Query" name:"Scene"`
	InstanceType         string `position:"Query" name:"InstanceType"`
	ProxyId              string `position:"Query" name:"ProxyId"`
}

func (req *DescribeRecommendInstanceTypeRequest) Invoke(client *sdk.Client) (resp *DescribeRecommendInstanceTypeResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeRecommendInstanceType", "ecs", "")
	resp = &DescribeRecommendInstanceTypeResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeRecommendInstanceTypeResponse struct {
	responses.BaseResponse
	RequestId common.String
	Data      DescribeRecommendInstanceTypeRecommendInstanceTypeList
}

type DescribeRecommendInstanceTypeRecommendInstanceType struct {
	RegionNo      common.String
	CommodityCode common.String
	Scene         common.String
	Zones         DescribeRecommendInstanceTypeZoneList
	InstanceType  DescribeRecommendInstanceTypeInstanceType
}

type DescribeRecommendInstanceTypeZone struct {
	ZoneNo       common.String
	NetworkTypes DescribeRecommendInstanceTypeNetworkTypeList
}

type DescribeRecommendInstanceTypeInstanceType struct {
	Generation         common.String
	InstanceTypeFamily common.String
	InstanceType       common.String
	SupportIoOptimized common.String
	Cores              common.Integer
	Memory             common.Integer
}

type DescribeRecommendInstanceTypeRecommendInstanceTypeList []DescribeRecommendInstanceTypeRecommendInstanceType

func (list *DescribeRecommendInstanceTypeRecommendInstanceTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRecommendInstanceTypeRecommendInstanceType)
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

type DescribeRecommendInstanceTypeZoneList []DescribeRecommendInstanceTypeZone

func (list *DescribeRecommendInstanceTypeZoneList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRecommendInstanceTypeZone)
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

type DescribeRecommendInstanceTypeNetworkTypeList []common.String

func (list *DescribeRecommendInstanceTypeNetworkTypeList) UnmarshalJSON(data []byte) error {
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
