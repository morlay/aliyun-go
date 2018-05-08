package domain_intl

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SaveSingleTaskForCreatingOrderTransferRequest struct {
	requests.RpcRequest
	PermitPremiumTransfer string `position:"Query" name:"PermitPremiumTransfer"`
	AuthorizationCode     string `position:"Query" name:"AuthorizationCode"`
	UserClientIp          string `position:"Query" name:"UserClientIp"`
	DomainName            string `position:"Query" name:"DomainName"`
	RegistrantProfileId   int64  `position:"Query" name:"RegistrantProfileId"`
	Lang                  string `position:"Query" name:"Lang"`
}

func (req *SaveSingleTaskForCreatingOrderTransferRequest) Invoke(client *sdk.Client) (resp *SaveSingleTaskForCreatingOrderTransferResponse, err error) {
	req.InitWithApiInfo("Domain-intl", "2017-12-18", "SaveSingleTaskForCreatingOrderTransfer", "domain", "")
	resp = &SaveSingleTaskForCreatingOrderTransferResponse{}
	err = client.DoAction(req, resp)
	return
}

type SaveSingleTaskForCreatingOrderTransferResponse struct {
	responses.BaseResponse
	RequestId common.String
	TaskNo    common.String
}
