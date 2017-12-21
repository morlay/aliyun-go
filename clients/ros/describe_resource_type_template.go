package ros

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeResourceTypeTemplateRequest struct {
	requests.RoaRequest
	TypeName string `position:"Path" name:"TypeName"`
}

func (req *DescribeResourceTypeTemplateRequest) Invoke(client *sdk.Client) (resp *DescribeResourceTypeTemplateResponse, err error) {
	req.InitWithApiInfo("ROS", "2015-09-01", "DescribeResourceTypeTemplate", "/resource_types/[TypeName]/template", "", "")
	req.Method = "GET"

	resp = &DescribeResourceTypeTemplateResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeResourceTypeTemplateResponse struct {
	responses.BaseResponse
}
