package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type BindMFADeviceRequest struct {
	SerialNumber        string `position:"Query" name:"SerialNumber"`
	AuthenticationCode2 string `position:"Query" name:"AuthenticationCode.2"`
	AuthenticationCode1 string `position:"Query" name:"AuthenticationCode.1"`
	UserName            string `position:"Query" name:"UserName"`
}

func (r BindMFADeviceRequest) Invoke(client *sdk.Client) (response *BindMFADeviceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		BindMFADeviceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ram", "2015-05-01", "BindMFADevice", "", "")

	resp := struct {
		*responses.BaseResponse
		BindMFADeviceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.BindMFADeviceResponse

	err = client.DoAction(&req, &resp)
	return
}

type BindMFADeviceResponse struct {
	RequestId string
}
