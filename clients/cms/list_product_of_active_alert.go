package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListProductOfActiveAlertRequest struct {
	requests.RpcRequest
	UserId string `position:"Query" name:"UserId"`
}

func (req *ListProductOfActiveAlertRequest) Invoke(client *sdk.Client) (resp *ListProductOfActiveAlertResponse, err error) {
	req.InitWithApiInfo("Cms", "2017-03-01", "ListProductOfActiveAlert", "cms", "")
	resp = &ListProductOfActiveAlertResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListProductOfActiveAlertResponse struct {
	responses.BaseResponse
	RequestId  string
	Success    bool
	Code       int
	Message    string
	Datapoints string
}
