package domain

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SaveBatchTaskForCreatingOrderActivateRequest struct {
	requests.RpcRequest
	OrderActivateParams *SaveBatchTaskForCreatingOrderActivateOrderActivateParamList `position:"Query" type:"Repeated" name:"OrderActivateParam"`
	UserClientIp        string                                                       `position:"Query" name:"UserClientIp"`
	Lang                string                                                       `position:"Query" name:"Lang"`
}

func (req *SaveBatchTaskForCreatingOrderActivateRequest) Invoke(client *sdk.Client) (resp *SaveBatchTaskForCreatingOrderActivateResponse, err error) {
	req.InitWithApiInfo("Domain", "2018-01-29", "SaveBatchTaskForCreatingOrderActivate", "", "")
	resp = &SaveBatchTaskForCreatingOrderActivateResponse{}
	err = client.DoAction(req, resp)
	return
}

type SaveBatchTaskForCreatingOrderActivateOrderActivateParam struct {
	DomainName               string `name:"DomainName"`
	SubscriptionDuration     int    `name:"SubscriptionDuration"`
	RegistrantProfileId      int64  `name:"RegistrantProfileId"`
	EnableDomainProxy        string `name:"EnableDomainProxy"`
	PermitPremiumActivation  string `name:"PermitPremiumActivation"`
	AliyunDns                string `name:"AliyunDns"`
	Dns1                     string `name:"Dns.1"`
	Dns2                     string `name:"Dns.2"`
	ZhCity                   string `name:"ZhCity"`
	ZhRegistrantOrganization string `name:"ZhRegistrantOrganization"`
	Country                  string `name:"Country"`
	ZhRegistrantName         string `name:"ZhRegistrantName"`
	ZhProvince               string `name:"ZhProvince"`
	ZhAddress                string `name:"ZhAddress"`
	City                     string `name:"City"`
	RegistrantOrganization   string `name:"RegistrantOrganization"`
	RegistrantName           string `name:"RegistrantName"`
	Province                 string `name:"Province"`
	Address                  string `name:"Address"`
	Email                    string `name:"Email"`
	PostalCode               string `name:"PostalCode"`
	TelArea                  string `name:"TelArea"`
	Telephone                string `name:"Telephone"`
	TelExt                   string `name:"TelExt"`
	RegistrantType           string `name:"RegistrantType"`
}

type SaveBatchTaskForCreatingOrderActivateResponse struct {
	responses.BaseResponse
	RequestId common.String
	TaskNo    common.String
}

type SaveBatchTaskForCreatingOrderActivateOrderActivateParamList []SaveBatchTaskForCreatingOrderActivateOrderActivateParam

func (list *SaveBatchTaskForCreatingOrderActivateOrderActivateParamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SaveBatchTaskForCreatingOrderActivateOrderActivateParam)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
