package nas

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeAccessRulesRequest struct {
	requests.RpcRequest
	PageSize        int    `position:"Query" name:"PageSize"`
	AccessGroupName string `position:"Query" name:"AccessGroupName"`
	AccessRuleId    string `position:"Query" name:"AccessRuleId"`
	PageNumber      int    `position:"Query" name:"PageNumber"`
}

func (req *DescribeAccessRulesRequest) Invoke(client *sdk.Client) (resp *DescribeAccessRulesResponse, err error) {
	req.InitWithApiInfo("NAS", "2017-06-26", "DescribeAccessRules", "nas", "")
	resp = &DescribeAccessRulesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeAccessRulesResponse struct {
	responses.BaseResponse
	RequestId   common.String
	TotalCount  common.Integer
	PageSize    common.Integer
	PageNumber  common.Integer
	AccessRules DescribeAccessRulesAccessRuleList
}

type DescribeAccessRulesAccessRule struct {
	SourceCidrIp common.String
	Priority     common.Integer
	AccessRuleId common.String
	RWAccess     common.String
	UserAccess   common.String
}

type DescribeAccessRulesAccessRuleList []DescribeAccessRulesAccessRule

func (list *DescribeAccessRulesAccessRuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAccessRulesAccessRule)
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
