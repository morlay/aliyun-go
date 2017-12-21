package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDomainPvDataRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeDomainPvDataRequest) Invoke(client *sdk.Client) (response *DescribeDomainPvDataResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeDomainPvDataRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainPvData", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeDomainPvDataResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeDomainPvDataResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeDomainPvDataResponse struct {
	RequestId      string
	DomainName     string
	DataInterval   string
	StartTime      string
	EndTime        string
	PvDataInterval DescribeDomainPvDataUsageDataList
}

type DescribeDomainPvDataUsageData struct {
	TimeStamp string
	Value     string
}

type DescribeDomainPvDataUsageDataList []DescribeDomainPvDataUsageData

func (list *DescribeDomainPvDataUsageDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainPvDataUsageData)
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
