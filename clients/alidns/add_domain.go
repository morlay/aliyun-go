package alidns

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddDomainRequest struct {
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	DomainName   string `position:"Query" name:"DomainName"`
	GroupId      string `position:"Query" name:"GroupId"`
}

func (r AddDomainRequest) Invoke(client *sdk.Client) (response *AddDomainResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AddDomainRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Alidns", "2015-01-09", "AddDomain", "", "")

	resp := struct {
		*responses.BaseResponse
		AddDomainResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AddDomainResponse

	err = client.DoAction(&req, &resp)
	return
}

type AddDomainResponse struct {
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
