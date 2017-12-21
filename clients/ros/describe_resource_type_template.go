package ros

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeResourceTypeTemplateRequest struct {
	TypeName string `position:"Path" name:"TypeName"`
}

func (r DescribeResourceTypeTemplateRequest) Invoke(client *sdk.Client) (response *DescribeResourceTypeTemplateResponse, err error) {
	req := struct {
		*requests.RoaRequest
		DescribeResourceTypeTemplateRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("ROS", "2015-09-01", "DescribeResourceTypeTemplate", "/resource_types/[TypeName]/template", "", "")
	req.Method = "GET"

	resp := struct {
		*responses.BaseResponse
		DescribeResourceTypeTemplateResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeResourceTypeTemplateResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeResourceTypeTemplateResponse struct {
}
