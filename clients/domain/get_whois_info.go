package domain

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetWhoisInfoRequest struct {
	DomainName string `position:"Query" name:"DomainName"`
}

func (r GetWhoisInfoRequest) Invoke(client *sdk.Client) (response *GetWhoisInfoResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetWhoisInfoRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Domain", "2016-05-11", "GetWhoisInfo", "", "")

	resp := struct {
		*responses.BaseResponse
		GetWhoisInfoResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetWhoisInfoResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetWhoisInfoResponse struct {
	RequestId                  string
	ReferralURL                string
	DomainName                 string
	Registrar                  string
	RegistrarWhoisServer       string
	StatusList                 string
	FormatCreationDate         string
	FormatExpirationDate       string
	FormatUpdatedDate          string
	NameServerList             string
	UpdatedDate                string
	CreationDate               string
	ExpirationDate             string
	RegistryDomainID           string
	RegistrarURL               string
	RegistrarIANAID            string
	RegistrarAbuseContactEmail string
	RegistrarAbuseContactPhone string
	RegistryRegistrantId       string
	RegistrantName             string
	RegistrantOrganization     string
	RegistrantStreet           string
	RegistrantCity             string
	RegistrantStateProvince    string
	RegistrantPostalCode       string
	RegistrantCountry          string
	RegistrantPhone            string
	RegistrantPhoneExt         string
	RegistrantFax              string
	RegistrantFaxExt           string
	RegistrantEmail            string
	RegistryAdminID            string
	RegistryTechID             string
	AdminName                  string
	AdminOrganization          string
	AdminStreet                string
	AdminCity                  string
	AdminStateProvince         string
	AdminPostalCode            string
	AdminCountry               string
	AdminPhone                 string
	AdminPhoneExt              string
	AdminFax                   string
	AdminFaxExt                string
	AdminEmail                 string
	TechName                   string
	TechOrganization           string
	TechStreet                 string
	TechCity                   string
	TechStateProvince          string
	TechPostalCode             string
	TechCountry                string
	TechPhone                  string
	TechPhoneExt               string
	TechFax                    string
	TechFaxExt                 string
	TechEmail                  string
	OriginalInfo               string
	CacheUpdatedDate           string
	WhoisProtected             string
	DomainStatusList           GetWhoisInfoDomainStatusList
}

type GetWhoisInfoDomainStatus struct {
	Status string
	Desc   string
	Tip    string
}

type GetWhoisInfoDomainStatusList []GetWhoisInfoDomainStatus

func (list *GetWhoisInfoDomainStatusList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetWhoisInfoDomainStatus)
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
