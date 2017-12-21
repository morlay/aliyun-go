package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeCdnDomainDetailRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeCdnDomainDetailRequest) Invoke(client *sdk.Client) (resp *DescribeCdnDomainDetailResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeCdnDomainDetail", "", "")
	resp = &DescribeCdnDomainDetailResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeCdnDomainDetailResponse struct {
	responses.BaseResponse
	RequestId            string
	GetDomainDetailModel DescribeCdnDomainDetailGetDomainDetailModel
}

type DescribeCdnDomainDetailGetDomainDetailModel struct {
	GmtCreated              string
	GmtModified             string
	SourceType              string
	DomainStatus            string
	SourcePort              int
	CdnType                 string
	Cname                   string
	HttpsCname              string
	DomainName              string
	Description             string
	ServerCertificateStatus string
	ServerCertificate       string
	Region                  string
	Scope                   string
	CertificateName         string
	ResourceGroupId         string
	SourceModels            DescribeCdnDomainDetailSourceModelList
	Sources                 DescribeCdnDomainDetailSourceList
}

type DescribeCdnDomainDetailSourceModel struct {
	Content  string
	Type     string
	Port     int
	Enabled  string
	Priority string
}

type DescribeCdnDomainDetailSourceModelList []DescribeCdnDomainDetailSourceModel

func (list *DescribeCdnDomainDetailSourceModelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCdnDomainDetailSourceModel)
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

type DescribeCdnDomainDetailSourceList []string

func (list *DescribeCdnDomainDetailSourceList) UnmarshalJSON(data []byte) error {
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
