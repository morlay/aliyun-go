package ess

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeScalingRulesRequest struct {
	requests.RpcRequest
	ScalingRuleName1     string `position:"Query" name:"ScalingRuleName.1"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ScalingRuleName2     string `position:"Query" name:"ScalingRuleName.2"`
	ScalingRuleName3     string `position:"Query" name:"ScalingRuleName.3"`
	ScalingRuleName4     string `position:"Query" name:"ScalingRuleName.4"`
	ScalingRuleName5     string `position:"Query" name:"ScalingRuleName.5"`
	ScalingGroupId       string `position:"Query" name:"ScalingGroupId"`
	ScalingRuleName6     string `position:"Query" name:"ScalingRuleName.6"`
	ScalingRuleName7     string `position:"Query" name:"ScalingRuleName.7"`
	ScalingRuleName8     string `position:"Query" name:"ScalingRuleName.8"`
	ScalingRuleAri9      string `position:"Query" name:"ScalingRuleAri.9"`
	ScalingRuleName9     string `position:"Query" name:"ScalingRuleName.9"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
	PageSize             int    `position:"Query" name:"PageSize"`
	ScalingRuleId10      string `position:"Query" name:"ScalingRuleId.10"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ScalingRuleAri1      string `position:"Query" name:"ScalingRuleAri.1"`
	ScalingRuleAri2      string `position:"Query" name:"ScalingRuleAri.2"`
	ScalingRuleName10    string `position:"Query" name:"ScalingRuleName.10"`
	ScalingRuleAri3      string `position:"Query" name:"ScalingRuleAri.3"`
	ScalingRuleAri4      string `position:"Query" name:"ScalingRuleAri.4"`
	ScalingRuleId8       string `position:"Query" name:"ScalingRuleId.8"`
	ScalingRuleAri5      string `position:"Query" name:"ScalingRuleAri.5"`
	ScalingRuleId9       string `position:"Query" name:"ScalingRuleId.9"`
	ScalingRuleAri6      string `position:"Query" name:"ScalingRuleAri.6"`
	ScalingRuleAri7      string `position:"Query" name:"ScalingRuleAri.7"`
	ScalingRuleAri10     string `position:"Query" name:"ScalingRuleAri.10"`
	ScalingRuleAri8      string `position:"Query" name:"ScalingRuleAri.8"`
	ScalingRuleId4       string `position:"Query" name:"ScalingRuleId.4"`
	ScalingRuleId5       string `position:"Query" name:"ScalingRuleId.5"`
	ScalingRuleId6       string `position:"Query" name:"ScalingRuleId.6"`
	ScalingRuleId7       string `position:"Query" name:"ScalingRuleId.7"`
	ScalingRuleId1       string `position:"Query" name:"ScalingRuleId.1"`
	ScalingRuleId2       string `position:"Query" name:"ScalingRuleId.2"`
	ScalingRuleId3       string `position:"Query" name:"ScalingRuleId.3"`
}

func (req *DescribeScalingRulesRequest) Invoke(client *sdk.Client) (resp *DescribeScalingRulesResponse, err error) {
	req.InitWithApiInfo("Ess", "2014-08-28", "DescribeScalingRules", "ess", "")
	resp = &DescribeScalingRulesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeScalingRulesResponse struct {
	responses.BaseResponse
	TotalCount   common.Integer
	PageNumber   common.Integer
	PageSize     common.Integer
	RequestId    common.String
	ScalingRules DescribeScalingRulesScalingRuleList
}

type DescribeScalingRulesScalingRule struct {
	ScalingRuleId   common.String
	ScalingGroupId  common.String
	ScalingRuleName common.String
	Cooldown        common.Integer
	AdjustmentType  common.String
	AdjustmentValue common.Integer
	MinSize         common.Integer
	MaxSize         common.Integer
	ScalingRuleAri  common.String
}

type DescribeScalingRulesScalingRuleList []DescribeScalingRulesScalingRule

func (list *DescribeScalingRulesScalingRuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeScalingRulesScalingRule)
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
