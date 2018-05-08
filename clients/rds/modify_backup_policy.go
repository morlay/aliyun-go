package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyBackupPolicyRequest struct {
	requests.RpcRequest
	PreferredBackupTime      string `position:"Query" name:"PreferredBackupTime"`
	PreferredBackupPeriod    string `position:"Query" name:"PreferredBackupPeriod"`
	BackupRetentionPeriod    string `position:"Query" name:"BackupRetentionPeriod"`
	ResourceOwnerId          int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount             string `position:"Query" name:"OwnerAccount"`
	DBInstanceId             string `position:"Query" name:"DBInstanceId"`
	BackupLog                string `position:"Query" name:"BackupLog"`
	OwnerId                  int64  `position:"Query" name:"OwnerId"`
	LogBackupRetentionPeriod string `position:"Query" name:"LogBackupRetentionPeriod"`
}

func (req *ModifyBackupPolicyRequest) Invoke(client *sdk.Client) (resp *ModifyBackupPolicyResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "ModifyBackupPolicy", "rds", "")
	resp = &ModifyBackupPolicyResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyBackupPolicyResponse struct {
	responses.BaseResponse
	RequestId common.String
}
