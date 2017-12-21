package polardb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateBackupRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string `position:"Query" name:"DBClusterId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
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
	req.InitWithApiInfo("polardb", "2017-08-01", "CreateBackup", "polardb", "")

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
