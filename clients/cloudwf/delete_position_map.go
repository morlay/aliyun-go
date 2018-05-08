package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeletePositionMapRequest struct {
	requests.RpcRequest
	MapId int64 `position:"Query" name:"MapId"`
}

func (req *DeletePositionMapRequest) Invoke(client *sdk.Client) (resp *DeletePositionMapResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "DeletePositionMap", "", "")
	resp = &DeletePositionMapResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeletePositionMapResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
	Message   common.String
	Data      common.String
	ErrorCode common.Integer
	ErrorMsg  common.String
}
