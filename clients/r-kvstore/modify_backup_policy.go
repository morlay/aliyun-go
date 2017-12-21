package r_kvstore

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyBackupPolicyRequest struct {
	requests.RpcRequest
	PreferredBackupTime   string `position:"Query" name:"PreferredBackupTime"`
	PreferredBackupPeriod string `position:"Query" name:"PreferredBackupPeriod"`
	ResourceOwnerId       int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId            string `position:"Query" name:"InstanceId"`
	SecurityToken         string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount          string `position:"Query" name:"OwnerAccount"`
	OwnerId               int64  `position:"Query" name:"OwnerId"`
}

func (req *ModifyBackupPolicyRequest) Invoke(client *sdk.Client) (resp *ModifyBackupPolicyResponse, err error) {
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "ModifyBackupPolicy", "redisa", "")
	resp = &ModifyBackupPolicyResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyBackupPolicyResponse struct {
	responses.BaseResponse
	RequestId string
}
