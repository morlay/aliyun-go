package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ApgroupBatchAddApRequest struct {
	ApAssetIds string `position:"Query" name:"ApAssetIds"`
	ApgroupId  int64  `position:"Query" name:"ApgroupId"`
}

func (r ApgroupBatchAddApRequest) Invoke(client *sdk.Client) (response *ApgroupBatchAddApResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ApgroupBatchAddApRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ApgroupBatchAddAp", "", "")

	resp := struct {
		*responses.BaseResponse
		ApgroupBatchAddApResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ApgroupBatchAddApResponse

	err = client.DoAction(&req, &resp)
	return
}

type ApgroupBatchAddApResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
