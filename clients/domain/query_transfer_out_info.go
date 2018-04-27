package domain

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
	req.InitWithApiInfo("Domain", "2018-01-29", "QueryTransferOutInfo", "", "")
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
