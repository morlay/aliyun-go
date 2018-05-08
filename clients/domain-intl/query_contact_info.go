package domain_intl

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId              common.String
	CreateDate             common.String
	RegistrantName         common.String
	RegistrantOrganization common.String
	Country                common.String
	Province               common.String
	City                   common.String
	Address                common.String
	Email                  common.String
	PostalCode             common.String
	TelArea                common.String
	Telephone              common.String
	TelExt                 common.String
}
