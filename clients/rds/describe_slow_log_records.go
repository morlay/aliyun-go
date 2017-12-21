package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeSlowLogRecordsRequest struct {
	requests.RpcRequest
	SQLId                int64  `position:"Query" name:"SQLId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	EndTime              string `position:"Query" name:"EndTime"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
	DBName               string `position:"Query" name:"DBName"`
	PageSize             int    `position:"Query" name:"PageSize"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
}

func (req *DescribeSlowLogRecordsRequest) Invoke(client *sdk.Client) (resp *DescribeSlowLogRecordsResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeSlowLogRecords", "rds", "")
	resp = &DescribeSlowLogRecordsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeSlowLogRecordsResponse struct {
	responses.BaseResponse
	RequestId        string
	Engine           string
	TotalRecordCount int
	PageNumber       int
	PageRecordCount  int
	Items            DescribeSlowLogRecordsSQLSlowRecordList
}

type DescribeSlowLogRecordsSQLSlowRecord struct {
	HostAddress        string
	DBName             string
	SQLText            string
	QueryTimes         int64
	LockTimes          int64
	ParseRowCounts     int64
	ReturnRowCounts    int64
	ExecutionStartTime string
}

type DescribeSlowLogRecordsSQLSlowRecordList []DescribeSlowLogRecordsSQLSlowRecord

func (list *DescribeSlowLogRecordsSQLSlowRecordList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSlowLogRecordsSQLSlowRecord)
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
