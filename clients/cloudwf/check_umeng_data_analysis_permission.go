package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CheckUmengDataAnalysisPermissionRequest struct {
	requests.RpcRequest
}

func (req *CheckUmengDataAnalysisPermissionRequest) Invoke(client *sdk.Client) (resp *CheckUmengDataAnalysisPermissionResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "CheckUmengDataAnalysisPermission", "", "")
	resp = &CheckUmengDataAnalysisPermissionResponse{}
	err = client.DoAction(req, resp)
	return
}

type CheckUmengDataAnalysisPermissionResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
