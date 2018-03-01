package cs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeImagesRequest struct {
	requests.RoaRequest
	ImageName     string `position:"Query" name:"ImageName"`
	DockerVersion string `position:"Query" name:"DockerVersion"`
}

func (req *DescribeImagesRequest) Invoke(client *sdk.Client) (resp *DescribeImagesResponse, err error) {
	req.InitWithApiInfo("CS", "2015-12-15", "DescribeImages", "/images", "", "")
	req.Method = "GET"

	resp = &DescribeImagesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeImagesResponse struct {
	responses.BaseResponse
}
