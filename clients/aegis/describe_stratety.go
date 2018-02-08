package aegis

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeStratetyRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
}

func (req *DescribeStratetyRequest) Invoke(client *sdk.Client) (resp *DescribeStratetyResponse, err error) {
	req.InitWithApiInfo("aegis", "2016-11-11", "DescribeStratety", "vipaegis", "")
	resp = &DescribeStratetyResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeStratetyResponse struct {
	responses.BaseResponse
	RequestId  string
	Count      int
	Strategies DescribeStratetyDataList
}

type DescribeStratetyData struct {
	CycleDays      int
	Id             int
	CycleStartTime int
	Type           int
	Name           string
	RiskCount      int
	EcsCount       int
	ConfigTargets  DescribeStratetyConfigTargetList
}

type DescribeStratetyConfigTarget struct {
	Flag       string
	TargetType string
	Target     string
}

type DescribeStratetyDataList []DescribeStratetyData

func (list *DescribeStratetyDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeStratetyData)
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

type DescribeStratetyConfigTargetList []DescribeStratetyConfigTarget

func (list *DescribeStratetyConfigTargetList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeStratetyConfigTarget)
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
