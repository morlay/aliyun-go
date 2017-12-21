package alidns

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeSubDomainRecordsRequest struct {
	requests.RpcRequest
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	SubDomain    string `position:"Query" name:"SubDomain"`
	PageNumber   int64  `position:"Query" name:"PageNumber"`
	PageSize     int64  `position:"Query" name:"PageSize"`
	Type         string `position:"Query" name:"Type"`
}

func (req *DescribeSubDomainRecordsRequest) Invoke(client *sdk.Client) (resp *DescribeSubDomainRecordsResponse, err error) {
	req.InitWithApiInfo("Alidns", "2015-01-09", "DescribeSubDomainRecords", "", "")
	resp = &DescribeSubDomainRecordsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeSubDomainRecordsResponse struct {
	responses.BaseResponse
	RequestId     string
	TotalCount    int64
	PageNumber    int64
	PageSize      int64
	DomainRecords DescribeSubDomainRecordsRecordList
}

type DescribeSubDomainRecordsRecord struct {
	DomainName string
	RecordId   string
	RR         string
	Type       string
	Value      string
	TTL        int64
	Priority   int64
	Line       string
	Status     string
	Locked     bool
	Weight     int
}

type DescribeSubDomainRecordsRecordList []DescribeSubDomainRecordsRecord

func (list *DescribeSubDomainRecordsRecordList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSubDomainRecordsRecord)
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
