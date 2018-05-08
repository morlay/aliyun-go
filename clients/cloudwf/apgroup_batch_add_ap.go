package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ApgroupBatchAddApRequest struct {
	requests.RpcRequest
	ApAssetIds string `position:"Query" name:"ApAssetIds"`
	ApgroupId  int64  `position:"Query" name:"ApgroupId"`
}

func (req *ApgroupBatchAddApRequest) Invoke(client *sdk.Client) (resp *ApgroupBatchAddApResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ApgroupBatchAddAp", "", "")
	resp = &ApgroupBatchAddApResponse{}
	err = client.DoAction(req, resp)
	return
}

type ApgroupBatchAddApResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
	Message   common.String
	Data      common.String
	ErrorCode common.Integer
	ErrorMsg  common.String
}
