package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId common.String
	Role      GetRoleRole
}

type GetRoleRole struct {
	RoleId                   common.String
	RoleName                 common.String
	Arn                      common.String
	Description              common.String
	AssumeRolePolicyDocument common.String
	CreateDate               common.String
	UpdateDate               common.String
}
