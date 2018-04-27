package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeMigrateTasksForSQLServerRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	PageSize             int    `position:"Query" name:"PageSize"`
	EndTime              string `position:"Query" name:"EndTime"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
}

func (req *DescribeMigrateTasksForSQLServerRequest) Invoke(client *sdk.Client) (resp *DescribeMigrateTasksForSQLServerResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeMigrateTasksForSQLServer", "rds", "")
	resp = &DescribeMigrateTasksForSQLServerResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeMigrateTasksForSQLServerResponse struct {
	responses.BaseResponse
	RequestId        string
	DBInstanceID     string
	DBInstanceName   string
	StartTime        string
	EndTime          string
	TotalRecordCount int
	PageNumber       int
	PageRecordCount  int
	Items            DescribeMigrateTasksForSQLServerMigrateTaskList
}

type DescribeMigrateTasksForSQLServerMigrateTask struct {
	DBName        string
	MigrateIaskId string
	CreateTime    string
	EndTime       string
	TaskType      string
	Status        string
	IsDBReplaced  string
	Desc          string
}

type DescribeMigrateTasksForSQLServerMigrateTaskList []DescribeMigrateTasksForSQLServerMigrateTask

func (list *DescribeMigrateTasksForSQLServerMigrateTaskList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeMigrateTasksForSQLServerMigrateTask)
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
