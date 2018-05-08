package domain

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SaveSingleTaskForQueryingTransferAuthorizationCodeRequest struct {
	requests.RpcRequest
	DomainName   string `position:"Query" name:"DomainName"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
}

func (req *SaveSingleTaskForQueryingTransferAuthorizationCodeRequest) Invoke(client *sdk.Client) (resp *SaveSingleTaskForQueryingTransferAuthorizationCodeResponse, err error) {
	req.InitWithApiInfo("Domain", "2018-01-29", "SaveSingleTaskForQueryingTransferAuthorizationCode", "", "")
	resp = &SaveSingleTaskForQueryingTransferAuthorizationCodeResponse{}
	err = client.DoAction(req, resp)
	return
}

type SaveSingleTaskForQueryingTransferAuthorizationCodeResponse struct {
	responses.BaseResponse
	RequestId common.String
	TaskNo    common.String
}
