package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateRoleRequest struct {
	RoleName                 string `position:"Query" name:"RoleName"`
	Description              string `position:"Query" name:"Description"`
	AssumeRolePolicyDocument string `position:"Query" name:"AssumeRolePolicyDocument"`
}

func (r CreateRoleRequest) Invoke(client *sdk.Client) (response *CreateRoleResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateRoleRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ram", "2015-05-01", "CreateRole", "", "")

	resp := struct {
		*responses.BaseResponse
		CreateRoleResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateRoleResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateRoleResponse struct {
	RequestId string
	Role      CreateRoleRole
}

type CreateRoleRole struct {
	RoleId                   string
	RoleName                 string
	Arn                      string
	Description              string
	AssumeRolePolicyDocument string
	CreateDate               string
}
