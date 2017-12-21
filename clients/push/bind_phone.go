package push

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type BindPhoneRequest struct {
	PhoneNumber string `position:"Query" name:"PhoneNumber"`
	AppKey      int64  `position:"Query" name:"AppKey"`
	DeviceId    string `position:"Query" name:"DeviceId"`
}

func (r BindPhoneRequest) Invoke(client *sdk.Client) (response *BindPhoneResponse, err error) {
	req := struct {
		*requests.RpcRequest
		BindPhoneRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Push", "2016-08-01", "BindPhone", "", "")

	resp := struct {
		*responses.BaseResponse
		BindPhoneResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.BindPhoneResponse

	err = client.DoAction(&req, &resp)
	return
}

type BindPhoneResponse struct {
	RequestId string
}
