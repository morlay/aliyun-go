package dcdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDcdnDomainDetailRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeDcdnDomainDetailRequest) Invoke(client *sdk.Client) (resp *DescribeDcdnDomainDetailResponse, err error) {
	req.InitWithApiInfo("dcdn", "2018-01-15", "DescribeDcdnDomainDetail", "dcdn", "")
	resp = &DescribeDcdnDomainDetailResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDcdnDomainDetailResponse struct {
	responses.BaseResponse
	RequestId    string
	DomainDetail DescribeDcdnDomainDetailDomainDetail
}

type DescribeDcdnDomainDetailDomainDetail struct {
	GmtCreated      string
	GmtModified     string
	DomainStatus    string
	Cname           string
	DomainName      string
	Description     string
	SSLProtocol     string
	SSLPub          string
	Scope           string
	CertName        string
	ResourceGroupId string
	Sources         DescribeDcdnDomainDetailSourceList
}

type DescribeDcdnDomainDetailSource struct {
	Content  string
	Type     string
	Port     int
	Enabled  string
	Priority string
}

type DescribeDcdnDomainDetailSourceList []DescribeDcdnDomainDetailSource

func (list *DescribeDcdnDomainDetailSourceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDcdnDomainDetailSource)
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
