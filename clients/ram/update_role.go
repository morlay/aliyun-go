package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdateRoleRequest struct {
	requests.RpcRequest
	NewAssumeRolePolicyDocument string `position:"Query" name:"NewAssumeRolePolicyDocument"`
	RoleName                    string `position:"Query" name:"RoleName"`
}

func (req *UpdateRoleRequest) Invoke(client *sdk.Client) (resp *UpdateRoleResponse, err error) {
	req.InitWithApiInfo("Ram", "2015-05-01", "UpdateRole", "", "")
	resp = &UpdateRoleResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpdateRoleResponse struct {
	responses.BaseResponse
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
