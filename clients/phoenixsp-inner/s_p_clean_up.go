package phoenixsp_inner

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId common.String
	Code      common.String
	Message   common.String
	Success   bool
}
