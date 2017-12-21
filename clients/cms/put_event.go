package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type PutEventRequest struct {
	EventInfo string `position:"Query" name:"EventInfo"`
}

func (r PutEventRequest) Invoke(client *sdk.Client) (response *PutEventResponse, err error) {
	req := struct {
		*requests.RpcRequest
		PutEventRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cms", "2017-03-01", "PutEvent", "cms", "")

	resp := struct {
		*responses.BaseResponse
		PutEventResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.PutEventResponse

	err = client.DoAction(&req, &resp)
	return
}

type PutEventResponse struct {
	Code    string
	Message string
	Data    string
}
