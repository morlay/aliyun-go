package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DetachPolicyFromRoleRequest struct {
	PolicyType string `position:"Query" name:"PolicyType"`
	RoleName   string `position:"Query" name:"RoleName"`
	PolicyName string `position:"Query" name:"PolicyName"`
}

func (r DetachPolicyFromRoleRequest) Invoke(client *sdk.Client) (response *DetachPolicyFromRoleResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DetachPolicyFromRoleRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ram", "2015-05-01", "DetachPolicyFromRole", "", "")

	resp := struct {
		*responses.BaseResponse
		DetachPolicyFromRoleResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DetachPolicyFromRoleResponse

	err = client.DoAction(&req, &resp)
	return
}

type DetachPolicyFromRoleResponse struct {
	RequestId string
}
