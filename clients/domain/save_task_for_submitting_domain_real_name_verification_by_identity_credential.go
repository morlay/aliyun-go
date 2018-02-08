package domain

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialRequest struct {
	requests.RpcRequest
	IdentityCredentialType string                                                                             `position:"Query" name:"IdentityCredentialType"`
	UserClientIp           string                                                                             `position:"Query" name:"UserClientIp"`
	IdentityCredential     string                                                                             `position:"Body" name:"IdentityCredential"`
	DomainNames            *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialDomainNameList `position:"Query" type:"Repeated" name:"DomainName"`
	Lang                   string                                                                             `position:"Query" name:"Lang"`
	IdentityCredentialNo   string                                                                             `position:"Query" name:"IdentityCredentialNo"`
}

func (req *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialRequest) Invoke(client *sdk.Client) (resp *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialResponse, err error) {
	req.InitWithApiInfo("Domain", "2018-01-29", "SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredential", "", "")
	resp = &SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialResponse{}
	err = client.DoAction(req, resp)
	return
}

type SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialResponse struct {
	responses.BaseResponse
	RequestId string
	TaskNo    string
}

type SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialDomainNameList []string

func (list *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialDomainNameList) UnmarshalJSON(data []byte) error {
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
