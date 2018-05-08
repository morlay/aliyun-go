package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ApgroupBatchDeleteApRequest struct {
	requests.RpcRequest
	ApAssetIds string `position:"Query" name:"ApAssetIds"`
}

func (req *ApgroupBatchDeleteApRequest) Invoke(client *sdk.Client) (resp *ApgroupBatchDeleteApResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ApgroupBatchDeleteAp", "", "")
	resp = &ApgroupBatchDeleteApResponse{}
	err = client.DoAction(req, resp)
	return
}

type ApgroupBatchDeleteApResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
	Message   common.String
	Data      common.String
	ErrorCode common.Integer
	ErrorMsg  common.String
}
