package dcdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId    common.String
	DomainDetail DescribeDcdnDomainDetailDomainDetail
}

type DescribeDcdnDomainDetailDomainDetail struct {
	GmtCreated      common.String
	GmtModified     common.String
	DomainStatus    common.String
	Cname           common.String
	DomainName      common.String
	Description     common.String
	SSLProtocol     common.String
	SSLPub          common.String
	Scope           common.String
	CertName        common.String
	ResourceGroupId common.String
	Sources         DescribeDcdnDomainDetailSourceList
}

type DescribeDcdnDomainDetailSource struct {
	Content  common.String
	Type     common.String
	Port     common.Integer
	Enabled  common.String
	Priority common.String
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
