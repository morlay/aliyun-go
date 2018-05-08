package domain

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	req.InitWithApiInfo("Domain", "2018-01-29", "SaveSingleTaskForUpdatingContactInfo", "", "")
	resp = &SaveSingleTaskForUpdatingContactInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type SaveSingleTaskForUpdatingContactInfoResponse struct {
	responses.BaseResponse
	RequestId common.String
	TaskNo    common.String
}
