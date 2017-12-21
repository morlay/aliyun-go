package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetRoleRequest struct {
	RoleName string `position:"Query" name:"RoleName"`
}

func (r GetRoleRequest) Invoke(client *sdk.Client) (response *GetRoleResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetRoleRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ram", "2015-05-01", "GetRole", "", "")

	resp := struct {
		*responses.BaseResponse
		GetRoleResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetRoleResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetRoleResponse struct {
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
