package dds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeAuditRecordsRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	EndTime              string `position:"Query" name:"EndTime"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	QueryKeywords        string `position:"Query" name:"QueryKeywords"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
	Database             string `position:"Query" name:"Database"`
	Form                 string `position:"Query" name:"Form"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	PageSize             int    `position:"Query" name:"PageSize"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	NodeId               string `position:"Query" name:"NodeId"`
	User                 string `position:"Query" name:"User"`
}

func (r DescribeAuditRecordsRequest) Invoke(client *sdk.Client) (response *DescribeAuditRecordsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeAuditRecordsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Dds", "2015-12-01", "DescribeAuditRecords", "dds", "")

	resp := struct {
		*responses.BaseResponse
		DescribeAuditRecordsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeAuditRecordsResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeAuditRecordsResponse struct {
	RequestId        string
	TotalRecordCount int
	PageNumber       int
	PageRecordCount  int
	Items            DescribeAuditRecordsSQLRecordList
}

type DescribeAuditRecordsSQLRecord struct {
	DBName              string
	AccountName         string
	HostAddress         string
	Syntax              string
	TotalExecutionTimes int64
	ReturnRowCounts     int64
	ExecuteTime         string
	ThreadID            string
}

type DescribeAuditRecordsSQLRecordList []DescribeAuditRecordsSQLRecord

func (list *DescribeAuditRecordsSQLRecordList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAuditRecordsSQLRecord)
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
