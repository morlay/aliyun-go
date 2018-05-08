package vod

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeDomainFlowDataRequest struct {
	requests.RpcRequest
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

func (req *DescribeDomainFlowDataRequest) Invoke(client *sdk.Client) (resp *DescribeDomainFlowDataResponse, err error) {
	req.InitWithApiInfo("vod", "2017-03-21", "DescribeDomainFlowData", "vod", "")
	resp = &DescribeDomainFlowDataResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDomainFlowDataResponse struct {
	responses.BaseResponse
	RequestId           common.String
	DomainName          common.String
	DataInterval        common.String
	StartTime           common.String
	EndTime             common.String
	FlowDataPerInterval DescribeDomainFlowDataDataModuleList
}

type DescribeDomainFlowDataDataModule struct {
	TimeStamp            common.String
	Value                common.String
	DomesticValue        common.String
	OverseasValue        common.String
	DynamicValue         common.String
	DynamicDomesticValue common.String
	DynamicOverseasValue common.String
	StaticValue          common.String
	StaticDomesticValue  common.String
	StaticOverseasValue  common.String
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
