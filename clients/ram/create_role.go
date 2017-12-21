package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateRoleRequest struct {
	requests.RpcRequest
	RoleName                 string `position:"Query" name:"RoleName"`
	Description              string `position:"Query" name:"Description"`
	AssumeRolePolicyDocument string `position:"Query" name:"AssumeRolePolicyDocument"`
}

func (req *CreateRoleRequest) Invoke(client *sdk.Client) (resp *CreateRoleResponse, err error) {
	req.InitWithApiInfo("Ram", "2015-05-01", "CreateRole", "", "")
	resp = &CreateRoleResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateRoleResponse struct {
	responses.BaseResponse
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
