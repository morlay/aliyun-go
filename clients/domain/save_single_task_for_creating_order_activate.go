package domain

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SaveSingleTaskForCreatingOrderActivateRequest struct {
	requests.RpcRequest
	Country                  string `position:"Query" name:"Country"`
	SubscriptionDuration     int    `position:"Query" name:"SubscriptionDuration"`
	PermitPremiumActivation  string `position:"Query" name:"PermitPremiumActivation"`
	City                     string `position:"Query" name:"City"`
	Dns2                     string `position:"Query" name:"Dns.2"`
	Dns1                     string `position:"Query" name:"Dns.1"`
	RegistrantProfileId      int64  `position:"Query" name:"RegistrantProfileId"`
	AliyunDns                string `position:"Query" name:"AliyunDns"`
	ZhCity                   string `position:"Query" name:"ZhCity"`
	TelExt                   string `position:"Query" name:"TelExt"`
	ZhRegistrantName         string `position:"Query" name:"ZhRegistrantName"`
	Province                 string `position:"Query" name:"Province"`
	PostalCode               string `position:"Query" name:"PostalCode"`
	Lang                     string `position:"Query" name:"Lang"`
	Email                    string `position:"Query" name:"Email"`
	ZhRegistrantOrganization string `position:"Query" name:"ZhRegistrantOrganization"`
	Address                  string `position:"Query" name:"Address"`
	TelArea                  string `position:"Query" name:"TelArea"`
	DomainName               string `position:"Query" name:"DomainName"`
	ZhAddress                string `position:"Query" name:"ZhAddress"`
	RegistrantType           string `position:"Query" name:"RegistrantType"`
	Telephone                string `position:"Query" name:"Telephone"`
	ZhProvince               string `position:"Query" name:"ZhProvince"`
	RegistrantOrganization   string `position:"Query" name:"RegistrantOrganization"`
	EnableDomainProxy        string `position:"Query" name:"EnableDomainProxy"`
	UserClientIp             string `position:"Query" name:"UserClientIp"`
	RegistrantName           string `position:"Query" name:"RegistrantName"`
}

func (req *SaveSingleTaskForCreatingOrderActivateRequest) Invoke(client *sdk.Client) (resp *SaveSingleTaskForCreatingOrderActivateResponse, err error) {
	req.InitWithApiInfo("Domain", "2018-01-29", "SaveSingleTaskForCreatingOrderActivate", "", "")
	resp = &SaveSingleTaskForCreatingOrderActivateResponse{}
	err = client.DoAction(req, resp)
	return
}

type SaveSingleTaskForCreatingOrderActivateResponse struct {
	responses.BaseResponse
	RequestId common.String
	TaskNo    common.String
}
