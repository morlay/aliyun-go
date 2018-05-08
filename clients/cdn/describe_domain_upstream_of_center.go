package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeDomainUpstreamOfCenterRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeDomainUpstreamOfCenterRequest) Invoke(client *sdk.Client) (resp *DescribeDomainUpstreamOfCenterResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainUpstreamOfCenter", "", "")
	resp = &DescribeDomainUpstreamOfCenterResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDomainUpstreamOfCenterResponse struct {
	responses.BaseResponse
	RequestId common.String
	BpsDatas  DescribeDomainUpstreamOfCenterDomainBpsModelList
}

type DescribeDomainUpstreamOfCenterDomainBpsModel struct {
	Time common.String
	Bps  common.Float
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
