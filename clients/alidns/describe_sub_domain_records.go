package alidns

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeSubDomainRecordsRequest struct {
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	SubDomain    string `position:"Query" name:"SubDomain"`
	PageNumber   int64  `position:"Query" name:"PageNumber"`
	PageSize     int64  `position:"Query" name:"PageSize"`
	Type         string `position:"Query" name:"Type"`
}

func (r DescribeSubDomainRecordsRequest) Invoke(client *sdk.Client) (response *DescribeSubDomainRecordsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeSubDomainRecordsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Alidns", "2015-01-09", "DescribeSubDomainRecords", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeSubDomainRecordsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeSubDomainRecordsResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeSubDomainRecordsResponse struct {
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
