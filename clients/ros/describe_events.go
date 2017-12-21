package ros

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeEventsRequest struct {
	StackName      string `position:"Path" name:"StackName"`
	StackId        string `position:"Path" name:"StackId"`
	ResourceStatus string `position:"Query" name:"ResourceStatus"`
	ResourceName   string `position:"Query" name:"ResourceName"`
	ResourceType   string `position:"Query" name:"ResourceType"`
	PageSize       int    `position:"Query" name:"PageSize"`
	PageNumber     int    `position:"Query" name:"PageNumber"`
}

func (r DescribeEventsRequest) Invoke(client *sdk.Client) (response *DescribeEventsResponse, err error) {
	req := struct {
		*requests.RoaRequest
		DescribeEventsRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("ROS", "2015-09-01", "DescribeEvents", "/stacks/[StackName]/[StackId]/events", "", "")
	req.Method = "GET"

	resp := struct {
		*responses.BaseResponse
		DescribeEventsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeEventsResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeEventsResponse struct {
}
