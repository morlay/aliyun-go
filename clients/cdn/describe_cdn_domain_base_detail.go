package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeCdnDomainBaseDetailRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeCdnDomainBaseDetailRequest) Invoke(client *sdk.Client) (resp *DescribeCdnDomainBaseDetailResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeCdnDomainBaseDetail", "", "")
	resp = &DescribeCdnDomainBaseDetailResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeCdnDomainBaseDetailResponse struct {
	responses.BaseResponse
	RequestId             common.String
	DomainBaseDetailModel DescribeCdnDomainBaseDetailDomainBaseDetailModel
}

type DescribeCdnDomainBaseDetailDomainBaseDetailModel struct {
	Cname        common.String
	CdnType      common.String
	DomainStatus common.String
	SourceType   common.String
	Region       common.String
	DomainName   common.String
	Remark       common.String
	GmtModified  common.String
	GmtCreated   common.String
	Sources      DescribeCdnDomainBaseDetailSourceList
}

type DescribeCdnDomainBaseDetailSourceList []common.String

func (list *DescribeCdnDomainBaseDetailSourceList) UnmarshalJSON(data []byte) error {
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
