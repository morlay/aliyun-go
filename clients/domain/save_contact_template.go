package domain

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SaveContactTemplateRequest struct {
	CCompany          string `position:"Query" name:"CCompany"`
	DefaultTemplate   string `position:"Query" name:"DefaultTemplate"`
	TelArea           string `position:"Query" name:"TelArea"`
	ECompany          string `position:"Query" name:"ECompany"`
	TelMain           string `position:"Query" name:"TelMain"`
	CName             string `position:"Query" name:"CName"`
	CProvince         string `position:"Query" name:"CProvince"`
	ECity             string `position:"Query" name:"ECity"`
	CCity             string `position:"Query" name:"CCity"`
	RegType           string `position:"Query" name:"RegType"`
	EName             string `position:"Query" name:"EName"`
	TelExt            string `position:"Query" name:"TelExt"`
	CVenu             string `position:"Query" name:"CVenu"`
	EProvince         string `position:"Query" name:"EProvince"`
	PostalCode        string `position:"Query" name:"PostalCode"`
	UserClientIp      string `position:"Query" name:"UserClientIp"`
	CCountry          string `position:"Query" name:"CCountry"`
	Lang              string `position:"Query" name:"Lang"`
	EVenu             string `position:"Query" name:"EVenu"`
	Email             string `position:"Query" name:"Email"`
	ContactTemplateId int64  `position:"Query" name:"ContactTemplateId"`
}

func (r SaveContactTemplateRequest) Invoke(client *sdk.Client) (response *SaveContactTemplateResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SaveContactTemplateRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Domain", "2016-05-11", "SaveContactTemplate", "", "")

	resp := struct {
		*responses.BaseResponse
		SaveContactTemplateResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SaveContactTemplateResponse

	err = client.DoAction(&req, &resp)
	return
}

type SaveContactTemplateResponse struct {
	RequestId         string
	Success           bool
	ContactTemplateId int64
}
