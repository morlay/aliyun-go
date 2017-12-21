package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeCdnDomainBaseDetailRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeCdnDomainBaseDetailRequest) Invoke(client *sdk.Client) (response *DescribeCdnDomainBaseDetailResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeCdnDomainBaseDetailRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeCdnDomainBaseDetail", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeCdnDomainBaseDetailResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeCdnDomainBaseDetailResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeCdnDomainBaseDetailResponse struct {
	RequestId             string
	DomainBaseDetailModel DescribeCdnDomainBaseDetailDomainBaseDetailModel
}

type DescribeCdnDomainBaseDetailDomainBaseDetailModel struct {
	Cname        string
	CdnType      string
	DomainStatus string
	SourceType   string
	Region       string
	DomainName   string
	Remark       string
	GmtModified  string
	GmtCreated   string
	Sources      DescribeCdnDomainBaseDetailSourceList
}

type DescribeCdnDomainBaseDetailSourceList []string

func (list *DescribeCdnDomainBaseDetailSourceList) UnmarshalJSON(data []byte) error {
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
