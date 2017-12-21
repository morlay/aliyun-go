package httpdns

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListDomainsRequest struct {
	PageSize   int `position:"Query" name:"PageSize"`
	PageNumber int `position:"Query" name:"PageNumber"`
}

func (r ListDomainsRequest) Invoke(client *sdk.Client) (response *ListDomainsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ListDomainsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Httpdns", "2016-02-01", "ListDomains", "", "")

	resp := struct {
		*responses.BaseResponse
		ListDomainsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ListDomainsResponse

	err = client.DoAction(&req, &resp)
	return
}

type ListDomainsResponse struct {
	RequestId   string
	TotalCount  int64
	PageNumber  int64
	PageSize    int64
	DomainInfos ListDomainsDomainInfoList
}

type ListDomainsDomainInfo struct {
	DomainName    string
	Resolved      int64
	ResolvedHttps int64
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
