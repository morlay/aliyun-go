package ram

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListPolicyVersionsRequest struct {
	requests.RpcRequest
	PolicyType string `position:"Query" name:"PolicyType"`
	PolicyName string `position:"Query" name:"PolicyName"`
}

func (req *ListPolicyVersionsRequest) Invoke(client *sdk.Client) (resp *ListPolicyVersionsResponse, err error) {
	req.InitWithApiInfo("Ram", "2015-05-01", "ListPolicyVersions", "", "")
	resp = &ListPolicyVersionsResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListPolicyVersionsResponse struct {
	responses.BaseResponse
	RequestId      common.String
	PolicyVersions ListPolicyVersionsPolicyVersionList
}

type ListPolicyVersionsPolicyVersion struct {
	VersionId        common.String
	IsDefaultVersion bool
	PolicyDocument   common.String
	CreateDate       common.String
}

type ListPolicyVersionsPolicyVersionList []ListPolicyVersionsPolicyVersion

func (list *ListPolicyVersionsPolicyVersionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListPolicyVersionsPolicyVersion)
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
