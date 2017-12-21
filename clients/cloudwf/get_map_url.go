package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetMapUrlRequest struct {
	MapId int64 `position:"Query" name:"MapId"`
}

func (r GetMapUrlRequest) Invoke(client *sdk.Client) (response *GetMapUrlResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetMapUrlRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "GetMapUrl", "", "")

	resp := struct {
		*responses.BaseResponse
		GetMapUrlResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetMapUrlResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetMapUrlResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
