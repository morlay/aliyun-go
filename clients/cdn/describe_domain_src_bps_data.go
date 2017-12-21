package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDomainSrcBpsDataRequest struct {
	FixTimeGap    string `position:"Query" name:"FixTimeGap"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	TimeMerge     string `position:"Query" name:"TimeMerge"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	Interval      string `position:"Query" name:"Interval"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeDomainSrcBpsDataRequest) Invoke(client *sdk.Client) (response *DescribeDomainSrcBpsDataResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeDomainSrcBpsDataRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainSrcBpsData", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeDomainSrcBpsDataResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeDomainSrcBpsDataResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeDomainSrcBpsDataResponse struct {
	RequestId             string
	DomainName            string
	DataInterval          string
	StartTime             string
	EndTime               string
	SrcBpsDataPerInterval DescribeDomainSrcBpsDataDataModuleList
}

type DescribeDomainSrcBpsDataDataModule struct {
	TimeStamp string
	Value     string
}

type DescribeDomainSrcBpsDataDataModuleList []DescribeDomainSrcBpsDataDataModule

func (list *DescribeDomainSrcBpsDataDataModuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainSrcBpsDataDataModule)
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
