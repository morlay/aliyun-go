package alidns

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyHichinaDomainDNSRequest struct {
	requests.RpcRequest
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	DomainName   string `position:"Query" name:"DomainName"`
}

func (req *ModifyHichinaDomainDNSRequest) Invoke(client *sdk.Client) (resp *ModifyHichinaDomainDNSResponse, err error) {
	req.InitWithApiInfo("Alidns", "2015-01-09", "ModifyHichinaDomainDNS", "", "")
	resp = &ModifyHichinaDomainDNSResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyHichinaDomainDNSResponse struct {
	responses.BaseResponse
	RequestId          string
	OriginalDnsServers ModifyHichinaDomainDNSOriginalDnsServerList
	NewDnsServers      ModifyHichinaDomainDNSNewDnsServerList
}

type ModifyHichinaDomainDNSOriginalDnsServerList []string

func (list *ModifyHichinaDomainDNSOriginalDnsServerList) UnmarshalJSON(data []byte) error {
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

type ModifyHichinaDomainDNSNewDnsServerList []string

func (list *ModifyHichinaDomainDNSNewDnsServerList) UnmarshalJSON(data []byte) error {
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
