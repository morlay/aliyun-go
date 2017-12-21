package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GroupIntimeRequest struct {
	requests.RpcRequest
	Gsid int64 `position:"Query" name:"Gsid"`
}

func (req *GroupIntimeRequest) Invoke(client *sdk.Client) (resp *GroupIntimeResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "GroupIntime", "", "")
	resp = &GroupIntimeResponse{}
	err = client.DoAction(req, resp)
	return
}

type GroupIntimeResponse struct {
	responses.BaseResponse
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
