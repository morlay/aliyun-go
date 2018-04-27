package domain_intl

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryTransferOutInfoRequest struct {
	requests.RpcRequest
	DomainName   string `position:"Query" name:"DomainName"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
}

func (req *QueryTransferOutInfoRequest) Invoke(client *sdk.Client) (resp *QueryTransferOutInfoResponse, err error) {
	req.InitWithApiInfo("Domain-intl", "2017-12-18", "QueryTransferOutInfo", "domain", "")
	resp = &QueryTransferOutInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryTransferOutInfoResponse struct {
	responses.BaseResponse
	RequestId                         string
	Status                            int
	Email                             string
	TransferAuthorizationCodeSendDate string
	ExpirationDate                    string
	PendingRequestDate                string
	ResultCode                        string
	ResultMsg                         string
}
