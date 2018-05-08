package dcdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeDcdnDomainQpsDataRequest struct {
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

func (req *DescribeDcdnDomainQpsDataRequest) Invoke(client *sdk.Client) (resp *DescribeDcdnDomainQpsDataResponse, err error) {
	req.InitWithApiInfo("dcdn", "2018-01-15", "DescribeDcdnDomainQpsData", "dcdn", "")
	resp = &DescribeDcdnDomainQpsDataResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDcdnDomainQpsDataResponse struct {
	responses.BaseResponse
	RequestId          common.String
	DomainName         common.String
	StartTime          common.String
	EndTime            common.String
	DataInterval       common.String
	QpsDataPerInterval DescribeDcdnDomainQpsDataDataModuleList
}

type DescribeDcdnDomainQpsDataDataModule struct {
	TimeStamp       common.String
	Qps             common.Float
	DynamicHttpQps  common.Float
	DynamicHttpsQps common.Float
	StaticHttpQps   common.Float
	StaticHttpsQps  common.Float
	Acc             common.Float
	DynamicHttpAcc  common.Float
	DynamicHttpsAcc common.Float
	StaticHttpAcc   common.Float
	StaticHttpsAcc  common.Float
}

type DescribeDcdnDomainQpsDataDataModuleList []DescribeDcdnDomainQpsDataDataModule

func (list *DescribeDcdnDomainQpsDataDataModuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDcdnDomainQpsDataDataModule)
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
