package domain

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type TransferInReenterTransferAuthorizationCodeRequest struct {
	requests.RpcRequest
	TransferAuthorizationCode string `position:"Query" name:"TransferAuthorizationCode"`
	DomainName                string `position:"Query" name:"DomainName"`
	UserClientIp              string `position:"Query" name:"UserClientIp"`
	Lang                      string `position:"Query" name:"Lang"`
}

func (req *TransferInReenterTransferAuthorizationCodeRequest) Invoke(client *sdk.Client) (resp *TransferInReenterTransferAuthorizationCodeResponse, err error) {
	req.InitWithApiInfo("Domain", "2018-01-29", "TransferInReenterTransferAuthorizationCode", "", "")
	resp = &TransferInReenterTransferAuthorizationCodeResponse{}
	err = client.DoAction(req, resp)
	return
}

type TransferInReenterTransferAuthorizationCodeResponse struct {
	responses.BaseResponse
	RequestId string
}
