package domain

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type VerifyEmailRequest struct {
	requests.RpcRequest
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
	Token        string `position:"Query" name:"Token"`
}

func (req *VerifyEmailRequest) Invoke(client *sdk.Client) (resp *VerifyEmailResponse, err error) {
	req.InitWithApiInfo("Domain", "2018-01-29", "VerifyEmail", "", "")
	resp = &VerifyEmailResponse{}
	err = client.DoAction(req, resp)
	return
}

type VerifyEmailResponse struct {
	responses.BaseResponse
	RequestId string
}
