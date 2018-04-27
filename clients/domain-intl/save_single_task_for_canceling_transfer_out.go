package domain_intl

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SaveSingleTaskForCancelingTransferOutRequest struct {
	requests.RpcRequest
	DomainName   string `position:"Query" name:"DomainName"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
}

func (req *SaveSingleTaskForCancelingTransferOutRequest) Invoke(client *sdk.Client) (resp *SaveSingleTaskForCancelingTransferOutResponse, err error) {
	req.InitWithApiInfo("Domain-intl", "2017-12-18", "SaveSingleTaskForCancelingTransferOut", "domain", "")
	resp = &SaveSingleTaskForCancelingTransferOutResponse{}
	err = client.DoAction(req, resp)
	return
}

type SaveSingleTaskForCancelingTransferOutResponse struct {
	responses.BaseResponse
	RequestId string
	TaskNo    string
}
