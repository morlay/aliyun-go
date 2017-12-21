package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreatePolicyVersionRequest struct {
	requests.RpcRequest
	SetAsDefault   string `position:"Query" name:"SetAsDefault"`
	PolicyName     string `position:"Query" name:"PolicyName"`
	PolicyDocument string `position:"Query" name:"PolicyDocument"`
}

func (req *CreatePolicyVersionRequest) Invoke(client *sdk.Client) (resp *CreatePolicyVersionResponse, err error) {
	req.InitWithApiInfo("Ram", "2015-05-01", "CreatePolicyVersion", "", "")
	resp = &CreatePolicyVersionResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreatePolicyVersionResponse struct {
	responses.BaseResponse
	RequestId     string
	PolicyVersion CreatePolicyVersionPolicyVersion
}

type CreatePolicyVersionPolicyVersion struct {
	VersionId        string
	IsDefaultVersion bool
	PolicyDocument   string
	CreateDate       string
}
