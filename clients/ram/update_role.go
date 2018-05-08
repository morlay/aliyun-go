package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId common.String
	Role      UpdateRoleRole
}

type UpdateRoleRole struct {
	RoleId                   common.String
	RoleName                 common.String
	Arn                      common.String
	Description              common.String
	AssumeRolePolicyDocument common.String
	CreateDate               common.String
	UpdateDate               common.String
}
