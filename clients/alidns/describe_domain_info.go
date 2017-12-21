package alidns

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDomainInfoRequest struct {
	requests.RpcRequest
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	DomainName   string `position:"Query" name:"DomainName"`
}

func (req *DescribeDomainInfoRequest) Invoke(client *sdk.Client) (resp *DescribeDomainInfoResponse, err error) {
	req.InitWithApiInfo("Alidns", "2015-01-09", "DescribeDomainInfo", "", "")
	resp = &DescribeDomainInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDomainInfoResponse struct {
	responses.BaseResponse
	RequestId       string
	DomainId        string
	DomainName      string
	PunyCode        string
	AliDomain       bool
	RegistrantEmail string
	GroupId         string
	GroupName       string
	InstanceId      string
	VersionCode     string
	VersionName     string
	DnsServers      DescribeDomainInfoDnsServerList
}

type DescribeDomainInfoDnsServerList []string

func (list *DescribeDomainInfoDnsServerList) UnmarshalJSON(data []byte) error {
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
