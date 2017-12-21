package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteRoleRequest struct {
	RoleName string `position:"Query" name:"RoleName"`
}

func (r DeleteRoleRequest) Invoke(client *sdk.Client) (response *DeleteRoleResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteRoleRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ram", "2015-05-01", "DeleteRole", "", "")

	resp := struct {
		*responses.BaseResponse
		DeleteRoleResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteRoleResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteRoleResponse struct {
	RequestId string
}
