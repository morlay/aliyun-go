package alidns

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeDomainNsRequest struct {
	requests.RpcRequest
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	DomainName   string `position:"Query" name:"DomainName"`
}

func (req *DescribeDomainNsRequest) Invoke(client *sdk.Client) (resp *DescribeDomainNsResponse, err error) {
	req.InitWithApiInfo("Alidns", "2015-01-09", "DescribeDomainNs", "", "")
	resp = &DescribeDomainNsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDomainNsResponse struct {
	responses.BaseResponse
	RequestId     common.String
	AllAliDns     bool
	IncludeAliDns bool
	DnsServers    DescribeDomainNsDnsServerList
}

type DescribeDomainNsDnsServerList []common.String

func (list *DescribeDomainNsDnsServerList) UnmarshalJSON(data []byte) error {
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
