
package r-kvstore

import (
    "github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
)


type DescribeBackupPolicyRequest struct {
ResourceOwnerId int64 `position:"Query" name:"ResourceOwnerId"`
InstanceId string `position:"Query" name:"InstanceId"`
SecurityToken string `position:"Query" name:"SecurityToken"`
ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
OwnerAccount string `position:"Query" name:"OwnerAccount"`
OwnerId int64 `position:"Query" name:"OwnerId"`
}

func (r DescribeBackupPolicyRequest) Invoke(client *sdk.Client) (response *DescribeBackupPolicyResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeBackupPolicyRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "DescribeBackupPolicy", "redisa", "")

	resp := struct {
		*responses.BaseResponse
		DescribeBackupPolicyResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
    response = &resp.DescribeBackupPolicyResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeBackupPolicyResponse struct {
RequestId string
BackupRetentionPeriod string
PreferredBackupTime string
PreferredBackupPeriod string
PreferredNextBackupTime string
}

