package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CheckRecoveryConditionsRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	RestoreTime          string `position:"Query" name:"RestoreTime"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	BackupFile           string `position:"Query" name:"BackupFile"`
	BackupId             string `position:"Query" name:"BackupId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r CheckRecoveryConditionsRequest) Invoke(client *sdk.Client) (response *CheckRecoveryConditionsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CheckRecoveryConditionsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "CheckRecoveryConditions", "rds", "")

	resp := struct {
		*responses.BaseResponse
		CheckRecoveryConditionsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CheckRecoveryConditionsResponse

	err = client.DoAction(&req, &resp)
	return
}

type CheckRecoveryConditionsResponse struct {
	RequestId      string
	DBInstanceId   string
	RecoveryStatus string
}
