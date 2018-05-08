package ram

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListPoliciesForUserRequest struct {
	requests.RpcRequest
	UserName string `position:"Query" name:"UserName"`
}

func (req *ListPoliciesForUserRequest) Invoke(client *sdk.Client) (resp *ListPoliciesForUserResponse, err error) {
	req.InitWithApiInfo("Ram", "2015-05-01", "ListPoliciesForUser", "", "")
	resp = &ListPoliciesForUserResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListPoliciesForUserResponse struct {
	responses.BaseResponse
	RequestId common.String
	Policies  ListPoliciesForUserPolicyList
}

type ListPoliciesForUserPolicy struct {
	PolicyName     common.String
	PolicyType     common.String
	Description    common.String
	DefaultVersion common.String
	AttachDate     common.String
}

type ListPoliciesForUserPolicyList []ListPoliciesForUserPolicy

func (list *ListPoliciesForUserPolicyList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListPoliciesForUserPolicy)
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
