package cs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeClustersRequest struct {
	Name string `position:"Query" name:"Name"`
}

func (r DescribeClustersRequest) Invoke(client *sdk.Client) (response *DescribeClustersResponse, err error) {
	req := struct {
		*requests.RoaRequest
		DescribeClustersRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("CS", "2015-12-15", "DescribeClusters", "/clusters", "", "")
	req.Method = "GET"

	resp := struct {
		*responses.BaseResponse
		DescribeClustersResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeClustersResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeClustersResponse struct {
}
