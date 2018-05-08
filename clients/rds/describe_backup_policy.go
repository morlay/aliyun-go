package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeBackupPolicyRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeBackupPolicyRequest) Invoke(client *sdk.Client) (resp *DescribeBackupPolicyResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeBackupPolicy", "rds", "")
	resp = &DescribeBackupPolicyResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeBackupPolicyResponse struct {
	responses.BaseResponse
	RequestId                common.String
	BackupRetentionPeriod    common.Integer
	PreferredNextBackupTime  common.String
	PreferredBackupTime      common.String
	PreferredBackupPeriod    common.String
	BackupLog                common.String
	LogBackupRetentionPeriod common.Integer
}
