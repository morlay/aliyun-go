package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeDomainHttpCodeDataRequest struct {
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

func (req *DescribeDomainHttpCodeDataRequest) Invoke(client *sdk.Client) (resp *DescribeDomainHttpCodeDataResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainHttpCodeData", "", "")
	resp = &DescribeDomainHttpCodeDataResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDomainHttpCodeDataResponse struct {
	responses.BaseResponse
	RequestId    common.String
	DomainName   common.String
	DataInterval common.String
	StartTime    common.String
	EndTime      common.String
	HttpCodeData DescribeDomainHttpCodeDataUsageDataList
}

type DescribeDomainHttpCodeDataUsageData struct {
	TimeStamp common.String
	Value     DescribeDomainHttpCodeDataCodeProportionDataList
}

type DescribeDomainHttpCodeDataCodeProportionData struct {
	Code       common.String
	Proportion common.String
	Count      common.String
}

type DescribeDomainHttpCodeDataUsageDataList []DescribeDomainHttpCodeDataUsageData

func (list *DescribeDomainHttpCodeDataUsageDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainHttpCodeDataUsageData)
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

type DescribeDomainHttpCodeDataCodeProportionDataList []DescribeDomainHttpCodeDataCodeProportionData

func (list *DescribeDomainHttpCodeDataCodeProportionDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainHttpCodeDataCodeProportionData)
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
