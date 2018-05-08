package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeDomainHitRateDataRequest struct {
	requests.RpcRequest
	SecurityToken  string `position:"Query" name:"SecurityToken"`
	TimeMerge      string `position:"Query" name:"TimeMerge"`
	DomainName     string `position:"Query" name:"DomainName"`
	EndTime        string `position:"Query" name:"EndTime"`
	LocationNameEn string `position:"Query" name:"LocationNameEn"`
	Interval       string `position:"Query" name:"Interval"`
	StartTime      string `position:"Query" name:"StartTime"`
	IspNameEn      string `position:"Query" name:"IspNameEn"`
	OwnerId        int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeDomainHitRateDataRequest) Invoke(client *sdk.Client) (resp *DescribeDomainHitRateDataResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainHitRateData", "", "")
	resp = &DescribeDomainHitRateDataResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDomainHitRateDataResponse struct {
	responses.BaseResponse
	RequestId       common.String
	DomainName      common.String
	DataInterval    common.String
	StartTime       common.String
	EndTime         common.String
	HitRateInterval DescribeDomainHitRateDataDataModuleList
}

type DescribeDomainHitRateDataDataModule struct {
	TimeStamp common.String
	Value     common.String
}

type DescribeDomainHitRateDataDataModuleList []DescribeDomainHitRateDataDataModule

func (list *DescribeDomainHitRateDataDataModuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainHitRateDataDataModule)
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
