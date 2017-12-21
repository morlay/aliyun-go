package r_kvstore

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateBackupRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *CreateBackupRequest) Invoke(client *sdk.Client) (resp *CreateBackupResponse, err error) {
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "CreateBackup", "redisa", "")
	resp = &CreateBackupResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateBackupResponse struct {
	responses.BaseResponse
	RequestId   string
	BackupJobID string
}
