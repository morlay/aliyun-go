package polardb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeBackupPolicyRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string `position:"Query" name:"DBClusterId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeBackupPolicyRequest) Invoke(client *sdk.Client) (response *DescribeBackupPolicyResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeBackupPolicyRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("polardb", "2017-08-01", "DescribeBackupPolicy", "polardb", "")

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
	RequestId               string
	BackupRetentionPeriod   int
	PreferredNextBackupTime string
	PreferredBackupTime     string
	PreferredBackupPeriod   string
}
