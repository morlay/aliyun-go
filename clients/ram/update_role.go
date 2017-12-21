package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdateRoleRequest struct {
	NewAssumeRolePolicyDocument string `position:"Query" name:"NewAssumeRolePolicyDocument"`
	RoleName                    string `position:"Query" name:"RoleName"`
}

func (r UpdateRoleRequest) Invoke(client *sdk.Client) (response *UpdateRoleResponse, err error) {
	req := struct {
		*requests.RpcRequest
		UpdateRoleRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ram", "2015-05-01", "UpdateRole", "", "")

	resp := struct {
		*responses.BaseResponse
		UpdateRoleResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.UpdateRoleResponse

	err = client.DoAction(&req, &resp)
	return
}

type UpdateRoleResponse struct {
	RequestId string
	Role      UpdateRoleRole
}

type UpdateRoleRole struct {
	RoleId                   string
	RoleName                 string
	Arn                      string
	Description              string
	AssumeRolePolicyDocument string
	CreateDate               string
	UpdateDate               string
}
