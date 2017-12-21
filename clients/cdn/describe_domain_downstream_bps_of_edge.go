package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDomainDownstreamBpsOfEdgeRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeDomainDownstreamBpsOfEdgeRequest) Invoke(client *sdk.Client) (resp *DescribeDomainDownstreamBpsOfEdgeResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainDownstreamBpsOfEdge", "", "")
	resp = &DescribeDomainDownstreamBpsOfEdgeResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDomainDownstreamBpsOfEdgeResponse struct {
	responses.BaseResponse
	RequestId string
	BpsDatas  DescribeDomainDownstreamBpsOfEdgeDomainBpsModelList
}

type DescribeDomainDownstreamBpsOfEdgeDomainBpsModel struct {
	Time string
	Bps  float32
}

type DescribeDomainDownstreamBpsOfEdgeDomainBpsModelList []DescribeDomainDownstreamBpsOfEdgeDomainBpsModel

func (list *DescribeDomainDownstreamBpsOfEdgeDomainBpsModelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainDownstreamBpsOfEdgeDomainBpsModel)
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
