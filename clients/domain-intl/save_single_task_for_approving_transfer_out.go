package domain_intl

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SaveSingleTaskForApprovingTransferOutRequest struct {
	requests.RpcRequest
	DomainName   string `position:"Query" name:"DomainName"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
}

func (req *SaveSingleTaskForApprovingTransferOutRequest) Invoke(client *sdk.Client) (resp *SaveSingleTaskForApprovingTransferOutResponse, err error) {
	req.InitWithApiInfo("Domain-intl", "2017-12-18", "SaveSingleTaskForApprovingTransferOut", "domain", "")
	resp = &SaveSingleTaskForApprovingTransferOutResponse{}
	err = client.DoAction(req, resp)
	return
}

type SaveSingleTaskForApprovingTransferOutResponse struct {
	responses.BaseResponse
	RequestId string
	TaskNo    string
}
