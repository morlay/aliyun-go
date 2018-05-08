package domain_intl

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SaveSingleTaskForCreatingOrderActivateRequest struct {
	requests.RpcRequest
	Country                 string `position:"Query" name:"Country"`
	SubscriptionDuration    int    `position:"Query" name:"SubscriptionDuration"`
	Address                 string `position:"Query" name:"Address"`
	PermitPremiumActivation string `position:"Query" name:"PermitPremiumActivation"`
	TelArea                 string `position:"Query" name:"TelArea"`
	City                    string `position:"Query" name:"City"`
	Dns2                    string `position:"Query" name:"Dns.2"`
	Dns1                    string `position:"Query" name:"Dns.1"`
	DomainName              string `position:"Query" name:"DomainName"`
	RegistrantProfileId     int64  `position:"Query" name:"RegistrantProfileId"`
	Telephone               string `position:"Query" name:"Telephone"`
	AliyunDns               string `position:"Query" name:"AliyunDns"`
	RegistrantOrganization  string `position:"Query" name:"RegistrantOrganization"`
	TelExt                  string `position:"Query" name:"TelExt"`
	Province                string `position:"Query" name:"Province"`
	PostalCode              string `position:"Query" name:"PostalCode"`
	UserClientIp            string `position:"Query" name:"UserClientIp"`
	EnableDomainProxy       string `position:"Query" name:"EnableDomainProxy"`
	Lang                    string `position:"Query" name:"Lang"`
	Email                   string `position:"Query" name:"Email"`
	RegistrantName          string `position:"Query" name:"RegistrantName"`
}

func (req *SaveSingleTaskForCreatingOrderActivateRequest) Invoke(client *sdk.Client) (resp *SaveSingleTaskForCreatingOrderActivateResponse, err error) {
	req.InitWithApiInfo("Domain-intl", "2017-12-18", "SaveSingleTaskForCreatingOrderActivate", "domain", "")
	resp = &SaveSingleTaskForCreatingOrderActivateResponse{}
	err = client.DoAction(req, resp)
	return
}

type SaveSingleTaskForCreatingOrderActivateResponse struct {
	responses.BaseResponse
	RequestId common.String
	TaskNo    common.String
}
