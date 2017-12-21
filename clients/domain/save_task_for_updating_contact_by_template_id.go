package domain

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SaveTaskForUpdatingContactByTemplateIdRequest struct {
	SaleId            string `position:"Query" name:"SaleId"`
	ContactType       string `position:"Query" name:"ContactType"`
	UserClientIp      string `position:"Query" name:"UserClientIp"`
	DomainName        string `position:"Query" name:"DomainName"`
	AddTransferLock   string `position:"Query" name:"AddTransferLock"`
	Lang              string `position:"Query" name:"Lang"`
	ContactTemplateId int64  `position:"Query" name:"ContactTemplateId"`
}

func (r SaveTaskForUpdatingContactByTemplateIdRequest) Invoke(client *sdk.Client) (response *SaveTaskForUpdatingContactByTemplateIdResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SaveTaskForUpdatingContactByTemplateIdRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Domain", "2016-05-11", "SaveTaskForUpdatingContactByTemplateId", "", "")

	resp := struct {
		*responses.BaseResponse
		SaveTaskForUpdatingContactByTemplateIdResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SaveTaskForUpdatingContactByTemplateIdResponse

	err = client.DoAction(&req, &resp)
	return
}

type SaveTaskForUpdatingContactByTemplateIdResponse struct {
	RequestId string
	Success   bool
	TaskNo    string
}
