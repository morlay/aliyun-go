package ram

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListPoliciesForGroupRequest struct {
	requests.RpcRequest
	GroupName string `position:"Query" name:"GroupName"`
}

func (req *ListPoliciesForGroupRequest) Invoke(client *sdk.Client) (resp *ListPoliciesForGroupResponse, err error) {
	req.InitWithApiInfo("Ram", "2015-05-01", "ListPoliciesForGroup", "", "")
	resp = &ListPoliciesForGroupResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListPoliciesForGroupResponse struct {
	responses.BaseResponse
	RequestId string
	Policies  ListPoliciesForGroupPolicyList
}

type ListPoliciesForGroupPolicy struct {
	PolicyName     string
	PolicyType     string
	Description    string
	DefaultVersion string
	AttachDate     string
}

type ListPoliciesForGroupPolicyList []ListPoliciesForGroupPolicy

func (list *ListPoliciesForGroupPolicyList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListPoliciesForGroupPolicy)
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
