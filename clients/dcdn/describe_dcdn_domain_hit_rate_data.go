package dcdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDcdnDomainHitRateDataRequest struct {
	requests.RpcRequest
	StartTime  string `position:"Query" name:"StartTime"`
	DomainName string `position:"Query" name:"DomainName"`
	EndTime    string `position:"Query" name:"EndTime"`
	OwnerId    int64  `position:"Query" name:"OwnerId"`
	Interval   string `position:"Query" name:"Interval"`
}

func (req *DescribeDcdnDomainHitRateDataRequest) Invoke(client *sdk.Client) (resp *DescribeDcdnDomainHitRateDataResponse, err error) {
	req.InitWithApiInfo("dcdn", "2018-01-15", "DescribeDcdnDomainHitRateData", "dcdn", "")
	resp = &DescribeDcdnDomainHitRateDataResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDcdnDomainHitRateDataResponse struct {
	responses.BaseResponse
	RequestId          string
	DomainName         string
	StartTime          string
	EndTime            string
	DataInterval       string
	HitRatePerInterval DescribeDcdnDomainHitRateDataDataModuleList
}

type DescribeDcdnDomainHitRateDataDataModule struct {
	TimeStamp   string
	ReqHitRate  float32
	ByteHitRate float32
}

type DescribeDcdnDomainHitRateDataDataModuleList []DescribeDcdnDomainHitRateDataDataModule

func (list *DescribeDcdnDomainHitRateDataDataModuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDcdnDomainHitRateDataDataModule)
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
