package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteBackupRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	BackupId             string `position:"Query" name:"BackupId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DeleteBackupRequest) Invoke(client *sdk.Client) (response *DeleteBackupResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteBackupRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "DeleteBackup", "rds", "")

	resp := struct {
		*responses.BaseResponse
		DeleteBackupResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteBackupResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteBackupResponse struct {
	RequestId string
}
