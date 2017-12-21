package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type BindMFADeviceRequest struct {
	requests.RpcRequest
	SerialNumber        string `position:"Query" name:"SerialNumber"`
	AuthenticationCode2 string `position:"Query" name:"AuthenticationCode.2"`
	AuthenticationCode1 string `position:"Query" name:"AuthenticationCode.1"`
	UserName            string `position:"Query" name:"UserName"`
}

func (req *BindMFADeviceRequest) Invoke(client *sdk.Client) (resp *BindMFADeviceResponse, err error) {
	req.InitWithApiInfo("Ram", "2015-05-01", "BindMFADevice", "", "")
	resp = &BindMFADeviceResponse{}
	err = client.DoAction(req, resp)
	return
}

type BindMFADeviceResponse struct {
	responses.BaseResponse
	RequestId string
}
