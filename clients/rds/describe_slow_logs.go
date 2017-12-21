package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeSlowLogsRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	EndTime              string `position:"Query" name:"EndTime"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
	SortKey              string `position:"Query" name:"SortKey"`
	DBName               string `position:"Query" name:"DBName"`
	PageSize             int    `position:"Query" name:"PageSize"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
}

func (r DescribeSlowLogsRequest) Invoke(client *sdk.Client) (response *DescribeSlowLogsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeSlowLogsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeSlowLogs", "rds", "")

	resp := struct {
		*responses.BaseResponse
		DescribeSlowLogsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeSlowLogsResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeSlowLogsResponse struct {
	RequestId        string
	Engine           string
	StartTime        string
	EndTime          string
	TotalRecordCount int
	PageNumber       int
	PageRecordCount  int
	Items            DescribeSlowLogsSQLSlowLogList
}

type DescribeSlowLogsSQLSlowLog struct {
	SlowLogId                     int64
	SQLId                         int64
	DBName                        string
	SQLText                       string
	MySQLTotalExecutionCounts     int64
	MySQLTotalExecutionTimes      int64
	TotalLockTimes                int64
	MaxLockTime                   int64
	ParseTotalRowCounts           int64
	ParseMaxRowCount              int64
	ReturnTotalRowCounts          int64
	ReturnMaxRowCount             int64
	CreateTime                    string
	SQLServerTotalExecutionCounts int64
	SQLServerTotalExecutionTimes  int64
	TotalLogicalReadCounts        int64
	TotalPhysicalReadCounts       int64
	ReportTime                    string
	MaxExecutionTime              int64
	AvgExecutionTime              int64
}

type DescribeSlowLogsSQLSlowLogList []DescribeSlowLogsSQLSlowLog

func (list *DescribeSlowLogsSQLSlowLogList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSlowLogsSQLSlowLog)
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
