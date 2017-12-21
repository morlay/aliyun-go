package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SaveGroupApRadioConfigRequest struct {
	JsonData string `position:"Query" name:"JsonData"`
}

func (r SaveGroupApRadioConfigRequest) Invoke(client *sdk.Client) (response *SaveGroupApRadioConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SaveGroupApRadioConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "SaveGroupApRadioConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		SaveGroupApRadioConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SaveGroupApRadioConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type SaveGroupApRadioConfigResponse struct {
	RequestId string
	Success   bool
	Message   string
	ErrorCode int
	ErrorMsg  string
}
