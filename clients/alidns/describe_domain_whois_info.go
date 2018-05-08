package alidns

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId        common.String
	RegistrantName   common.String
	RegistrantEmail  common.String
	Registrar        common.String
	RegistrationDate common.String
	ExpirationDate   common.String
	StatusList       DescribeDomainWhoisInfoStatusListList
	DnsServers       DescribeDomainWhoisInfoDnsServerList
}

type DescribeDomainWhoisInfoStatusListList []common.String

func (list *DescribeDomainWhoisInfoStatusListList) UnmarshalJSON(data []byte) error {
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

type DescribeDomainWhoisInfoDnsServerList []common.String

func (list *DescribeDomainWhoisInfoDnsServerList) UnmarshalJSON(data []byte) error {
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
