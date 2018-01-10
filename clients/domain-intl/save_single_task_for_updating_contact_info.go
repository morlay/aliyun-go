package domain_intl

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SaveSingleTaskForUpdatingContactInfoRequest struct {
	requests.RpcRequest
	InstanceId          string `position:"Query" name:"InstanceId"`
	ContactType         string `position:"Query" name:"ContactType"`
	UserClientIp        string `position:"Query" name:"UserClientIp"`
	DomainName          string `position:"Query" name:"DomainName"`
	RegistrantProfileId int64  `position:"Query" name:"RegistrantProfileId"`
	AddTransferLock     string `position:"Query" name:"AddTransferLock"`
	Lang                string `position:"Query" name:"Lang"`
}

func (req *SaveSingleTaskForUpdatingContactInfoRequest) Invoke(client *sdk.Client) (resp *SaveSingleTaskForUpdatingContactInfoResponse, err error) {
	req.InitWithApiInfo("Domain-intl", "2017-12-18", "SaveSingleTaskForUpdatingContactInfo", "domain", "")
	resp = &SaveSingleTaskForUpdatingContactInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type SaveSingleTaskForUpdatingContactInfoResponse struct {
	responses.BaseResponse
	RequestId string
	TaskNo    string
}
