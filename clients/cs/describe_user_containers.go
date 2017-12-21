package cs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeUserContainersRequest struct {
	requests.RoaRequest
	ServiceId string `position:"Query" name:"ServiceId"`
}

func (req *DescribeUserContainersRequest) Invoke(client *sdk.Client) (resp *DescribeUserContainersResponse, err error) {
	req.InitWithApiInfo("CS", "2015-12-15", "DescribeUserContainers", "/region/[RegionId]/containers", "", "")
	req.Method = "GET"

	resp = &DescribeUserContainersResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeUserContainersResponse struct {
	responses.BaseResponse
}
