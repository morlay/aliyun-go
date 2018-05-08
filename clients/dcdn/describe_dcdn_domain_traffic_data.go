package dcdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeDcdnDomainTrafficDataRequest struct {
	requests.RpcRequest
	LocationNameEn string `position:"Query" name:"LocationNameEn"`
	StartTime      string `position:"Query" name:"StartTime"`
	IspNameEn      string `position:"Query" name:"IspNameEn"`
	FixTimeGap     string `position:"Query" name:"FixTimeGap"`
	TimeMerge      string `position:"Query" name:"TimeMerge"`
	DomainName     string `position:"Query" name:"DomainName"`
	EndTime        string `position:"Query" name:"EndTime"`
	OwnerId        int64  `position:"Query" name:"OwnerId"`
	Interval       string `position:"Query" name:"Interval"`
}

func (req *DescribeDcdnDomainTrafficDataRequest) Invoke(client *sdk.Client) (resp *DescribeDcdnDomainTrafficDataResponse, err error) {
	req.InitWithApiInfo("dcdn", "2018-01-15", "DescribeDcdnDomainTrafficData", "dcdn", "")
	resp = &DescribeDcdnDomainTrafficDataResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDcdnDomainTrafficDataResponse struct {
	responses.BaseResponse
	RequestId              common.String
	DomainName             common.String
	StartTime              common.String
	EndTime                common.String
	DataInterval           common.String
	TrafficDataPerInterval DescribeDcdnDomainTrafficDataDataModuleList
}

type DescribeDcdnDomainTrafficDataDataModule struct {
	TimeStamp           common.String
	Traffic             common.Float
	DynamicHttpTraffic  common.Float
	DynamicHttpsTraffic common.Float
	StaticHttpTraffic   common.Float
	StaticHttpsTraffic  common.Float
}

type DescribeDcdnDomainTrafficDataDataModuleList []DescribeDcdnDomainTrafficDataDataModule

func (list *DescribeDcdnDomainTrafficDataDataModuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDcdnDomainTrafficDataDataModule)
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
