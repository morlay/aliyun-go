package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDomainRealTimeDataRequest struct {
	Field         string `position:"Query" name:"Field"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeDomainRealTimeDataRequest) Invoke(client *sdk.Client) (response *DescribeDomainRealTimeDataResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeDomainRealTimeDataRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainRealTimeData", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeDomainRealTimeDataResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeDomainRealTimeDataResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeDomainRealTimeDataResponse struct {
	RequestId       string
	DomainName      string
	Field           string
	StartTime       string
	EndTime         string
	DataPerInterval DescribeDomainRealTimeDataDataModuleList
}

type DescribeDomainRealTimeDataDataModule struct {
	TimeStamp string
	Value     string
}

type DescribeDomainRealTimeDataDataModuleList []DescribeDomainRealTimeDataDataModule

func (list *DescribeDomainRealTimeDataDataModuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainRealTimeDataDataModule)
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
