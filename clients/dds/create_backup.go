package dds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateBackupRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r CreateBackupRequest) Invoke(client *sdk.Client) (response *CreateBackupResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateBackupRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Dds", "2015-12-01", "CreateBackup", "dds", "")

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
	BackupId  string
}
