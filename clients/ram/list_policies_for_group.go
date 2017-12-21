package ram

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListPoliciesForGroupRequest struct {
	GroupName string `position:"Query" name:"GroupName"`
}

func (r ListPoliciesForGroupRequest) Invoke(client *sdk.Client) (response *ListPoliciesForGroupResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ListPoliciesForGroupRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ram", "2015-05-01", "ListPoliciesForGroup", "", "")

	resp := struct {
		*responses.BaseResponse
		ListPoliciesForGroupResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ListPoliciesForGroupResponse

	err = client.DoAction(&req, &resp)
	return
}

type ListPoliciesForGroupResponse struct {
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
