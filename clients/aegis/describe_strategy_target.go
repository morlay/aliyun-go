package aegis

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeStrategyTargetRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	Type            string `position:"Query" name:"Type"`
	Config          string `position:"Query" name:"Config"`
	Target          string `position:"Query" name:"Target"`
}

func (req *DescribeStrategyTargetRequest) Invoke(client *sdk.Client) (resp *DescribeStrategyTargetResponse, err error) {
	req.InitWithApiInfo("aegis", "2016-11-11", "DescribeStrategyTarget", "vipaegis", "")
	resp = &DescribeStrategyTargetResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeStrategyTargetResponse struct {
	responses.BaseResponse
	RequestId       string
	Count           int
	StrategyTargets DescribeStrategyTargetStringItemList
}

type DescribeStrategyTargetStringItem struct {
	Flag       string
	Target     string
	TargetType string
}

type DescribeStrategyTargetStringItemList []DescribeStrategyTargetStringItem

func (list *DescribeStrategyTargetStringItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeStrategyTargetStringItem)
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
