package dcdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeDcdnDomainBpsDataRequest struct {
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

func (req *DescribeDcdnDomainBpsDataRequest) Invoke(client *sdk.Client) (resp *DescribeDcdnDomainBpsDataResponse, err error) {
	req.InitWithApiInfo("dcdn", "2018-01-15", "DescribeDcdnDomainBpsData", "dcdn", "")
	resp = &DescribeDcdnDomainBpsDataResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDcdnDomainBpsDataResponse struct {
	responses.BaseResponse
	RequestId          common.String
	DomainName         common.String
	StartTime          common.String
	EndTime            common.String
	DataInterval       common.String
	BpsDataPerInterval DescribeDcdnDomainBpsDataDataModuleList
}

type DescribeDcdnDomainBpsDataDataModule struct {
	TimeStamp       common.String
	Bps             common.Float
	DynamicHttpBps  common.Float
	DynamicHttpsBps common.Float
	StaticHttpBps   common.Float
	StaticHttpsBps  common.Float
}

type DescribeDcdnDomainBpsDataDataModuleList []DescribeDcdnDomainBpsDataDataModule

func (list *DescribeDcdnDomainBpsDataDataModuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDcdnDomainBpsDataDataModule)
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
