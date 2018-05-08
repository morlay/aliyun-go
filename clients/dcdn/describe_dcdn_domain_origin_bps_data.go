package dcdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeDcdnDomainOriginBpsDataRequest struct {
	requests.RpcRequest
	StartTime  string `position:"Query" name:"StartTime"`
	FixTimeGap string `position:"Query" name:"FixTimeGap"`
	TimeMerge  string `position:"Query" name:"TimeMerge"`
	DomainName string `position:"Query" name:"DomainName"`
	EndTime    string `position:"Query" name:"EndTime"`
	OwnerId    int64  `position:"Query" name:"OwnerId"`
	Interval   string `position:"Query" name:"Interval"`
}

func (req *DescribeDcdnDomainOriginBpsDataRequest) Invoke(client *sdk.Client) (resp *DescribeDcdnDomainOriginBpsDataResponse, err error) {
	req.InitWithApiInfo("dcdn", "2018-01-15", "DescribeDcdnDomainOriginBpsData", "dcdn", "")
	resp = &DescribeDcdnDomainOriginBpsDataResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDcdnDomainOriginBpsDataResponse struct {
	responses.BaseResponse
	RequestId                common.String
	DomainName               common.String
	StartTime                common.String
	EndTime                  common.String
	DataInterval             common.String
	OriginBpsDataPerInterval DescribeDcdnDomainOriginBpsDataDataModuleList
}

type DescribeDcdnDomainOriginBpsDataDataModule struct {
	TimeStamp             common.String
	OriginBps             common.Float
	DynamicHttpOriginBps  common.Float
	DynamicHttpsOriginBps common.Float
	StaticHttpOriginBps   common.Float
	StaticHttpsOriginBps  common.Float
}

type DescribeDcdnDomainOriginBpsDataDataModuleList []DescribeDcdnDomainOriginBpsDataDataModule

func (list *DescribeDcdnDomainOriginBpsDataDataModuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDcdnDomainOriginBpsDataDataModule)
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
