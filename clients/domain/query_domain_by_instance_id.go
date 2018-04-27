package domain

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryDomainByInstanceIdRequest struct {
	requests.RpcRequest
	InstanceId   string `position:"Query" name:"InstanceId"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
}

func (req *QueryDomainByInstanceIdRequest) Invoke(client *sdk.Client) (resp *QueryDomainByInstanceIdResponse, err error) {
	req.InitWithApiInfo("Domain", "2018-01-29", "QueryDomainByInstanceId", "", "")
	resp = &QueryDomainByInstanceIdResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryDomainByInstanceIdResponse struct {
	responses.BaseResponse
	UserId                       string
	DomainName                   string
	InstanceId                   string
	RegistrationDate             string
	ExpirationDate               string
	RegistrantOrganization       string
	RegistrantName               string
	Email                        string
	UpdateProhibitionLock        string
	TransferProhibitionLock      string
	DomainNameProxyService       bool
	Premium                      bool
	EmailVerificationStatus      int
	EmailVerificationClientHold  bool
	RealNameStatus               string
	RegistrantUpdatingStatus     string
	TransferOutStatus            string
	RegistrantType               string
	DomainNameVerificationStatus string
	RequestId                    string
	DnsList                      QueryDomainByInstanceIdDnsListList
}

type QueryDomainByInstanceIdDnsListList []string

func (list *QueryDomainByInstanceIdDnsListList) UnmarshalJSON(data []byte) error {
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
