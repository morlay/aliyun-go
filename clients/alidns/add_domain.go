package alidns

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId  common.String
	DomainId   common.String
	DomainName common.String
	PunyCode   common.String
	GroupId    common.String
	GroupName  common.String
	DnsServers AddDomainDnsServerList
}

type AddDomainDnsServerList []common.String

func (list *AddDomainDnsServerList) UnmarshalJSON(data []byte) error {
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
