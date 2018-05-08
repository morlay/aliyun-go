package push

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CancelPushRequest struct {
	requests.RpcRequest
	MessageId int64 `position:"Query" name:"MessageId"`
	AppKey    int64 `position:"Query" name:"AppKey"`
}

func (req *CancelPushRequest) Invoke(client *sdk.Client) (resp *CancelPushResponse, err error) {
	req.InitWithApiInfo("Push", "2016-08-01", "CancelPush", "", "")
	resp = &CancelPushResponse{}
	err = client.DoAction(req, resp)
	return
}

type CancelPushResponse struct {
	responses.BaseResponse
	RequestId common.String
}
