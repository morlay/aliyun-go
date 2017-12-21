package vod

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDomainFlowDataRequest struct {
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	TimeMerge            string `position:"Query" name:"TimeMerge"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DomainName           string `position:"Query" name:"DomainName"`
	EndTime              string `position:"Query" name:"EndTime"`
	LocationNameEn       string `position:"Query" name:"LocationNameEn"`
	StartTime            string `position:"Query" name:"StartTime"`
	IspNameEn            string `position:"Query" name:"IspNameEn"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	Interval             string `position:"Query" name:"Interval"`
}

func (r DescribeDomainFlowDataRequest) Invoke(client *sdk.Client) (response *DescribeDomainFlowDataResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeDomainFlowDataRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("vod", "2017-03-21", "DescribeDomainFlowData", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeDomainFlowDataResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeDomainFlowDataResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeDomainFlowDataResponse struct {
	RequestId           string
	DomainName          string
	DataInterval        string
	StartTime           string
	EndTime             string
	FlowDataPerInterval DescribeDomainFlowDataDataModuleList
}

type DescribeDomainFlowDataDataModule struct {
	TimeStamp            string
	Value                string
	DomesticValue        string
	OverseasValue        string
	DynamicValue         string
	DynamicDomesticValue string
	DynamicOverseasValue string
	StaticValue          string
	StaticDomesticValue  string
	StaticOverseasValue  string
}

type DescribeDomainFlowDataDataModuleList []DescribeDomainFlowDataDataModule

func (list *DescribeDomainFlowDataDataModuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainFlowDataDataModule)
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
