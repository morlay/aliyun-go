package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type PutSystemEventRequest struct {
	EventInfo string `position:"Query" name:"EventInfo"`
}

func (r PutSystemEventRequest) Invoke(client *sdk.Client) (response *PutSystemEventResponse, err error) {
	req := struct {
		*requests.RpcRequest
		PutSystemEventRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cms", "2017-03-01", "PutSystemEvent", "cms", "")

	resp := struct {
		*responses.BaseResponse
		PutSystemEventResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.PutSystemEventResponse

	err = client.DoAction(&req, &resp)
	return
}

type PutSystemEventResponse struct {
	Code    string
	Message string
	Data    string
}
