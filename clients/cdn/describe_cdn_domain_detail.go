package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId            common.String
	GetDomainDetailModel DescribeCdnDomainDetailGetDomainDetailModel
}

type DescribeCdnDomainDetailGetDomainDetailModel struct {
	GmtCreated              common.String
	GmtModified             common.String
	SourceType              common.String
	DomainStatus            common.String
	SourcePort              common.Integer
	CdnType                 common.String
	Cname                   common.String
	HttpsCname              common.String
	DomainName              common.String
	Description             common.String
	ServerCertificateStatus common.String
	ServerCertificate       common.String
	Region                  common.String
	Scope                   common.String
	CertificateName         common.String
	ResourceGroupId         common.String
	SourceModels            DescribeCdnDomainDetailSourceModelList
	Sources                 DescribeCdnDomainDetailSourceList
}

type DescribeCdnDomainDetailSourceModel struct {
	Content  common.String
	Type     common.String
	Port     common.Integer
	Enabled  common.String
	Priority common.String
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

type DescribeCdnDomainDetailSourceList []common.String

func (list *DescribeCdnDomainDetailSourceList) UnmarshalJSON(data []byte) error {
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
