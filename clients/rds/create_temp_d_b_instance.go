package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateTempDBInstanceRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	RestoreTime          string `position:"Query" name:"RestoreTime"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	BackupId             int    `position:"Query" name:"BackupId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *CreateTempDBInstanceRequest) Invoke(client *sdk.Client) (resp *CreateTempDBInstanceResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "CreateTempDBInstance", "rds", "")
	resp = &CreateTempDBInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateTempDBInstanceResponse struct {
	responses.BaseResponse
	RequestId        string
	TempDBInstanceId string
}
