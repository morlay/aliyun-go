package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type PutSystemEventRequest struct {
	requests.RpcRequest
	EventInfo string `position:"Query" name:"EventInfo"`
}

func (req *PutSystemEventRequest) Invoke(client *sdk.Client) (resp *PutSystemEventResponse, err error) {
	req.InitWithApiInfo("Cms", "2017-03-01", "PutSystemEvent", "cms", "")
	resp = &PutSystemEventResponse{}
	err = client.DoAction(req, resp)
	return
}

type PutSystemEventResponse struct {
	responses.BaseResponse
	Code    string
	Message string
	Data    string
}
