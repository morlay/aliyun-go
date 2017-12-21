package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetOssServerSignRequest struct {
}

func (r GetOssServerSignRequest) Invoke(client *sdk.Client) (response *GetOssServerSignResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetOssServerSignRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "GetOssServerSign", "", "")

	resp := struct {
		*responses.BaseResponse
		GetOssServerSignResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetOssServerSignResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetOssServerSignResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
