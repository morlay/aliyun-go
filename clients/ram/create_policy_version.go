package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreatePolicyVersionRequest struct {
	SetAsDefault   string `position:"Query" name:"SetAsDefault"`
	PolicyName     string `position:"Query" name:"PolicyName"`
	PolicyDocument string `position:"Query" name:"PolicyDocument"`
}

func (r CreatePolicyVersionRequest) Invoke(client *sdk.Client) (response *CreatePolicyVersionResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreatePolicyVersionRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ram", "2015-05-01", "CreatePolicyVersion", "", "")

	resp := struct {
		*responses.BaseResponse
		CreatePolicyVersionResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreatePolicyVersionResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreatePolicyVersionResponse struct {
	RequestId     string
	PolicyVersion CreatePolicyVersionPolicyVersion
}

type CreatePolicyVersionPolicyVersion struct {
	VersionId        string
	IsDefaultVersion bool
	PolicyDocument   string
	CreateDate       string
}
