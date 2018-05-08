package domain

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId                         common.String
	Status                            common.Integer
	Email                             common.String
	TransferAuthorizationCodeSendDate common.String
	ExpirationDate                    common.String
	PendingRequestDate                common.String
	ResultCode                        common.String
	ResultMsg                         common.String
}
