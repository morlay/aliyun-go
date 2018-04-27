package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateMigrateTaskForSQLServerRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	TaskType             string `position:"Query" name:"TaskType"`
	DBName               string `position:"Query" name:"DBName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	IsOnlineDB           string `position:"Query" name:"IsOnlineDB"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	OSSUrls              string `position:"Query" name:"OSSUrls"`
}

func (req *CreateMigrateTaskForSQLServerRequest) Invoke(client *sdk.Client) (resp *CreateMigrateTaskForSQLServerResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "CreateMigrateTaskForSQLServer", "rds", "")
	resp = &CreateMigrateTaskForSQLServerResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateMigrateTaskForSQLServerResponse struct {
	responses.BaseResponse
	RequestId      string
	DBInstanceId   string
	DBInstanceName string
	TaskId         string
	DBName         string
	MigrateIaskId  string
	TaskType       string
}
