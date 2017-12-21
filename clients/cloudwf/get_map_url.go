package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetMapUrlRequest struct {
	requests.RpcRequest
	MapId int64 `position:"Query" name:"MapId"`
}

func (req *GetMapUrlRequest) Invoke(client *sdk.Client) (resp *GetMapUrlResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "GetMapUrl", "", "")
	resp = &GetMapUrlResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetMapUrlResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
