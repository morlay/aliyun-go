package cs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeUserContainersRequest struct {
	ServiceId string `position:"Query" name:"ServiceId"`
}

func (r DescribeUserContainersRequest) Invoke(client *sdk.Client) (response *DescribeUserContainersResponse, err error) {
	req := struct {
		*requests.RoaRequest
		DescribeUserContainersRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("CS", "2015-12-15", "DescribeUserContainers", "/region/[RegionId]/containers", "", "")
	req.Method = "GET"

	resp := struct {
		*responses.BaseResponse
		DescribeUserContainersResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeUserContainersResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeUserContainersResponse struct {
}
