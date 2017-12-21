package ram

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListPoliciesForUserRequest struct {
	UserName string `position:"Query" name:"UserName"`
}

func (r ListPoliciesForUserRequest) Invoke(client *sdk.Client) (response *ListPoliciesForUserResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ListPoliciesForUserRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ram", "2015-05-01", "ListPoliciesForUser", "", "")

	resp := struct {
		*responses.BaseResponse
		ListPoliciesForUserResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ListPoliciesForUserResponse

	err = client.DoAction(&req, &resp)
	return
}

type ListPoliciesForUserResponse struct {
	RequestId string
	Policies  ListPoliciesForUserPolicyList
}

type ListPoliciesForUserPolicy struct {
	PolicyName     string
	PolicyType     string
	Description    string
	DefaultVersion string
	AttachDate     string
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
