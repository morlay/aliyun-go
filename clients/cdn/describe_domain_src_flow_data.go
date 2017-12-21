package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDomainSrcFlowDataRequest struct {
	FixTimeGap    string `position:"Query" name:"FixTimeGap"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	TimeMerge     string `position:"Query" name:"TimeMerge"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	Interval      string `position:"Query" name:"Interval"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeDomainSrcFlowDataRequest) Invoke(client *sdk.Client) (response *DescribeDomainSrcFlowDataResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeDomainSrcFlowDataRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainSrcFlowData", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeDomainSrcFlowDataResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeDomainSrcFlowDataResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeDomainSrcFlowDataResponse struct {
	RequestId              string
	DomainName             string
	DataInterval           string
	StartTime              string
	EndTime                string
	SrcFlowDataPerInterval DescribeDomainSrcFlowDataDataModuleList
}

type DescribeDomainSrcFlowDataDataModule struct {
	TimeStamp string
	Value     string
}

type DescribeDomainSrcFlowDataDataModuleList []DescribeDomainSrcFlowDataDataModule

func (list *DescribeDomainSrcFlowDataDataModuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainSrcFlowDataDataModule)
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
