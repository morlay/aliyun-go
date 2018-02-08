package domain

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
	req.InitWithApiInfo("Domain", "2018-01-29", "SaveTaskForSubmittingDomainDelete", "", "")
	resp = &SaveTaskForSubmittingDomainDeleteResponse{}
	err = client.DoAction(req, resp)
	return
}

type SaveTaskForSubmittingDomainDeleteResponse struct {
	responses.BaseResponse
	RequestId string
	TaskNo    string
}
