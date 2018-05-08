package dcdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeDcdnDomainHttpCodeDataRequest struct {
	requests.RpcRequest
	LocationNameEn string `position:"Query" name:"LocationNameEn"`
	StartTime      string `position:"Query" name:"StartTime"`
	IspNameEn      string `position:"Query" name:"IspNameEn"`
	DomainName     string `position:"Query" name:"DomainName"`
	EndTime        string `position:"Query" name:"EndTime"`
	OwnerId        int64  `position:"Query" name:"OwnerId"`
	Interval       string `position:"Query" name:"Interval"`
}

func (req *DescribeDcdnDomainHttpCodeDataRequest) Invoke(client *sdk.Client) (resp *DescribeDcdnDomainHttpCodeDataResponse, err error) {
	req.InitWithApiInfo("dcdn", "2018-01-15", "DescribeDcdnDomainHttpCodeData", "dcdn", "")
	resp = &DescribeDcdnDomainHttpCodeDataResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDcdnDomainHttpCodeDataResponse struct {
	responses.BaseResponse
	RequestId       common.String
	DomainName      common.String
	StartTime       common.String
	EndTime         common.String
	DataInterval    common.String
	DataPerInterval DescribeDcdnDomainHttpCodeDataDataModuleList
}

type DescribeDcdnDomainHttpCodeDataDataModule struct {
	TimeStamp               common.String
	HttpCodeDataPerInterval DescribeDcdnDomainHttpCodeDataHttpCodeDataModuleList
}

type DescribeDcdnDomainHttpCodeDataHttpCodeDataModule struct {
	Code       common.Integer
	Proportion common.Float
	Count      common.Float
}

type DescribeDcdnDomainHttpCodeDataDataModuleList []DescribeDcdnDomainHttpCodeDataDataModule

func (list *DescribeDcdnDomainHttpCodeDataDataModuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDcdnDomainHttpCodeDataDataModule)
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

type DescribeDcdnDomainHttpCodeDataHttpCodeDataModuleList []DescribeDcdnDomainHttpCodeDataHttpCodeDataModule

func (list *DescribeDcdnDomainHttpCodeDataHttpCodeDataModuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDcdnDomainHttpCodeDataHttpCodeDataModule)
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
