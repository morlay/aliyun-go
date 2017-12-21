package domain

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SaveTaskForUpdatingContactByTemplateIdRequest struct {
	requests.RpcRequest
	SaleId            string `position:"Query" name:"SaleId"`
	ContactType       string `position:"Query" name:"ContactType"`
	UserClientIp      string `position:"Query" name:"UserClientIp"`
	DomainName        string `position:"Query" name:"DomainName"`
	AddTransferLock   string `position:"Query" name:"AddTransferLock"`
	Lang              string `position:"Query" name:"Lang"`
	ContactTemplateId int64  `position:"Query" name:"ContactTemplateId"`
}

func (req *SaveTaskForUpdatingContactByTemplateIdRequest) Invoke(client *sdk.Client) (resp *SaveTaskForUpdatingContactByTemplateIdResponse, err error) {
	req.InitWithApiInfo("Domain", "2016-05-11", "SaveTaskForUpdatingContactByTemplateId", "", "")
	resp = &SaveTaskForUpdatingContactByTemplateIdResponse{}
	err = client.DoAction(req, resp)
	return
}

type SaveTaskForUpdatingContactByTemplateIdResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	TaskNo    string
}
