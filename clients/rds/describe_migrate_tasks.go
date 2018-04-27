package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeMigrateTasksRequest struct {
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

func (req *DescribeMigrateTasksRequest) Invoke(client *sdk.Client) (resp *DescribeMigrateTasksResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeMigrateTasks", "rds", "")
	resp = &DescribeMigrateTasksResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeMigrateTasksResponse struct {
	responses.BaseResponse
	RequestId        string
	DBInstanceId     string
	TotalRecordCount int
	PageNumber       int
	PageRecordCount  int
	Items            DescribeMigrateTasksMigrateTaskList
}

type DescribeMigrateTasksMigrateTask struct {
	DBName        string
	MigrateTaskId string
	CreateTime    string
	EndTime       string
	BackupMode    string
	Status        string
	IsDBReplaced  string
	Description   string
}

type DescribeMigrateTasksMigrateTaskList []DescribeMigrateTasksMigrateTask

func (list *DescribeMigrateTasksMigrateTaskList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeMigrateTasksMigrateTask)
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
