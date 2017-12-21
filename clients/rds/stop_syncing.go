package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type StopSyncingRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ImportId             int    `position:"Query" name:"ImportId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r StopSyncingRequest) Invoke(client *sdk.Client) (response *StopSyncingResponse, err error) {
	req := struct {
		*requests.RpcRequest
		StopSyncingRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "StopSyncing", "rds", "")

	resp := struct {
		*responses.BaseResponse
		StopSyncingResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.StopSyncingResponse

	err = client.DoAction(&req, &resp)
	return
}

type StopSyncingResponse struct {
	RequestId string
}
