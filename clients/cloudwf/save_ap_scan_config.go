package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SaveApScanConfigRequest struct {
	JsonData   string `position:"Query" name:"JsonData"`
	ApConfigId int64  `position:"Query" name:"ApConfigId"`
}

func (r SaveApScanConfigRequest) Invoke(client *sdk.Client) (response *SaveApScanConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SaveApScanConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "SaveApScanConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		SaveApScanConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SaveApScanConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type SaveApScanConfigResponse struct {
	RequestId string
	Success   bool
	Message   string
	ErrorCode int
	ErrorMsg  string
}
