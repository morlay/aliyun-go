package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeDomainsBySourceRequest struct {
	requests.RpcRequest
	Sources       string `position:"Query" name:"Sources"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeDomainsBySourceRequest) Invoke(client *sdk.Client) (resp *DescribeDomainsBySourceResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainsBySource", "", "")
	resp = &DescribeDomainsBySourceResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDomainsBySourceResponse struct {
	responses.BaseResponse
	RequestId   common.String
	Sources     common.String
	DomainsList DescribeDomainsBySourceDomainsDataList
}

type DescribeDomainsBySourceDomainsData struct {
	Source      common.String
	DomainInfos DescribeDomainsBySourceDomainInfoList
	Domains     DescribeDomainsBySourceDomainList
}

type DescribeDomainsBySourceDomainInfo struct {
	DomainName  common.String
	DomainCname common.String
	CreateTime  common.String
	UpdateTime  common.String
	Status      common.String
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

type DescribeDomainsBySourceDomainList []common.String

func (list *DescribeDomainsBySourceDomainList) UnmarshalJSON(data []byte) error {
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
