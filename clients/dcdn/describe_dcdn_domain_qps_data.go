package dcdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId          string
	DomainName         string
	StartTime          string
	EndTime            string
	DataInterval       string
	QpsDataPerInterval DescribeDcdnDomainQpsDataDataModuleList
}

type DescribeDcdnDomainQpsDataDataModule struct {
	TimeStamp       string
	Qps             float32
	DynamicHttpQps  float32
	DynamicHttpsQps float32
	StaticHttpQps   float32
	StaticHttpsQps  float32
	Acc             float32
	DynamicHttpAcc  float32
	DynamicHttpsAcc float32
	StaticHttpAcc   float32
	StaticHttpsAcc  float32
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
