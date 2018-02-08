package aegis

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeStratetyDetailRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	Id              string `position:"Query" name:"Id"`
}

func (req *DescribeStratetyDetailRequest) Invoke(client *sdk.Client) (resp *DescribeStratetyDetailResponse, err error) {
	req.InitWithApiInfo("aegis", "2016-11-11", "DescribeStratetyDetail", "vipaegis", "")
	resp = &DescribeStratetyDetailResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeStratetyDetailResponse struct {
	responses.BaseResponse
	RequestId string
	Strategy  DescribeStratetyDetailStrategy
}

type DescribeStratetyDetailStrategy struct {
	CycleDays                        int
	Name                             string
	Id                               int
	CycleStartTime                   int
	Type                             int
	RiskTypeWhiteListQueryResultList DescribeStratetyDetailRiskTypeWhiteListQueryResultList
}

type DescribeStratetyDetailRiskTypeWhiteListQueryResult struct {
	TypeName string
	Alias    string
	On       bool
	SubTypes DescribeStratetyDetailSubTypList
}

type DescribeStratetyDetailSubTyp struct {
	TypeName string
	Alias    string
	On       bool
}

type DescribeStratetyDetailRiskTypeWhiteListQueryResultList []DescribeStratetyDetailRiskTypeWhiteListQueryResult

func (list *DescribeStratetyDetailRiskTypeWhiteListQueryResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeStratetyDetailRiskTypeWhiteListQueryResult)
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

type DescribeStratetyDetailSubTypList []DescribeStratetyDetailSubTyp

func (list *DescribeStratetyDetailSubTypList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeStratetyDetailSubTyp)
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
