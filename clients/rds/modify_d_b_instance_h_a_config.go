package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyDBInstanceHAConfigRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	SyncMode             string `position:"Query" name:"SyncMode"`
	DbInstanceId         string `position:"Query" name:"DbInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	HAMode               string `position:"Query" name:"HAMode"`
}

func (req *ModifyDBInstanceHAConfigRequest) Invoke(client *sdk.Client) (resp *ModifyDBInstanceHAConfigResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "ModifyDBInstanceHAConfig", "rds", "")
	resp = &ModifyDBInstanceHAConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyDBInstanceHAConfigResponse struct {
	responses.BaseResponse
	RequestId string
}
