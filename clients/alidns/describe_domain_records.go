package alidns

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeDomainRecordsRequest struct {
	requests.RpcRequest
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	DomainName   string `position:"Query" name:"DomainName"`
	PageNumber   int64  `position:"Query" name:"PageNumber"`
	PageSize     int64  `position:"Query" name:"PageSize"`
	RRKeyWord    string `position:"Query" name:"RRKeyWord"`
	TypeKeyWord  string `position:"Query" name:"TypeKeyWord"`
	ValueKeyWord string `position:"Query" name:"ValueKeyWord"`
}

func (req *DescribeDomainRecordsRequest) Invoke(client *sdk.Client) (resp *DescribeDomainRecordsResponse, err error) {
	req.InitWithApiInfo("Alidns", "2015-01-09", "DescribeDomainRecords", "", "")
	resp = &DescribeDomainRecordsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDomainRecordsResponse struct {
	responses.BaseResponse
	RequestId     common.String
	TotalCount    common.Long
	PageNumber    common.Long
	PageSize      common.Long
	DomainRecords DescribeDomainRecordsRecordList
}

type DescribeDomainRecordsRecord struct {
	DomainName common.String
	RecordId   common.String
	RR         common.String
	Type       common.String
	Value      common.String
	TTL        common.Long
	Priority   common.Long
	Line       common.String
	Status     common.String
	Locked     bool
	Weight     common.Integer
}

type DescribeDomainRecordsRecordList []DescribeDomainRecordsRecord

func (list *DescribeDomainRecordsRecordList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainRecordsRecord)
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
