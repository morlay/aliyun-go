package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type UpgradeDBInstanceKernelVersionRequest struct {
	requests.RpcRequest
	SwitchTimeMode       string `position:"Query" name:"SwitchTimeMode"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	SwitchTime           string `position:"Query" name:"SwitchTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *UpgradeDBInstanceKernelVersionRequest) Invoke(client *sdk.Client) (resp *UpgradeDBInstanceKernelVersionResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "UpgradeDBInstanceKernelVersion", "rds", "")
	resp = &UpgradeDBInstanceKernelVersionResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpgradeDBInstanceKernelVersionResponse struct {
	responses.BaseResponse
	RequestId          common.String
	DBInstanceName     common.String
	TaskId             common.String
	TargetMinorVersion common.String
}
