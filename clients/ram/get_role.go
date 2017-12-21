package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetRoleRequest struct {
	requests.RpcRequest
	RoleName string `position:"Query" name:"RoleName"`
}

func (req *GetRoleRequest) Invoke(client *sdk.Client) (resp *GetRoleResponse, err error) {
	req.InitWithApiInfo("Ram", "2015-05-01", "GetRole", "", "")
	resp = &GetRoleResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetRoleResponse struct {
	responses.BaseResponse
	RequestId string
	Role      GetRoleRole
}

type GetRoleRole struct {
	RoleId                   string
	RoleName                 string
	Arn                      string
	Description              string
	AssumeRolePolicyDocument string
	CreateDate               string
	UpdateDate               string
}
