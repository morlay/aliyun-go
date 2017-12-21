package alidns

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDomainWhoisInfoRequest struct {
	requests.RpcRequest
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	DomainName   string `position:"Query" name:"DomainName"`
	GroupId      string `position:"Query" name:"GroupId"`
}

func (req *DescribeDomainWhoisInfoRequest) Invoke(client *sdk.Client) (resp *DescribeDomainWhoisInfoResponse, err error) {
	req.InitWithApiInfo("Alidns", "2015-01-09", "DescribeDomainWhoisInfo", "", "")
	resp = &DescribeDomainWhoisInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDomainWhoisInfoResponse struct {
	responses.BaseResponse
	RequestId        string
	RegistrantName   string
	RegistrantEmail  string
	Registrar        string
	RegistrationDate string
	ExpirationDate   string
	StatusList       DescribeDomainWhoisInfoStatusListList
	DnsServers       DescribeDomainWhoisInfoDnsServerList
}

type DescribeDomainWhoisInfoStatusListList []string

func (list *DescribeDomainWhoisInfoStatusListList) UnmarshalJSON(data []byte) error {
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

type DescribeDomainWhoisInfoDnsServerList []string

func (list *DescribeDomainWhoisInfoDnsServerList) UnmarshalJSON(data []byte) error {
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
