package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteFCTriggerRequest struct {
	requests.RpcRequest
	TriggerARN string `position:"Query" name:"TriggerARN"`
	OwnerId    int64  `position:"Query" name:"OwnerId"`
}

func (req *DeleteFCTriggerRequest) Invoke(client *sdk.Client) (resp *DeleteFCTriggerResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DeleteFCTrigger", "", "")
	resp = &DeleteFCTriggerResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteFCTriggerResponse struct {
	responses.BaseResponse
	RequestId common.String
}
