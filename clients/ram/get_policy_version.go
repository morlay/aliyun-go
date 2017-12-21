package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetPolicyVersionRequest struct {
	VersionId  string `position:"Query" name:"VersionId"`
	PolicyType string `position:"Query" name:"PolicyType"`
	PolicyName string `position:"Query" name:"PolicyName"`
}

func (r GetPolicyVersionRequest) Invoke(client *sdk.Client) (response *GetPolicyVersionResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetPolicyVersionRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ram", "2015-05-01", "GetPolicyVersion", "", "")

	resp := struct {
		*responses.BaseResponse
		GetPolicyVersionResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetPolicyVersionResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetPolicyVersionResponse struct {
	RequestId     string
	PolicyVersion GetPolicyVersionPolicyVersion
}

type GetPolicyVersionPolicyVersion struct {
	VersionId        string
	IsDefaultVersion bool
	PolicyDocument   string
	CreateDate       string
}
