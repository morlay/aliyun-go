package domain_intl

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SaveTaskForSubmittingDomainDeleteRequest struct {
	requests.RpcRequest
	InstanceId   string `position:"Query" name:"InstanceId"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
}

func (req *SaveTaskForSubmittingDomainDeleteRequest) Invoke(client *sdk.Client) (resp *SaveTaskForSubmittingDomainDeleteResponse, err error) {
	req.InitWithApiInfo("Domain-intl", "2017-12-18", "SaveTaskForSubmittingDomainDelete", "domain", "")
	resp = &SaveTaskForSubmittingDomainDeleteResponse{}
	err = client.DoAction(req, resp)
	return
}

type SaveTaskForSubmittingDomainDeleteResponse struct {
	responses.BaseResponse
	RequestId string
	TaskNo    string
}
