package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryCustomEventDetailRequest struct {
	requests.RpcRequest
	QueryJson string `position:"Query" name:"QueryJson"`
}

func (req *QueryCustomEventDetailRequest) Invoke(client *sdk.Client) (resp *QueryCustomEventDetailResponse, err error) {
	req.InitWithApiInfo("Cms", "2018-03-08", "QueryCustomEventDetail", "cms", "")
	resp = &QueryCustomEventDetailResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryCustomEventDetailResponse struct {
	responses.BaseResponse
	Code      string
	Message   string
	Data      string
	RequestId string
	Success   string
}
