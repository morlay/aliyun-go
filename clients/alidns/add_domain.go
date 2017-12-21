package alidns

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddDomainRequest struct {
	requests.RpcRequest
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	DomainName   string `position:"Query" name:"DomainName"`
	GroupId      string `position:"Query" name:"GroupId"`
}

func (req *AddDomainRequest) Invoke(client *sdk.Client) (resp *AddDomainResponse, err error) {
	req.InitWithApiInfo("Alidns", "2015-01-09", "AddDomain", "", "")
	resp = &AddDomainResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddDomainResponse struct {
	responses.BaseResponse
	RequestId  string
	DomainId   string
	DomainName string
	PunyCode   string
	GroupId    string
	GroupName  string
	DnsServers AddDomainDnsServerList
}

type AddDomainDnsServerList []string

func (list *AddDomainDnsServerList) UnmarshalJSON(data []byte) error {
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
