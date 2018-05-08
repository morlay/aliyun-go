package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type PutEventRequest struct {
	requests.RpcRequest
	EventInfo string `position:"Query" name:"EventInfo"`
}

func (req *PutEventRequest) Invoke(client *sdk.Client) (resp *PutEventResponse, err error) {
	req.InitWithApiInfo("Cms", "2018-03-08", "PutEvent", "cms", "")
	resp = &PutEventResponse{}
	err = client.DoAction(req, resp)
	return
}

type PutEventResponse struct {
	responses.BaseResponse
	Code    common.String
	Message common.String
	Data    common.String
}
