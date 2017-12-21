package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type BusinessInfoRequest struct {
	requests.RpcRequest
	Bid int64 `position:"Query" name:"Bid"`
}

func (req *BusinessInfoRequest) Invoke(client *sdk.Client) (resp *BusinessInfoResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "BusinessInfo", "", "")
	resp = &BusinessInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type BusinessInfoResponse struct {
	responses.BaseResponse
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
