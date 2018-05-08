package domain_intl

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type VerifyEmailRequest struct {
	requests.RpcRequest
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
	Token        string `position:"Query" name:"Token"`
}

func (req *VerifyEmailRequest) Invoke(client *sdk.Client) (resp *VerifyEmailResponse, err error) {
	req.InitWithApiInfo("Domain-intl", "2017-12-18", "VerifyEmail", "domain", "")
	resp = &VerifyEmailResponse{}
	err = client.DoAction(req, resp)
	return
}

type VerifyEmailResponse struct {
	responses.BaseResponse
	RequestId common.String
}
