package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type StopSyncingRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ImportId             int    `position:"Query" name:"ImportId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *StopSyncingRequest) Invoke(client *sdk.Client) (resp *StopSyncingResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "StopSyncing", "rds", "")
	resp = &StopSyncingResponse{}
	err = client.DoAction(req, resp)
	return
}

type StopSyncingResponse struct {
	responses.BaseResponse
	RequestId string
}
