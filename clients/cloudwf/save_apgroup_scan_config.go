package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SaveApgroupScanConfigRequest struct {
	JsonData  string `position:"Query" name:"JsonData"`
	ApgroupId int64  `position:"Query" name:"ApgroupId"`
}

func (r SaveApgroupScanConfigRequest) Invoke(client *sdk.Client) (response *SaveApgroupScanConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SaveApgroupScanConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "SaveApgroupScanConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		SaveApgroupScanConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SaveApgroupScanConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type SaveApgroupScanConfigResponse struct {
	RequestId string
	Success   bool
	Message   string
	ErrorCode int
	ErrorMsg  string
}
