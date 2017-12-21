package polardb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyBackupPolicyRequest struct {
	PreferredBackupTime   string `position:"Query" name:"PreferredBackupTime"`
	PreferredBackupPeriod string `position:"Query" name:"PreferredBackupPeriod"`
	BackupRetentionPeriod string `position:"Query" name:"BackupRetentionPeriod"`
	ResourceOwnerId       int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId           string `position:"Query" name:"DBClusterId"`
	OwnerAccount          string `position:"Query" name:"OwnerAccount"`
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
	req.InitWithApiInfo("polardb", "2017-08-01", "ModifyBackupPolicy", "polardb", "")

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
