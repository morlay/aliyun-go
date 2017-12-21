package cs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeClustersRequest struct {
	requests.RoaRequest
	Name string `position:"Query" name:"Name"`
}

func (req *DescribeClustersRequest) Invoke(client *sdk.Client) (resp *DescribeClustersResponse, err error) {
	req.InitWithApiInfo("CS", "2015-12-15", "DescribeClusters", "/clusters", "", "")
	req.Method = "GET"

	resp = &DescribeClustersResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeClustersResponse struct {
	responses.BaseResponse
}
