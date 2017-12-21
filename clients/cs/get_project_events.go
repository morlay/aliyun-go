package cs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetProjectEventsRequest struct {
	ClusterId string `position:"Path" name:"ClusterId"`
	ProjectId string `position:"Path" name:"ProjectId"`
}

func (r GetProjectEventsRequest) Invoke(client *sdk.Client) (response *GetProjectEventsResponse, err error) {
	req := struct {
		*requests.RoaRequest
		GetProjectEventsRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("CS", "2015-12-15", "GetProjectEvents", "/clusters/[ClusterId]/projects/[ProjectId]/events", "", "")
	req.Method = "GET"

	resp := struct {
		*responses.BaseResponse
		GetProjectEventsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetProjectEventsResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetProjectEventsResponse struct {
}
