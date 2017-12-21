package cs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeServiceContainersRequest struct {
	ClusterId string `position:"Path" name:"ClusterId"`
	ServiceId string `position:"Path" name:"ServiceId"`
}

func (r DescribeServiceContainersRequest) Invoke(client *sdk.Client) (response *DescribeServiceContainersResponse, err error) {
	req := struct {
		*requests.RoaRequest
		DescribeServiceContainersRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("CS", "2015-12-15", "DescribeServiceContainers", "/clusters/[ClusterId]/services/[ServiceId]/containers", "", "")
	req.Method = "GET"

	resp := struct {
		*responses.BaseResponse
		DescribeServiceContainersResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeServiceContainersResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeServiceContainersResponse struct {
}
