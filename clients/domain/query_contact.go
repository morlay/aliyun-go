package domain

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryContactRequest struct {
	requests.RpcRequest
	ContactType  string `position:"Query" name:"ContactType"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	DomainName   string `position:"Query" name:"DomainName"`
	Lang         string `position:"Query" name:"Lang"`
}

func (req *QueryContactRequest) Invoke(client *sdk.Client) (resp *QueryContactResponse, err error) {
	req.InitWithApiInfo("Domain", "2016-05-11", "QueryContact", "", "")
	resp = &QueryContactResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryContactResponse struct {
	responses.BaseResponse
	RequestId  string
	CreateDate string
	UpdateDate string
	CName      string
	EName      string
	CCompany   string
	ECompany   string
	CCountry   string
	CProvince  string
	EProvince  string
	CCity      string
	ECity      string
	CVenu      string
	EVenu      string
	Email      string
	TelArea    string
	PostalCode string
	TelMain    string
	TelExt     string
	RegType    string
}
