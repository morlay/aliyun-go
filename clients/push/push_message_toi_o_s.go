package push

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type PushMessageToiOSRequest struct {
	requests.RpcRequest
	AppKey      int64  `position:"Query" name:"AppKey"`
	TargetValue string `position:"Query" name:"TargetValue"`
	Title       string `position:"Query" name:"Title"`
	Body        string `position:"Query" name:"Body"`
	JobKey      string `position:"Query" name:"JobKey"`
	Target      string `position:"Query" name:"Target"`
}

func (req *PushMessageToiOSRequest) Invoke(client *sdk.Client) (resp *PushMessageToiOSResponse, err error) {
	req.InitWithApiInfo("Push", "2016-08-01", "PushMessageToiOS", "", "")
	resp = &PushMessageToiOSResponse{}
	err = client.DoAction(req, resp)
	return
}

type PushMessageToiOSResponse struct {
	responses.BaseResponse
	RequestId common.String
	MessageId common.String
}
