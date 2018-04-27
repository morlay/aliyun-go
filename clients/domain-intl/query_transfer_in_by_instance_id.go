package domain_intl

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryTransferInByInstanceIdRequest struct {
	requests.RpcRequest
	InstanceId   string `position:"Query" name:"InstanceId"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
}

func (req *QueryTransferInByInstanceIdRequest) Invoke(client *sdk.Client) (resp *QueryTransferInByInstanceIdResponse, err error) {
	req.InitWithApiInfo("Domain-intl", "2017-12-18", "QueryTransferInByInstanceId", "domain", "")
	resp = &QueryTransferInByInstanceIdResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryTransferInByInstanceIdResponse struct {
	responses.BaseResponse
	RequestId                                   string
	SubmissionDate                              string
	ModificationDate                            string
	UserId                                      string
	InstanceId                                  string
	DomainName                                  string
	Status                                      int
	SimpleTransferInStatus                      string
	ResultCode                                  string
	ResultDate                                  string
	ResultMsg                                   string
	TransferAuthorizationCodeSubmissionDate     string
	NeedMailCheck                               bool
	Email                                       string
	WhoisMailStatus                             bool
	ExpirationDate                              string
	ProgressBarType                             int
	SubmissionDateLong                          int64
	ModificationDateLong                        int64
	ResultDateLong                              int64
	ExpirationDateLong                          int64
	TransferAuthorizationCodeSubmissionDateLong int64
}
