package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeDomainBpsDataRequest struct {
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

func (req *DescribeDomainBpsDataRequest) Invoke(client *sdk.Client) (resp *DescribeDomainBpsDataResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainBpsData", "", "")
	resp = &DescribeDomainBpsDataResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDomainBpsDataResponse struct {
	responses.BaseResponse
	RequestId          common.String
	DomainName         common.String
	DataInterval       common.String
	StartTime          common.String
	EndTime            common.String
	LocationNameEn     common.String
	IspNameEn          common.String
	LocationName       common.String
	IspName            common.String
	BpsDataPerInterval DescribeDomainBpsDataDataModuleList
	SupplyBpsDatas     DescribeDomainBpsDataDataModule1List
}

type DescribeDomainBpsDataDataModule struct {
	TimeStamp            common.String
	Value                common.String
	DomesticValue        common.String
	OverseasValue        common.String
	DynamicValue         common.String
	DynamicDomesticValue common.String
	DynamicOverseasValue common.String
	StaticValue          common.String
	StaticDomesticValue  common.String
	StaticOverseasValue  common.String
	L2Value              common.String
	DomesticL2Value      common.String
	OverseasL2Value      common.String
}

type DescribeDomainBpsDataDataModule1 struct {
	TimeStamp common.String
	Value     common.String
}

type DescribeDomainBpsDataDataModuleList []DescribeDomainBpsDataDataModule

func (list *DescribeDomainBpsDataDataModuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainBpsDataDataModule)
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

type DescribeDomainBpsDataDataModule1List []DescribeDomainBpsDataDataModule1

func (list *DescribeDomainBpsDataDataModule1List) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainBpsDataDataModule1)
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
