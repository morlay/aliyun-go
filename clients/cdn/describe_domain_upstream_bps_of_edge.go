package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDomainUpstreamBpsOfEdgeRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeDomainUpstreamBpsOfEdgeRequest) Invoke(client *sdk.Client) (resp *DescribeDomainUpstreamBpsOfEdgeResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainUpstreamBpsOfEdge", "", "")
	resp = &DescribeDomainUpstreamBpsOfEdgeResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDomainUpstreamBpsOfEdgeResponse struct {
	responses.BaseResponse
	RequestId string
	BpsDatas  DescribeDomainUpstreamBpsOfEdgeDomainBpsModelList
}

type DescribeDomainUpstreamBpsOfEdgeDomainBpsModel struct {
	Time string
	Bps  float32
}

type DescribeDomainUpstreamBpsOfEdgeDomainBpsModelList []DescribeDomainUpstreamBpsOfEdgeDomainBpsModel

func (list *DescribeDomainUpstreamBpsOfEdgeDomainBpsModelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainUpstreamBpsOfEdgeDomainBpsModel)
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
