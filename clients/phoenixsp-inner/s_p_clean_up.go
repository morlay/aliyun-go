package phoenixsp_inner

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SPCleanUpRequest struct {
	requests.RpcRequest
	UserId int64 `position:"Query" name:"UserId"`
}

func (req *SPCleanUpRequest) Invoke(client *sdk.Client) (resp *SPCleanUpResponse, err error) {
	req.InitWithApiInfo("Phoenixsp-inner", "2016-08-05", "SPCleanUp", "", "")
	resp = &SPCleanUpResponse{}
	err = client.DoAction(req, resp)
	return
}

type SPCleanUpResponse struct {
	responses.BaseResponse
	RequestId string
	Code      string
	Message   string
	Success   bool
}
