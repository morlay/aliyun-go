package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CheckRecoveryConditionsRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	RestoreTime          string `position:"Query" name:"RestoreTime"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	BackupFile           string `position:"Query" name:"BackupFile"`
	BackupId             string `position:"Query" name:"BackupId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *CheckRecoveryConditionsRequest) Invoke(client *sdk.Client) (resp *CheckRecoveryConditionsResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "CheckRecoveryConditions", "rds", "")
	resp = &CheckRecoveryConditionsResponse{}
	err = client.DoAction(req, resp)
	return
}

type CheckRecoveryConditionsResponse struct {
	responses.BaseResponse
	RequestId      common.String
	DBInstanceId   common.String
	RecoveryStatus common.String
}
