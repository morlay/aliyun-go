package push

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UnbindPhoneRequest struct {
	requests.RpcRequest
	AppKey   int64  `position:"Query" name:"AppKey"`
	DeviceId string `position:"Query" name:"DeviceId"`
}

func (req *UnbindPhoneRequest) Invoke(client *sdk.Client) (resp *UnbindPhoneResponse, err error) {
	req.InitWithApiInfo("Push", "2016-08-01", "UnbindPhone", "", "")
	resp = &UnbindPhoneResponse{}
	err = client.DoAction(req, resp)
	return
}

type UnbindPhoneResponse struct {
	responses.BaseResponse
	RequestId string
}
