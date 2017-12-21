package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDomainsBySourceRequest struct {
	Sources       string `position:"Query" name:"Sources"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeDomainsBySourceRequest) Invoke(client *sdk.Client) (response *DescribeDomainsBySourceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeDomainsBySourceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainsBySource", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeDomainsBySourceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeDomainsBySourceResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeDomainsBySourceResponse struct {
	RequestId   string
	Sources     string
	DomainsList DescribeDomainsBySourceDomainsDataList
}

type DescribeDomainsBySourceDomainsData struct {
	Source      string
	DomainInfos DescribeDomainsBySourceDomainInfoList
	Domains     DescribeDomainsBySourceDomainList
}

type DescribeDomainsBySourceDomainInfo struct {
	DomainName  string
	DomainCname string
	CreateTime  string
	UpdateTime  string
	Status      string
}

type DescribeDomainsBySourceDomainsDataList []DescribeDomainsBySourceDomainsData

func (list *DescribeDomainsBySourceDomainsDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainsBySourceDomainsData)
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

type DescribeDomainsBySourceDomainInfoList []DescribeDomainsBySourceDomainInfo

func (list *DescribeDomainsBySourceDomainInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainsBySourceDomainInfo)
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

type DescribeDomainsBySourceDomainList []string

func (list *DescribeDomainsBySourceDomainList) UnmarshalJSON(data []byte) error {
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
