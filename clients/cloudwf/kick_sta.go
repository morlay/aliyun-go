package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type KickStaRequest struct {
	requests.RpcRequest
	Id int64 `position:"Query" name:"Id"`
}

func (req *KickStaRequest) Invoke(client *sdk.Client) (resp *KickStaResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "KickSta", "", "")
	resp = &KickStaResponse{}
	err = client.DoAction(req, resp)
	return
}

type KickStaResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Message   string
	ErrorCode int
	ErrorMsg  string
}
