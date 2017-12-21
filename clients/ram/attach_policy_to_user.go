package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AttachPolicyToUserRequest struct {
	PolicyType string `position:"Query" name:"PolicyType"`
	PolicyName string `position:"Query" name:"PolicyName"`
	UserName   string `position:"Query" name:"UserName"`
}

func (r AttachPolicyToUserRequest) Invoke(client *sdk.Client) (response *AttachPolicyToUserResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AttachPolicyToUserRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ram", "2015-05-01", "AttachPolicyToUser", "", "")

	resp := struct {
		*responses.BaseResponse
		AttachPolicyToUserResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AttachPolicyToUserResponse

	err = client.DoAction(&req, &resp)
	return
}

type AttachPolicyToUserResponse struct {
	RequestId string
}
