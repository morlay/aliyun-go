package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyDBInstanceConnectionModeRequest struct {
	requests.RpcRequest
	ConnectionMode       string `position:"Query" name:"ConnectionMode"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ModifyDBInstanceConnectionModeRequest) Invoke(client *sdk.Client) (resp *ModifyDBInstanceConnectionModeResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "ModifyDBInstanceConnectionMode", "rds", "")
	resp = &ModifyDBInstanceConnectionModeResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyDBInstanceConnectionModeResponse struct {
	responses.BaseResponse
	RequestId string
}
