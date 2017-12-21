package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyDBInstanceHAConfigRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	SyncMode             string `position:"Query" name:"SyncMode"`
	DbInstanceId         string `position:"Query" name:"DbInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	HAMode               string `position:"Query" name:"HAMode"`
}

func (r ModifyDBInstanceHAConfigRequest) Invoke(client *sdk.Client) (response *ModifyDBInstanceHAConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyDBInstanceHAConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "ModifyDBInstanceHAConfig", "rds", "")

	resp := struct {
		*responses.BaseResponse
		ModifyDBInstanceHAConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyDBInstanceHAConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyDBInstanceHAConfigResponse struct {
	RequestId string
}
