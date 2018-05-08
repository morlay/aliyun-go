package domain

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDRequest struct {
	requests.RpcRequest
	InstanceId          string `position:"Query" name:"InstanceId"`
	UserClientIp        string `position:"Query" name:"UserClientIp"`
	DomainName          string `position:"Query" name:"DomainName"`
	RegistrantProfileId int64  `position:"Query" name:"RegistrantProfileId"`
	Lang                string `position:"Query" name:"Lang"`
}

func (req *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDRequest) Invoke(client *sdk.Client) (resp *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDResponse, err error) {
	req.InitWithApiInfo("Domain", "2018-01-29", "SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileID", "", "")
	resp = &SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDResponse{}
	err = client.DoAction(req, resp)
	return
}

type SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDResponse struct {
	responses.BaseResponse
	RequestId common.String
	TaskNo    common.String
}
