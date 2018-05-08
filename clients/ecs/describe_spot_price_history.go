package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeSpotPriceHistoryRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	IoOptimized          string `position:"Query" name:"IoOptimized"`
	NetworkType          string `position:"Query" name:"NetworkType"`
	StartTime            string `position:"Query" name:"StartTime"`
	InstanceType         string `position:"Query" name:"InstanceType"`
	Offset               int    `position:"Query" name:"Offset"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	EndTime              string `position:"Query" name:"EndTime"`
	OSType               string `position:"Query" name:"OSType"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ZoneId               string `position:"Query" name:"ZoneId"`
}

func (req *DescribeSpotPriceHistoryRequest) Invoke(client *sdk.Client) (resp *DescribeSpotPriceHistoryResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeSpotPriceHistory", "ecs", "")
	resp = &DescribeSpotPriceHistoryResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeSpotPriceHistoryResponse struct {
	responses.BaseResponse
	RequestId  common.String
	NextOffset common.Integer
	Currency   common.String
	SpotPrices DescribeSpotPriceHistorySpotPriceTypeList
}

type DescribeSpotPriceHistorySpotPriceType struct {
	ZoneId       common.String
	InstanceType common.String
	IoOptimized  common.String
	Timestamp    common.String
	NetworkType  common.String
	SpotPrice    common.Float
	OriginPrice  common.Float
}

type DescribeSpotPriceHistorySpotPriceTypeList []DescribeSpotPriceHistorySpotPriceType

func (list *DescribeSpotPriceHistorySpotPriceTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSpotPriceHistorySpotPriceType)
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
