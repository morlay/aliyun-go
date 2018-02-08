package domain

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest struct {
	requests.RpcRequest
	Country                  string                                                               `position:"Query" name:"Country"`
	IdentityCredentialType   string                                                               `position:"Query" name:"IdentityCredentialType"`
	Address                  string                                                               `position:"Query" name:"Address"`
	TelArea                  string                                                               `position:"Query" name:"TelArea"`
	City                     string                                                               `position:"Query" name:"City"`
	ZhAddress                string                                                               `position:"Query" name:"ZhAddress"`
	RegistrantType           string                                                               `position:"Query" name:"RegistrantType"`
	DomainNames              *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialDomainNameList `position:"Query" type:"Repeated" name:"DomainName"`
	IdentityCredential       string                                                               `position:"Body" name:"IdentityCredential"`
	Telephone                string                                                               `position:"Query" name:"Telephone"`
	TransferOutProhibited    string                                                               `position:"Query" name:"TransferOutProhibited"`
	ZhCity                   string                                                               `position:"Query" name:"ZhCity"`
	ZhProvince               string                                                               `position:"Query" name:"ZhProvince"`
	RegistrantOrganization   string                                                               `position:"Query" name:"RegistrantOrganization"`
	TelExt                   string                                                               `position:"Query" name:"TelExt"`
	Province                 string                                                               `position:"Query" name:"Province"`
	ZhRegistrantName         string                                                               `position:"Query" name:"ZhRegistrantName"`
	PostalCode               string                                                               `position:"Query" name:"PostalCode"`
	UserClientIp             string                                                               `position:"Query" name:"UserClientIp"`
	Lang                     string                                                               `position:"Query" name:"Lang"`
	IdentityCredentialNo     string                                                               `position:"Query" name:"IdentityCredentialNo"`
	Email                    string                                                               `position:"Query" name:"Email"`
	RegistrantName           string                                                               `position:"Query" name:"RegistrantName"`
	ZhRegistrantOrganization string                                                               `position:"Query" name:"ZhRegistrantOrganization"`
}

func (req *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) Invoke(client *sdk.Client) (resp *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponse, err error) {
	req.InitWithApiInfo("Domain", "2018-01-29", "SaveTaskForUpdatingRegistrantInfoByIdentityCredential", "", "")
	resp = &SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponse{}
	err = client.DoAction(req, resp)
	return
}

type SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponse struct {
	responses.BaseResponse
	RequestId string
	TaskNo    string
}

type SaveTaskForUpdatingRegistrantInfoByIdentityCredentialDomainNameList []string

func (list *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialDomainNameList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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
