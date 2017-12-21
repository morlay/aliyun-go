package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteBackupRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	BackupId             string `position:"Query" name:"BackupId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DeleteBackupRequest) Invoke(client *sdk.Client) (resp *DeleteBackupResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DeleteBackup", "rds", "")
	resp = &DeleteBackupResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteBackupResponse struct {
	responses.BaseResponse
	RequestId string
}
