package push

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UnbindPhoneRequest struct {
	AppKey   int64  `position:"Query" name:"AppKey"`
	DeviceId string `position:"Query" name:"DeviceId"`
}

func (r UnbindPhoneRequest) Invoke(client *sdk.Client) (response *UnbindPhoneResponse, err error) {
	req := struct {
		*requests.RpcRequest
		UnbindPhoneRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Push", "2016-08-01", "UnbindPhone", "", "")

	resp := struct {
		*responses.BaseResponse
		UnbindPhoneResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.UnbindPhoneResponse

	err = client.DoAction(&req, &resp)
	return
}

type UnbindPhoneResponse struct {
	RequestId string
}
