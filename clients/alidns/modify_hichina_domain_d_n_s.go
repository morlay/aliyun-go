package alidns

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyHichinaDomainDNSRequest struct {
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	DomainName   string `position:"Query" name:"DomainName"`
}

func (r ModifyHichinaDomainDNSRequest) Invoke(client *sdk.Client) (response *ModifyHichinaDomainDNSResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyHichinaDomainDNSRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Alidns", "2015-01-09", "ModifyHichinaDomainDNS", "", "")

	resp := struct {
		*responses.BaseResponse
		ModifyHichinaDomainDNSResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyHichinaDomainDNSResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyHichinaDomainDNSResponse struct {
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
