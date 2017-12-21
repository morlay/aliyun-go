package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteRoleRequest struct {
	requests.RpcRequest
	RoleName string `position:"Query" name:"RoleName"`
}

func (req *DeleteRoleRequest) Invoke(client *sdk.Client) (resp *DeleteRoleResponse, err error) {
	req.InitWithApiInfo("Ram", "2015-05-01", "DeleteRole", "", "")
	resp = &DeleteRoleResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteRoleResponse struct {
	responses.BaseResponse
	RequestId string
}
