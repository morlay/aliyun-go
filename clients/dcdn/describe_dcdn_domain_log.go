package dcdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDcdnDomainLogRequest struct {
	requests.RpcRequest
	StartTime  string `position:"Query" name:"StartTime"`
	PageNumber int64  `position:"Query" name:"PageNumber"`
	PageSize   int64  `position:"Query" name:"PageSize"`
	DomainName string `position:"Query" name:"DomainName"`
	EndTime    string `position:"Query" name:"EndTime"`
	OwnerId    int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeDcdnDomainLogRequest) Invoke(client *sdk.Client) (resp *DescribeDcdnDomainLogResponse, err error) {
	req.InitWithApiInfo("dcdn", "2018-01-15", "DescribeDcdnDomainLog", "dcdn", "")
	resp = &DescribeDcdnDomainLogResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDcdnDomainLogResponse struct {
	responses.BaseResponse
	RequestId        string
	DomainName       string
	DomainLogDetails DescribeDcdnDomainLogDomainLogDetailList
}

type DescribeDcdnDomainLogDomainLogDetail struct {
	LogCount  int64
	PageInfos DescribeDcdnDomainLogPageInfoDetailList
	LogInfos  DescribeDcdnDomainLogLogInfoDetailList
}

type DescribeDcdnDomainLogPageInfoDetail struct {
	PageIndex int64
	PageSize  int64
	Total     int64
}

type DescribeDcdnDomainLogLogInfoDetail struct {
	LogName   string
	LogPath   string
	LogSize   int64
	StartTime string
	EndTime   string
}

type DescribeDcdnDomainLogDomainLogDetailList []DescribeDcdnDomainLogDomainLogDetail

func (list *DescribeDcdnDomainLogDomainLogDetailList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDcdnDomainLogDomainLogDetail)
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

type DescribeDcdnDomainLogPageInfoDetailList []DescribeDcdnDomainLogPageInfoDetail

func (list *DescribeDcdnDomainLogPageInfoDetailList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDcdnDomainLogPageInfoDetail)
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

type DescribeDcdnDomainLogLogInfoDetailList []DescribeDcdnDomainLogLogInfoDetail

func (list *DescribeDcdnDomainLogLogInfoDetailList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDcdnDomainLogLogInfoDetail)
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
