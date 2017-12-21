package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeRecommendInstanceTypeRequest struct {
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

func (r DescribeRecommendInstanceTypeRequest) Invoke(client *sdk.Client) (response *DescribeRecommendInstanceTypeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeRecommendInstanceTypeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeRecommendInstanceType", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		DescribeRecommendInstanceTypeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeRecommendInstanceTypeResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeRecommendInstanceTypeResponse struct {
	RequestId string
	Data      DescribeRecommendInstanceTypeRecommendInstanceTypeList
}

type DescribeRecommendInstanceTypeRecommendInstanceType struct {
	RegionNo      string
	CommodityCode string
	Scene         string
	Zones         DescribeRecommendInstanceTypeZoneList
	InstanceType  DescribeRecommendInstanceTypeInstanceType
}

type DescribeRecommendInstanceTypeZone struct {
	ZoneNo       string
	NetworkTypes DescribeRecommendInstanceTypeNetworkTypeList
}

type DescribeRecommendInstanceTypeInstanceType struct {
	Generation         string
	InstanceTypeFamily string
	InstanceType       string
	SupportIoOptimized string
	Cores              int
	Memory             int
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

type DescribeRecommendInstanceTypeNetworkTypeList []string

func (list *DescribeRecommendInstanceTypeNetworkTypeList) UnmarshalJSON(data []byte) error {
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
