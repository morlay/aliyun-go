package polardb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateBackupRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string `position:"Query" name:"DBClusterId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *CreateBackupRequest) Invoke(client *sdk.Client) (resp *CreateBackupResponse, err error) {
	req.InitWithApiInfo("polardb", "2017-08-01", "CreateBackup", "polardb", "")
	resp = &CreateBackupResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateBackupResponse struct {
	responses.BaseResponse
	RequestId string
}
