package domain_intl

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type QueryDomainByInstanceIdRequest struct {
	requests.RpcRequest
	InstanceId   string `position:"Query" name:"InstanceId"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
}

func (req *QueryDomainByInstanceIdRequest) Invoke(client *sdk.Client) (resp *QueryDomainByInstanceIdResponse, err error) {
	req.InitWithApiInfo("Domain-intl", "2017-12-18", "QueryDomainByInstanceId", "domain", "")
	resp = &QueryDomainByInstanceIdResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryDomainByInstanceIdResponse struct {
	responses.BaseResponse
	UserId                      common.String
	DomainName                  common.String
	InstanceId                  common.String
	RegistrationDate            common.String
	ExpirationDate              common.String
	RegistrantOrganization      common.String
	RegistrantName              common.String
	Email                       common.String
	UpdateProhibitionLock       common.String
	TransferProhibitionLock     common.String
	DomainNameProxyService      bool
	Premium                     bool
	EmailVerificationStatus     common.Integer
	EmailVerificationClientHold bool
	DnsList                     QueryDomainByInstanceIdDnsListList
}

type QueryDomainByInstanceIdDnsListList []common.String

func (list *QueryDomainByInstanceIdDnsListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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
