package cs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CheckAliyunCSServiceRoleRequest struct {
	requests.RoaRequest
}

func (req *CheckAliyunCSServiceRoleRequest) Invoke(client *sdk.Client) (resp *CheckAliyunCSServiceRoleResponse, err error) {
	req.InitWithApiInfo("CS", "2015-12-15", "CheckAliyunCSServiceRole", "/aliyuncsrole/status", "", "")
	req.Method = "GET"

	resp = &CheckAliyunCSServiceRoleResponse{}
	err = client.DoAction(req, resp)
	return
}

type CheckAliyunCSServiceRoleResponse struct {
	responses.BaseResponse
}
