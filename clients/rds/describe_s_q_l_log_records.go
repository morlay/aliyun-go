package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeSQLLogRecordsRequest struct {
	requests.RpcRequest
	SQLId                int64  `position:"Query" name:"SQLId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	EndTime              string `position:"Query" name:"EndTime"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	QueryKeywords        string `position:"Query" name:"QueryKeywords"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
	Database             string `position:"Query" name:"Database"`
	Form                 string `position:"Query" name:"Form"`
	PageSize             int    `position:"Query" name:"PageSize"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	User                 string `position:"Query" name:"User"`
}

func (req *DescribeSQLLogRecordsRequest) Invoke(client *sdk.Client) (resp *DescribeSQLLogRecordsResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeSQLLogRecords", "rds", "")
	resp = &DescribeSQLLogRecordsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeSQLLogRecordsResponse struct {
	responses.BaseResponse
	RequestId        string
	TotalRecordCount int64
	PageNumber       int
	PageRecordCount  int
	Items            DescribeSQLLogRecordsSQLRecordList
}

type DescribeSQLLogRecordsSQLRecord struct {
	DBName              string
	AccountName         string
	HostAddress         string
	SQLText             string
	TotalExecutionTimes int64
	ReturnRowCounts     int64
	ExecuteTime         string
	ThreadID            string
}

type DescribeSQLLogRecordsSQLRecordList []DescribeSQLLogRecordsSQLRecord

func (list *DescribeSQLLogRecordsSQLRecordList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSQLLogRecordsSQLRecord)
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
