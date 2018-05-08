package ram

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListPoliciesRequest struct {
	requests.RpcRequest
	PolicyType string `position:"Query" name:"PolicyType"`
	Marker     string `position:"Query" name:"Marker"`
	MaxItems   int    `position:"Query" name:"MaxItems"`
}

func (req *ListPoliciesRequest) Invoke(client *sdk.Client) (resp *ListPoliciesResponse, err error) {
	req.InitWithApiInfo("Ram", "2015-05-01", "ListPolicies", "", "")
	resp = &ListPoliciesResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListPoliciesResponse struct {
	responses.BaseResponse
	RequestId   common.String
	IsTruncated bool
	Marker      common.String
	Policies    ListPoliciesPolicyList
}

type ListPoliciesPolicy struct {
	PolicyName      common.String
	PolicyType      common.String
	Description     common.String
	DefaultVersion  common.String
	CreateDate      common.String
	UpdateDate      common.String
	AttachmentCount common.Integer
}

type ListPoliciesPolicyList []ListPoliciesPolicy

func (list *ListPoliciesPolicyList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListPoliciesPolicy)
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
