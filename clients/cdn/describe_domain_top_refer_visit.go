package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDomainTopReferVisitRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	SortBy        string `position:"Query" name:"SortBy"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeDomainTopReferVisitRequest) Invoke(client *sdk.Client) (response *DescribeDomainTopReferVisitResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeDomainTopReferVisitRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainTopReferVisit", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeDomainTopReferVisitResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeDomainTopReferVisitResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeDomainTopReferVisitResponse struct {
	RequestId    string
	DomainName   string
	StartTime    string
	TopReferList DescribeDomainTopReferVisitReferListList
}

type DescribeDomainTopReferVisitReferList struct {
	ReferDetail     string
	VisitData       string
	VisitProportion float32
	Flow            string
	FlowProportion  float32
}

type DescribeDomainTopReferVisitReferListList []DescribeDomainTopReferVisitReferList

func (list *DescribeDomainTopReferVisitReferListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainTopReferVisitReferList)
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
