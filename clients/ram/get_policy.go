package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetPolicyRequest struct {
	PolicyType string `position:"Query" name:"PolicyType"`
	PolicyName string `position:"Query" name:"PolicyName"`
}

func (r GetPolicyRequest) Invoke(client *sdk.Client) (response *GetPolicyResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetPolicyRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ram", "2015-05-01", "GetPolicy", "", "")

	resp := struct {
		*responses.BaseResponse
		GetPolicyResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetPolicyResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetPolicyResponse struct {
	RequestId string
	Policy    GetPolicyPolicy
}

type GetPolicyPolicy struct {
	PolicyName      string
	PolicyType      string
	Description     string
	DefaultVersion  string
	PolicyDocument  string
	CreateDate      string
	UpdateDate      string
	AttachmentCount int
}
