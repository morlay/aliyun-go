package domain

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type TransferInResendMailTokenRequest struct {
	requests.RpcRequest
	DomainName   string `position:"Query" name:"DomainName"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
}

func (req *TransferInResendMailTokenRequest) Invoke(client *sdk.Client) (resp *TransferInResendMailTokenResponse, err error) {
	req.InitWithApiInfo("Domain", "2018-01-29", "TransferInResendMailToken", "", "")
	resp = &TransferInResendMailTokenResponse{}
	err = client.DoAction(req, resp)
	return
}

type TransferInResendMailTokenResponse struct {
	responses.BaseResponse
	RequestId common.String
}
