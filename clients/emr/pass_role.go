package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type PassRoleRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ConsoleRoleName string `position:"Query" name:"ConsoleRoleName"`
}

func (req *PassRoleRequest) Invoke(client *sdk.Client) (resp *PassRoleResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "PassRole", "", "")
	resp = &PassRoleResponse{}
	err = client.DoAction(req, resp)
	return
}

type PassRoleResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
}
