package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SavePortalConfigRequest struct {
	JsonData string `position:"Query" name:"JsonData"`
}

func (r SavePortalConfigRequest) Invoke(client *sdk.Client) (response *SavePortalConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SavePortalConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "SavePortalConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		SavePortalConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SavePortalConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type SavePortalConfigResponse struct {
	RequestId string
	Success   bool
	Message   string
	ErrorCode int
	ErrorMsg  string
}
