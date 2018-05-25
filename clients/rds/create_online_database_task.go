package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateOnlineDatabaseTaskRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	MigrateTaskId        string `position:"Query" name:"MigrateTaskId"`
	DBName               string `position:"Query" name:"DBName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	CheckDBMode          string `position:"Query" name:"CheckDBMode"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *CreateOnlineDatabaseTaskRequest) Invoke(client *sdk.Client) (resp *CreateOnlineDatabaseTaskResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "CreateOnlineDatabaseTask", "rds", "")
	resp = &CreateOnlineDatabaseTaskResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateOnlineDatabaseTaskResponse struct {
	responses.BaseResponse
	RequestId common.String
}
