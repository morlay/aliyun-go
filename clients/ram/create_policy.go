package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreatePolicyRequest struct {
	Description    string `position:"Query" name:"Description"`
	PolicyName     string `position:"Query" name:"PolicyName"`
	PolicyDocument string `position:"Query" name:"PolicyDocument"`
}

func (r CreatePolicyRequest) Invoke(client *sdk.Client) (response *CreatePolicyResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreatePolicyRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ram", "2015-05-01", "CreatePolicy", "", "")

	resp := struct {
		*responses.BaseResponse
		CreatePolicyResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreatePolicyResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreatePolicyResponse struct {
	RequestId string
	Policy    CreatePolicyPolicy
}

type CreatePolicyPolicy struct {
	PolicyName     string
	PolicyType     string
	Description    string
	DefaultVersion string
	CreateDate     string
}
