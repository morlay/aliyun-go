package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId        common.String
	Engine           common.String
	TotalRecordCount common.Integer
	PageNumber       common.Integer
	PageRecordCount  common.Integer
	Items            DescribeSlowLogRecordsSQLSlowRecordList
}

type DescribeSlowLogRecordsSQLSlowRecord struct {
	HostAddress        common.String
	DBName             common.String
	SQLText            common.String
	QueryTimes         common.Long
	LockTimes          common.Long
	ParseRowCounts     common.Long
	ReturnRowCounts    common.Long
	ExecutionStartTime common.String
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
