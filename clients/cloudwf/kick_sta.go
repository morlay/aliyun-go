package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type KickStaRequest struct {
	Id int64 `position:"Query" name:"Id"`
}

func (r KickStaRequest) Invoke(client *sdk.Client) (response *KickStaResponse, err error) {
	req := struct {
		*requests.RpcRequest
		KickStaRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "KickSta", "", "")

	resp := struct {
		*responses.BaseResponse
		KickStaResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.KickStaResponse

	err = client.DoAction(&req, &resp)
	return
}

type KickStaResponse struct {
	RequestId string
	Success   bool
	Message   string
	ErrorCode int
	ErrorMsg  string
}
