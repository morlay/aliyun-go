package aegis

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId common.String
	Strategy  DescribeStratetyDetailStrategy
}

type DescribeStratetyDetailStrategy struct {
	CycleDays                        common.Integer
	Name                             common.String
	Id                               common.Integer
	CycleStartTime                   common.Integer
	Type                             common.Integer
	RiskTypeWhiteListQueryResultList DescribeStratetyDetailRiskTypeWhiteListQueryResultList
}

type DescribeStratetyDetailRiskTypeWhiteListQueryResult struct {
	TypeName common.String
	Alias    common.String
	On       bool
	SubTypes DescribeStratetyDetailSubTypList
}

type DescribeStratetyDetailSubTyp struct {
	TypeName common.String
	Alias    common.String
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
