package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateMigrateTaskRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	DBName               string `position:"Query" name:"DBName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	IsOnlineDB           string `position:"Query" name:"IsOnlineDB"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	BackupMode           string `position:"Query" name:"BackupMode"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	OSSUrls              string `position:"Query" name:"OSSUrls"`
}

func (req *CreateMigrateTaskRequest) Invoke(client *sdk.Client) (resp *CreateMigrateTaskResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "CreateMigrateTask", "rds", "")
	resp = &CreateMigrateTaskResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateMigrateTaskResponse struct {
	responses.BaseResponse
	RequestId     common.String
	DBInstanceId  common.String
	TaskId        common.String
	DBName        common.String
	MigrateIaskId common.String
	BackupMode    common.String
}
