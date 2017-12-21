package qualitycheck

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId string
	Success   bool
	Code      string
	Message   string
	Data      int
}
