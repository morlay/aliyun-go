package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDomainQpsDataRequest struct {
	requests.RpcRequest
	FixTimeGap     string `position:"Query" name:"FixTimeGap"`
	TimeMerge      string `position:"Query" name:"TimeMerge"`
	DomainName     string `position:"Query" name:"DomainName"`
	EndTime        string `position:"Query" name:"EndTime"`
	LocationNameEn string `position:"Query" name:"LocationNameEn"`
	StartTime      string `position:"Query" name:"StartTime"`
	IspNameEn      string `position:"Query" name:"IspNameEn"`
	OwnerId        int64  `position:"Query" name:"OwnerId"`
	DomainType     string `position:"Query" name:"DomainType"`
	SecurityToken  string `position:"Query" name:"SecurityToken"`
	Interval       string `position:"Query" name:"Interval"`
}

func (req *DescribeDomainQpsDataRequest) Invoke(client *sdk.Client) (resp *DescribeDomainQpsDataResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainQpsData", "", "")
	resp = &DescribeDomainQpsDataResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDomainQpsDataResponse struct {
	responses.BaseResponse
	RequestId       string
	DomainName      string
	DataInterval    string
	StartTime       string
	EndTime         string
	QpsDataInterval DescribeDomainQpsDataDataModuleList
}

type DescribeDomainQpsDataDataModule struct {
	TimeStamp            string
	Value                string
	DomesticValue        string
	OverseasValue        string
	AccValue             string
	AccDomesticValue     string
	AccOverseasValue     string
	DynamicValue         string
	DynamicDomesticValue string
	DynamicOverseasValue string
	StaticValue          string
	StaticDomesticValue  string
	StaticOverseasValue  string
}

type DescribeDomainQpsDataDataModuleList []DescribeDomainQpsDataDataModule

func (list *DescribeDomainQpsDataDataModuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainQpsDataDataModule)
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
