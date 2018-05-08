package domain_intl

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SaveBatchTaskForUpdatingContactInfoByNewContactRequest struct {
	requests.RpcRequest
	Country                string                                                         `position:"Query" name:"Country"`
	Address                string                                                         `position:"Query" name:"Address"`
	TelArea                string                                                         `position:"Query" name:"TelArea"`
	ContactType            string                                                         `position:"Query" name:"ContactType"`
	City                   string                                                         `position:"Query" name:"City"`
	DomainNames            *SaveBatchTaskForUpdatingContactInfoByNewContactDomainNameList `position:"Query" type:"Repeated" name:"DomainName"`
	Telephone              string                                                         `position:"Query" name:"Telephone"`
	TransferOutProhibited  string                                                         `position:"Query" name:"TransferOutProhibited"`
	RegistrantOrganization string                                                         `position:"Query" name:"RegistrantOrganization"`
	TelExt                 string                                                         `position:"Query" name:"TelExt"`
	Province               string                                                         `position:"Query" name:"Province"`
	PostalCode             string                                                         `position:"Query" name:"PostalCode"`
	UserClientIp           string                                                         `position:"Query" name:"UserClientIp"`
	Lang                   string                                                         `position:"Query" name:"Lang"`
	Email                  string                                                         `position:"Query" name:"Email"`
	RegistrantName         string                                                         `position:"Query" name:"RegistrantName"`
}

func (req *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) Invoke(client *sdk.Client) (resp *SaveBatchTaskForUpdatingContactInfoByNewContactResponse, err error) {
	req.InitWithApiInfo("Domain-intl", "2017-12-18", "SaveBatchTaskForUpdatingContactInfoByNewContact", "domain", "")
	resp = &SaveBatchTaskForUpdatingContactInfoByNewContactResponse{}
	err = client.DoAction(req, resp)
	return
}

type SaveBatchTaskForUpdatingContactInfoByNewContactResponse struct {
	responses.BaseResponse
	RequestId common.String
	TaskNo    common.String
}

type SaveBatchTaskForUpdatingContactInfoByNewContactDomainNameList []string

func (list *SaveBatchTaskForUpdatingContactInfoByNewContactDomainNameList) UnmarshalJSON(data []byte) error {
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
