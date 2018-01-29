package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QuerySystemEventDetailRequest struct {
	requests.RpcRequest
	QueryJson string `position:"Query" name:"QueryJson"`
}

func (req *QuerySystemEventDetailRequest) Invoke(client *sdk.Client) (resp *QuerySystemEventDetailResponse, err error) {
	req.InitWithApiInfo("Cms", "2017-03-01", "QuerySystemEventDetail", "cms", "")
	resp = &QuerySystemEventDetailResponse{}
	err = client.DoAction(req, resp)
	return
}

type QuerySystemEventDetailResponse struct {
	responses.BaseResponse
	Code    string
	Message string
	Data    string
}
