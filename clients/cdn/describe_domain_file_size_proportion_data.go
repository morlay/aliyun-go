package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDomainFileSizeProportionDataRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeDomainFileSizeProportionDataRequest) Invoke(client *sdk.Client) (response *DescribeDomainFileSizeProportionDataResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeDomainFileSizeProportionDataRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainFileSizeProportionData", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeDomainFileSizeProportionDataResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeDomainFileSizeProportionDataResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeDomainFileSizeProportionDataResponse struct {
	RequestId                      string
	DomainName                     string
	DataInterval                   string
	StartTime                      string
	EndTime                        string
	FileSizeProportionDataInterval DescribeDomainFileSizeProportionDataUsageDataList
}

type DescribeDomainFileSizeProportionDataUsageData struct {
	TimeStamp string
	Value     DescribeDomainFileSizeProportionDataFileSizeProportionDataList
}

type DescribeDomainFileSizeProportionDataFileSizeProportionData struct {
	FileSize   string
	Proportion string
}

type DescribeDomainFileSizeProportionDataUsageDataList []DescribeDomainFileSizeProportionDataUsageData

func (list *DescribeDomainFileSizeProportionDataUsageDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainFileSizeProportionDataUsageData)
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

type DescribeDomainFileSizeProportionDataFileSizeProportionDataList []DescribeDomainFileSizeProportionDataFileSizeProportionData

func (list *DescribeDomainFileSizeProportionDataFileSizeProportionDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainFileSizeProportionDataFileSizeProportionData)
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
