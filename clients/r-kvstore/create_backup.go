
package r-kvstore

import (
    "github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
)


type CreateBackupRequest struct {
ResourceOwnerId int64 `position:"Query" name:"ResourceOwnerId"`
InstanceId string `position:"Query" name:"InstanceId"`
SecurityToken string `position:"Query" name:"SecurityToken"`
ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
OwnerAccount string `position:"Query" name:"OwnerAccount"`
OwnerId int64 `position:"Query" name:"OwnerId"`
}

func (r CreateBackupRequest) Invoke(client *sdk.Client) (response *CreateBackupResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateBackupRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "CreateBackup", "redisa", "")

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
BackupJobID string
}

