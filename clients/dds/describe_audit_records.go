package dds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeAuditRecordsRequest struct {
	requests.RpcRequest
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

func (req *DescribeAuditRecordsRequest) Invoke(client *sdk.Client) (resp *DescribeAuditRecordsResponse, err error) {
	req.InitWithApiInfo("Dds", "2015-12-01", "DescribeAuditRecords", "dds", "")
	resp = &DescribeAuditRecordsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeAuditRecordsResponse struct {
	responses.BaseResponse
	RequestId        common.String
	TotalRecordCount common.Integer
	PageNumber       common.Integer
	PageRecordCount  common.Integer
	Items            DescribeAuditRecordsSQLRecordList
}

type DescribeAuditRecordsSQLRecord struct {
	DBName              common.String
	AccountName         common.String
	HostAddress         common.String
	Syntax              common.String
	TotalExecutionTimes common.Long
	ReturnRowCounts     common.Long
	ExecuteTime         common.String
	ThreadID            common.String
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
