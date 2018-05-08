package domain_intl

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SaveSingleTaskForCancelingTransferInRequest struct {
	requests.RpcRequest
	DomainName   string `position:"Query" name:"DomainName"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
}

func (req *SaveSingleTaskForCancelingTransferInRequest) Invoke(client *sdk.Client) (resp *SaveSingleTaskForCancelingTransferInResponse, err error) {
	req.InitWithApiInfo("Domain-intl", "2017-12-18", "SaveSingleTaskForCancelingTransferIn", "domain", "")
	resp = &SaveSingleTaskForCancelingTransferInResponse{}
	err = client.DoAction(req, resp)
	return
}

type SaveSingleTaskForCancelingTransferInResponse struct {
	responses.BaseResponse
	RequestId common.String
	TaskNo    common.String
}
