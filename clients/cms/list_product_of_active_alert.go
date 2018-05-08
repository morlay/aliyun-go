package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListProductOfActiveAlertRequest struct {
	requests.RpcRequest
	UserId string `position:"Query" name:"UserId"`
}

func (req *ListProductOfActiveAlertRequest) Invoke(client *sdk.Client) (resp *ListProductOfActiveAlertResponse, err error) {
	req.InitWithApiInfo("Cms", "2018-03-08", "ListProductOfActiveAlert", "cms", "")
	resp = &ListProductOfActiveAlertResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListProductOfActiveAlertResponse struct {
	responses.BaseResponse
	RequestId  common.String
	Success    bool
	Code       common.Integer
	Message    common.String
	Datapoints common.String
}
