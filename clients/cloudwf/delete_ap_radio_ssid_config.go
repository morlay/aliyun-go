package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteApRadioSsidConfigRequest struct {
	InstantlyEffective int   `position:"Query" name:"InstantlyEffective"`
	Id                 int64 `position:"Query" name:"Id"`
}

func (r DeleteApRadioSsidConfigRequest) Invoke(client *sdk.Client) (response *DeleteApRadioSsidConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteApRadioSsidConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "DeleteApRadioSsidConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		DeleteApRadioSsidConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteApRadioSsidConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteApRadioSsidConfigResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
