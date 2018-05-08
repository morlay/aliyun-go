package alidns

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId     common.String
	TotalCount    common.Long
	PageNumber    common.Long
	PageSize      common.Long
	DomainRecords DescribeSubDomainRecordsRecordList
}

type DescribeSubDomainRecordsRecord struct {
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
