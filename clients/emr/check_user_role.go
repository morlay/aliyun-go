package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CheckUserRoleRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ConsoleRoleName string `position:"Query" name:"ConsoleRoleName"`
}

func (req *CheckUserRoleRequest) Invoke(client *sdk.Client) (resp *CheckUserRoleResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "CheckUserRole", "", "")
	resp = &CheckUserRoleResponse{}
	err = client.DoAction(req, resp)
	return
}

type CheckUserRoleResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
}
