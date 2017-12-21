package cs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeServiceContainersRequest struct {
	requests.RoaRequest
	ClusterId string `position:"Path" name:"ClusterId"`
	ServiceId string `position:"Path" name:"ServiceId"`
}

func (req *DescribeServiceContainersRequest) Invoke(client *sdk.Client) (resp *DescribeServiceContainersResponse, err error) {
	req.InitWithApiInfo("CS", "2015-12-15", "DescribeServiceContainers", "/clusters/[ClusterId]/services/[ServiceId]/containers", "", "")
	req.Method = "GET"

	resp = &DescribeServiceContainersResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeServiceContainersResponse struct {
	responses.BaseResponse
}
