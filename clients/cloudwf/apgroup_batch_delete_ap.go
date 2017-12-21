package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ApgroupBatchDeleteApRequest struct {
	ApAssetIds string `position:"Query" name:"ApAssetIds"`
}

func (r ApgroupBatchDeleteApRequest) Invoke(client *sdk.Client) (response *ApgroupBatchDeleteApResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ApgroupBatchDeleteApRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ApgroupBatchDeleteAp", "", "")

	resp := struct {
		*responses.BaseResponse
		ApgroupBatchDeleteApResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ApgroupBatchDeleteApResponse

	err = client.DoAction(&req, &resp)
	return
}

type ApgroupBatchDeleteApResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
