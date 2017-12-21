package push

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type PushMessageToiOSRequest struct {
	AppKey      int64  `position:"Query" name:"AppKey"`
	TargetValue string `position:"Query" name:"TargetValue"`
	Title       string `position:"Query" name:"Title"`
	Body        string `position:"Query" name:"Body"`
	JobKey      string `position:"Query" name:"JobKey"`
	Target      string `position:"Query" name:"Target"`
}

func (r PushMessageToiOSRequest) Invoke(client *sdk.Client) (response *PushMessageToiOSResponse, err error) {
	req := struct {
		*requests.RpcRequest
		PushMessageToiOSRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Push", "2016-08-01", "PushMessageToiOS", "", "")

	resp := struct {
		*responses.BaseResponse
		PushMessageToiOSResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.PushMessageToiOSResponse

	err = client.DoAction(&req, &resp)
	return
}

type PushMessageToiOSResponse struct {
	RequestId string
	MessageId string
}
