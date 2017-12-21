package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AttachPolicyToRoleRequest struct {
	PolicyType string `position:"Query" name:"PolicyType"`
	RoleName   string `position:"Query" name:"RoleName"`
	PolicyName string `position:"Query" name:"PolicyName"`
}

func (r AttachPolicyToRoleRequest) Invoke(client *sdk.Client) (response *AttachPolicyToRoleResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AttachPolicyToRoleRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ram", "2015-05-01", "AttachPolicyToRole", "", "")

	resp := struct {
		*responses.BaseResponse
		AttachPolicyToRoleResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AttachPolicyToRoleResponse

	err = client.DoAction(&req, &resp)
	return
}

type AttachPolicyToRoleResponse struct {
	RequestId string
}
