package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId  string
	PageNumber int
	PageSize   int
	Total      int
	Rule       ListHealthRuleRuleItemList
}

type ListHealthRuleRuleItem struct {
	Id          int64
	Name        string
	Status      string
	Service     string
	Component   string
	Title       string
	Description string
	Explanation string
	Solution    string
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
