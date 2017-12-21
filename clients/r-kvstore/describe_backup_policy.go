package r_kvstore

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeBackupPolicyRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeBackupPolicyRequest) Invoke(client *sdk.Client) (resp *DescribeBackupPolicyResponse, err error) {
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "DescribeBackupPolicy", "redisa", "")
	resp = &DescribeBackupPolicyResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeBackupPolicyResponse struct {
	responses.BaseResponse
	RequestId               string
	BackupRetentionPeriod   string
	PreferredBackupTime     string
	PreferredBackupPeriod   string
	PreferredNextBackupTime string
}
