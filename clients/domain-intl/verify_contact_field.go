package domain_intl

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type VerifyContactFieldRequest struct {
	requests.RpcRequest
	Country                string `position:"Query" name:"Country"`
	Address                string `position:"Query" name:"Address"`
	TelArea                string `position:"Query" name:"TelArea"`
	City                   string `position:"Query" name:"City"`
	Telephone              string `position:"Query" name:"Telephone"`
	RegistrantOrganization string `position:"Query" name:"RegistrantOrganization"`
	TelExt                 string `position:"Query" name:"TelExt"`
	Province               string `position:"Query" name:"Province"`
	PostalCode             string `position:"Query" name:"PostalCode"`
	UserClientIp           string `position:"Query" name:"UserClientIp"`
	Lang                   string `position:"Query" name:"Lang"`
	Email                  string `position:"Query" name:"Email"`
	RegistrantName         string `position:"Query" name:"RegistrantName"`
}

func (req *VerifyContactFieldRequest) Invoke(client *sdk.Client) (resp *VerifyContactFieldResponse, err error) {
	req.InitWithApiInfo("Domain-intl", "2017-12-18", "VerifyContactField", "domain", "")
	resp = &VerifyContactFieldResponse{}
	err = client.DoAction(req, resp)
	return
}

type VerifyContactFieldResponse struct {
	responses.BaseResponse
	RequestId common.String
}
