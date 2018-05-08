package dcdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDcdnDomainOriginTrafficDataRequest struct {
	requests.RpcRequest
	StartTime  string `position:"Query" name:"StartTime"`
	FixTimeGap string `position:"Query" name:"FixTimeGap"`
	TimeMerge  string `position:"Query" name:"TimeMerge"`
	DomainName string `position:"Query" name:"DomainName"`
	EndTime    string `position:"Query" name:"EndTime"`
	OwnerId    int64  `position:"Query" name:"OwnerId"`
	Interval   string `position:"Query" name:"Interval"`
}

func (req *DescribeDcdnDomainOriginTrafficDataRequest) Invoke(client *sdk.Client) (resp *DescribeDcdnDomainOriginTrafficDataResponse, err error) {
	req.InitWithApiInfo("dcdn", "2018-01-15", "DescribeDcdnDomainOriginTrafficData", "dcdn", "")
	resp = &DescribeDcdnDomainOriginTrafficDataResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDcdnDomainOriginTrafficDataResponse struct {
	responses.BaseResponse
	RequestId                    string
	DomainName                   string
	StartTime                    string
	EndTime                      string
	DataInterval                 string
	OriginTrafficDataPerInterval DescribeDcdnDomainOriginTrafficDataDataModuleList
}

type DescribeDcdnDomainOriginTrafficDataDataModule struct {
	TimeStamp                 string
	OriginTraffic             float32
	DynamicHttpOriginTraffic  float32
	DynamicHttpsOriginTraffic float32
	StaticHttpOriginTraffic   float32
	StaticHttpsOriginTraffic  float32
}

type DescribeDcdnDomainOriginTrafficDataDataModuleList []DescribeDcdnDomainOriginTrafficDataDataModule

func (list *DescribeDcdnDomainOriginTrafficDataDataModuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDcdnDomainOriginTrafficDataDataModule)
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
