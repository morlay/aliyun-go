package dds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeBackupPolicyRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeBackupPolicyRequest) Invoke(client *sdk.Client) (resp *DescribeBackupPolicyResponse, err error) {
	req.InitWithApiInfo("Dds", "2015-12-01", "DescribeBackupPolicy", "dds", "")
	resp = &DescribeBackupPolicyResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeBackupPolicyResponse struct {
	responses.BaseResponse
	RequestId             common.String
	BackupRetentionPeriod common.String
	PreferredBackupTime   common.String
	PreferredBackupPeriod common.String
}
