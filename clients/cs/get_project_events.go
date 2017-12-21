package cs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetProjectEventsRequest struct {
	requests.RoaRequest
	ClusterId string `position:"Path" name:"ClusterId"`
	ProjectId string `position:"Path" name:"ProjectId"`
}

func (req *GetProjectEventsRequest) Invoke(client *sdk.Client) (resp *GetProjectEventsResponse, err error) {
	req.InitWithApiInfo("CS", "2015-12-15", "GetProjectEvents", "/clusters/[ClusterId]/projects/[ProjectId]/events", "", "")
	req.Method = "GET"

	resp = &GetProjectEventsResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetProjectEventsResponse struct {
	responses.BaseResponse
}
