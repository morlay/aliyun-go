package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDomainPvDataRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeDomainPvDataRequest) Invoke(client *sdk.Client) (resp *DescribeDomainPvDataResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainPvData", "", "")
	resp = &DescribeDomainPvDataResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDomainPvDataResponse struct {
	responses.BaseResponse
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
