package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CheckUmengDataAnalysisPermissionRequest struct {
}

func (r CheckUmengDataAnalysisPermissionRequest) Invoke(client *sdk.Client) (response *CheckUmengDataAnalysisPermissionResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CheckUmengDataAnalysisPermissionRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "CheckUmengDataAnalysisPermission", "", "")

	resp := struct {
		*responses.BaseResponse
		CheckUmengDataAnalysisPermissionResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CheckUmengDataAnalysisPermissionResponse

	err = client.DoAction(&req, &resp)
	return
}

type CheckUmengDataAnalysisPermissionResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
