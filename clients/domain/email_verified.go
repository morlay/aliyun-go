package domain

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type EmailVerifiedRequest struct {
	requests.RpcRequest
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
	Email        string `position:"Query" name:"Email"`
}

func (req *EmailVerifiedRequest) Invoke(client *sdk.Client) (resp *EmailVerifiedResponse, err error) {
	req.InitWithApiInfo("Domain", "2018-01-29", "EmailVerified", "", "")
	resp = &EmailVerifiedResponse{}
	err = client.DoAction(req, resp)
	return
}

type EmailVerifiedResponse struct {
	responses.BaseResponse
	RequestId string
}
