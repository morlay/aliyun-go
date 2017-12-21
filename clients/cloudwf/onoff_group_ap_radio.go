package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OnoffGroupApRadioRequest struct {
	ApgroupId int64 `position:"Query" name:"ApgroupId"`
	Disabled  int   `position:"Query" name:"Disabled"`
}

func (r OnoffGroupApRadioRequest) Invoke(client *sdk.Client) (response *OnoffGroupApRadioResponse, err error) {
	req := struct {
		*requests.RpcRequest
		OnoffGroupApRadioRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "OnoffGroupApRadio", "", "")

	resp := struct {
		*responses.BaseResponse
		OnoffGroupApRadioResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.OnoffGroupApRadioResponse

	err = client.DoAction(&req, &resp)
	return
}

type OnoffGroupApRadioResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
