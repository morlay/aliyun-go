package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId        common.String
	DBInstanceId     common.String
	TotalRecordCount common.Integer
	PageNumber       common.Integer
	PageRecordCount  common.Integer
	Items            DescribeMigrateTasksMigrateTaskList
}

type DescribeMigrateTasksMigrateTask struct {
	DBName        common.String
	MigrateTaskId common.String
	CreateTime    common.String
	EndTime       common.String
	BackupMode    common.String
	Status        common.String
	IsDBReplaced  common.String
	Description   common.String
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
