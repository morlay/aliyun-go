package push

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type BindPhoneRequest struct {
	requests.RpcRequest
	PhoneNumber string `position:"Query" name:"PhoneNumber"`
	AppKey      int64  `position:"Query" name:"AppKey"`
	DeviceId    string `position:"Query" name:"DeviceId"`
}

func (req *BindPhoneRequest) Invoke(client *sdk.Client) (resp *BindPhoneResponse, err error) {
	req.InitWithApiInfo("Push", "2016-08-01", "BindPhone", "", "")
	resp = &BindPhoneResponse{}
	err = client.DoAction(req, resp)
	return
}

type BindPhoneResponse struct {
	responses.BaseResponse
	RequestId common.String
}
