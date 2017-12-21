package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SaveApgroupSsidConfigRequest struct {
	JsonData string `position:"Query" name:"JsonData"`
}

func (r SaveApgroupSsidConfigRequest) Invoke(client *sdk.Client) (response *SaveApgroupSsidConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SaveApgroupSsidConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "SaveApgroupSsidConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		SaveApgroupSsidConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SaveApgroupSsidConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type SaveApgroupSsidConfigResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
