package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateBackupRequest struct {
	BackupMethod         string `position:"Query" name:"BackupMethod"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	DBName               string `position:"Query" name:"DBName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	BackupType           string `position:"Query" name:"BackupType"`
}

func (r CreateBackupRequest) Invoke(client *sdk.Client) (response *CreateBackupResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateBackupRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "CreateBackup", "rds", "")

	resp := struct {
		*responses.BaseResponse
		CreateBackupResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateBackupResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateBackupResponse struct {
	RequestId string
}
