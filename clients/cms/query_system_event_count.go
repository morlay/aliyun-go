package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QuerySystemEventCountRequest struct {
	requests.RpcRequest
	QueryJson string `position:"Query" name:"QueryJson"`
}

func (req *QuerySystemEventCountRequest) Invoke(client *sdk.Client) (resp *QuerySystemEventCountResponse, err error) {
	req.InitWithApiInfo("Cms", "2017-03-01", "QuerySystemEventCount", "cms", "")
	resp = &QuerySystemEventCountResponse{}
	err = client.DoAction(req, resp)
	return
}

type QuerySystemEventCountResponse struct {
	responses.BaseResponse
	Code    string
	Message string
	Data    string
}
