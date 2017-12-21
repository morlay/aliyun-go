package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyDBInstanceTDERequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	DBName               string `position:"Query" name:"DBName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	TDEStatus            string `position:"Query" name:"TDEStatus"`
}

func (req *ModifyDBInstanceTDERequest) Invoke(client *sdk.Client) (resp *ModifyDBInstanceTDEResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "ModifyDBInstanceTDE", "rds", "")
	resp = &ModifyDBInstanceTDEResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyDBInstanceTDEResponse struct {
	responses.BaseResponse
	RequestId string
}
