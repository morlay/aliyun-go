package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeSpotPriceHistoryRequest struct {
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

func (r DescribeSpotPriceHistoryRequest) Invoke(client *sdk.Client) (response *DescribeSpotPriceHistoryResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeSpotPriceHistoryRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeSpotPriceHistory", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		DescribeSpotPriceHistoryResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeSpotPriceHistoryResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeSpotPriceHistoryResponse struct {
	RequestId  string
	NextOffset int
	Currency   string
	SpotPrices DescribeSpotPriceHistorySpotPriceTypeList
}

type DescribeSpotPriceHistorySpotPriceType struct {
	ZoneId       string
	InstanceType string
	IoOptimized  string
	Timestamp    string
	NetworkType  string
	SpotPrice    float32
	OriginPrice  float32
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
