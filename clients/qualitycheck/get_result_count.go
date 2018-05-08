package qualitycheck

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetResultCountRequest struct {
	requests.RpcRequest
	JsonStr string `position:"Query" name:"JsonStr"`
}

func (req *GetResultCountRequest) Invoke(client *sdk.Client) (resp *GetResultCountResponse, err error) {
	req.InitWithApiInfo("Qualitycheck", "2016-08-01", "GetResultCount", "", "")
	resp = &GetResultCountResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetResultCountResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
	Code      common.String
	Message   common.String
	Data      common.Integer
}
