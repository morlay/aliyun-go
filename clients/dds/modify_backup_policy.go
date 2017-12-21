package dds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyBackupPolicyRequest struct {
	PreferredBackupTime   string `position:"Query" name:"PreferredBackupTime"`
	PreferredBackupPeriod string `position:"Query" name:"PreferredBackupPeriod"`
	ResourceOwnerId       int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken         string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount          string `position:"Query" name:"OwnerAccount"`
	DBInstanceId          string `position:"Query" name:"DBInstanceId"`
	OwnerId               int64  `position:"Query" name:"OwnerId"`
}

func (r ModifyBackupPolicyRequest) Invoke(client *sdk.Client) (response *ModifyBackupPolicyResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyBackupPolicyRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Dds", "2015-12-01", "ModifyBackupPolicy", "dds", "")

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
