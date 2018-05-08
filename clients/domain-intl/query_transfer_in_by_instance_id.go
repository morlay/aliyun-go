package domain_intl

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId                                   common.String
	SubmissionDate                              common.String
	ModificationDate                            common.String
	UserId                                      common.String
	InstanceId                                  common.String
	DomainName                                  common.String
	Status                                      common.Integer
	SimpleTransferInStatus                      common.String
	ResultCode                                  common.String
	ResultDate                                  common.String
	ResultMsg                                   common.String
	TransferAuthorizationCodeSubmissionDate     common.String
	NeedMailCheck                               bool
	Email                                       common.String
	WhoisMailStatus                             bool
	ExpirationDate                              common.String
	ProgressBarType                             common.Integer
	SubmissionDateLong                          common.Long
	ModificationDateLong                        common.Long
	ResultDateLong                              common.Long
	ExpirationDateLong                          common.Long
	TransferAuthorizationCodeSubmissionDateLong common.Long
}
