package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CancelOrderRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
}

func (req *CancelOrderRequest) Invoke(client *sdk.Client) (resp *CancelOrderResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "CancelOrder", "", "")
	resp = &CancelOrderResponse{}
	err = client.DoAction(req, resp)
	return
}

type CancelOrderResponse struct {
	responses.BaseResponse
	RequestId common.String
	ClusterId common.String
}
