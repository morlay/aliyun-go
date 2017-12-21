package ram

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListPoliciesForRoleRequest struct {
	requests.RpcRequest
	RoleName string `position:"Query" name:"RoleName"`
}

func (req *ListPoliciesForRoleRequest) Invoke(client *sdk.Client) (resp *ListPoliciesForRoleResponse, err error) {
	req.InitWithApiInfo("Ram", "2015-05-01", "ListPoliciesForRole", "", "")
	resp = &ListPoliciesForRoleResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListPoliciesForRoleResponse struct {
	responses.BaseResponse
	RequestId string
	Policies  ListPoliciesForRolePolicyList
}

type ListPoliciesForRolePolicy struct {
	PolicyName     string
	PolicyType     string
	Description    string
	DefaultVersion string
	AttachDate     string
}

type ListPoliciesForRolePolicyList []ListPoliciesForRolePolicy

func (list *ListPoliciesForRolePolicyList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListPoliciesForRolePolicy)
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
