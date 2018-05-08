package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DetachPolicyFromRoleRequest struct {
	requests.RpcRequest
	PolicyType string `position:"Query" name:"PolicyType"`
	RoleName   string `position:"Query" name:"RoleName"`
	PolicyName string `position:"Query" name:"PolicyName"`
}

func (req *DetachPolicyFromRoleRequest) Invoke(client *sdk.Client) (resp *DetachPolicyFromRoleResponse, err error) {
	req.InitWithApiInfo("Ram", "2015-05-01", "DetachPolicyFromRole", "", "")
	resp = &DetachPolicyFromRoleResponse{}
	err = client.DoAction(req, resp)
	return
}

type DetachPolicyFromRoleResponse struct {
	responses.BaseResponse
	RequestId common.String
}
