package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDomainUpstreamOfCenterRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeDomainUpstreamOfCenterRequest) Invoke(client *sdk.Client) (response *DescribeDomainUpstreamOfCenterResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeDomainUpstreamOfCenterRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainUpstreamOfCenter", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeDomainUpstreamOfCenterResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeDomainUpstreamOfCenterResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeDomainUpstreamOfCenterResponse struct {
	RequestId string
	BpsDatas  DescribeDomainUpstreamOfCenterDomainBpsModelList
}

type DescribeDomainUpstreamOfCenterDomainBpsModel struct {
	Time string
	Bps  float32
}

type DescribeDomainUpstreamOfCenterDomainBpsModelList []DescribeDomainUpstreamOfCenterDomainBpsModel

func (list *DescribeDomainUpstreamOfCenterDomainBpsModelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainUpstreamOfCenterDomainBpsModel)
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
