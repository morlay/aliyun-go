package domain_intl

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryContactInfoRequest struct {
	requests.RpcRequest
	ContactType  string `position:"Query" name:"ContactType"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	DomainName   string `position:"Query" name:"DomainName"`
	Lang         string `position:"Query" name:"Lang"`
}

func (req *QueryContactInfoRequest) Invoke(client *sdk.Client) (resp *QueryContactInfoResponse, err error) {
	req.InitWithApiInfo("Domain-intl", "2017-12-18", "QueryContactInfo", "domain", "")
	resp = &QueryContactInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryContactInfoResponse struct {
	responses.BaseResponse
	RequestId              string
	CreateDate             string
	RegistrantName         string
	RegistrantOrganization string
	Country                string
	Province               string
	City                   string
	Address                string
	Email                  string
	PostalCode             string
	TelArea                string
	Telephone              string
	TelExt                 string
}
