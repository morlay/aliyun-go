package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CancelJobRequest struct {
	requests.RpcRequest
	JobId                string `position:"Query" name:"JobId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *CancelJobRequest) Invoke(client *sdk.Client) (resp *CancelJobResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "CancelJob", "mts", "")
	resp = &CancelJobResponse{}
	err = client.DoAction(req, resp)
	return
}

type CancelJobResponse struct {
	responses.BaseResponse
	RequestId common.String
	JobId     common.String
}
