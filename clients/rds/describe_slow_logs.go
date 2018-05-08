package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeSlowLogsRequest struct {
	requests.RpcRequest
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

func (req *DescribeSlowLogsRequest) Invoke(client *sdk.Client) (resp *DescribeSlowLogsResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeSlowLogs", "rds", "")
	resp = &DescribeSlowLogsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeSlowLogsResponse struct {
	responses.BaseResponse
	RequestId        common.String
	Engine           common.String
	StartTime        common.String
	EndTime          common.String
	TotalRecordCount common.Integer
	PageNumber       common.Integer
	PageRecordCount  common.Integer
	Items            DescribeSlowLogsSQLSlowLogList
}

type DescribeSlowLogsSQLSlowLog struct {
	SlowLogId                     common.Long
	SQLId                         common.Long
	DBName                        common.String
	SQLText                       common.String
	MySQLTotalExecutionCounts     common.Long
	MySQLTotalExecutionTimes      common.Long
	TotalLockTimes                common.Long
	MaxLockTime                   common.Long
	ParseTotalRowCounts           common.Long
	ParseMaxRowCount              common.Long
	ReturnTotalRowCounts          common.Long
	ReturnMaxRowCount             common.Long
	CreateTime                    common.String
	SQLServerTotalExecutionCounts common.Long
	SQLServerTotalExecutionTimes  common.Long
	TotalLogicalReadCounts        common.Long
	TotalPhysicalReadCounts       common.Long
	ReportTime                    common.String
	MaxExecutionTime              common.Long
	AvgExecutionTime              common.Long
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
