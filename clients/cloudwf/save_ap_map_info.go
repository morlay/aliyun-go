package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SaveApMapInfoRequest struct {
	JsonData string `position:"Query" name:"JsonData"`
}

func (r SaveApMapInfoRequest) Invoke(client *sdk.Client) (response *SaveApMapInfoResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SaveApMapInfoRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "SaveApMapInfo", "", "")

	resp := struct {
		*responses.BaseResponse
		SaveApMapInfoResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SaveApMapInfoResponse

	err = client.DoAction(&req, &resp)
	return
}

type SaveApMapInfoResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
