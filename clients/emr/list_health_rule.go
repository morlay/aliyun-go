package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListHealthRuleRequest struct {
	requests.RpcRequest
	Component       string `position:"Query" name:"Component"`
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Service         string `position:"Query" name:"Service"`
	PageSize        int    `position:"Query" name:"PageSize"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	PageNumber      int    `position:"Query" name:"PageNumber"`
}

func (req *ListHealthRuleRequest) Invoke(client *sdk.Client) (resp *ListHealthRuleResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ListHealthRule", "", "")
	resp = &ListHealthRuleResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListHealthRuleResponse struct {
	responses.BaseResponse
	RequestId  common.String
	PageNumber common.Integer
	PageSize   common.Integer
	Total      common.Integer
	Rule       ListHealthRuleRuleItemList
}

type ListHealthRuleRuleItem struct {
	Id          common.Long
	Name        common.String
	Status      common.String
	Service     common.String
	Component   common.String
	Title       common.String
	Description common.String
	Explanation common.String
	Solution    common.String
}

type ListHealthRuleRuleItemList []ListHealthRuleRuleItem

func (list *ListHealthRuleRuleItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListHealthRuleRuleItem)
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
