
package r-kvstore

import (
    "github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
)


type ModifyBackupPolicyRequest struct {
PreferredBackupTime string `position:"Query" name:"PreferredBackupTime"`
PreferredBackupPeriod string `position:"Query" name:"PreferredBackupPeriod"`
ResourceOwnerId int64 `position:"Query" name:"ResourceOwnerId"`
InstanceId string `position:"Query" name:"InstanceId"`
SecurityToken string `position:"Query" name:"SecurityToken"`
ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
OwnerAccount string `position:"Query" name:"OwnerAccount"`
OwnerId int64 `position:"Query" name:"OwnerId"`
}

func (r ModifyBackupPolicyRequest) Invoke(client *sdk.Client) (response *ModifyBackupPolicyResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyBackupPolicyRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "ModifyBackupPolicy", "redisa", "")

	resp := struct {
		*responses.BaseResponse
		ModifyBackupPolicyResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
    response = &resp.ModifyBackupPolicyResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyBackupPolicyResponse struct {
RequestId string
}

