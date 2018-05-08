package httpdns

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListDomainsRequest struct {
	requests.RpcRequest
	PageSize   int `position:"Query" name:"PageSize"`
	PageNumber int `position:"Query" name:"PageNumber"`
}

func (req *ListDomainsRequest) Invoke(client *sdk.Client) (resp *ListDomainsResponse, err error) {
	req.InitWithApiInfo("Httpdns", "2016-02-01", "ListDomains", "", "")
	resp = &ListDomainsResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListDomainsResponse struct {
	responses.BaseResponse
	RequestId   common.String
	TotalCount  common.Long
	PageNumber  common.Long
	PageSize    common.Long
	DomainInfos ListDomainsDomainInfoList
}

type ListDomainsDomainInfo struct {
	DomainName    common.String
	Resolved      common.Long
	ResolvedHttps common.Long
}

type ListDomainsDomainInfoList []ListDomainsDomainInfo

func (list *ListDomainsDomainInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListDomainsDomainInfo)
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
