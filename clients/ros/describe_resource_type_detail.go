package ros

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeResourceTypeDetailRequest struct {
	requests.RoaRequest
	TypeName string `position:"Path" name:"TypeName"`
}

func (req *DescribeResourceTypeDetailRequest) Invoke(client *sdk.Client) (resp *DescribeResourceTypeDetailResponse, err error) {
	req.InitWithApiInfo("ROS", "2015-09-01", "DescribeResourceTypeDetail", "/resource_types/[TypeName]", "", "")
	req.Method = "GET"

	resp = &DescribeResourceTypeDetailResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeResourceTypeDetailResponse struct {
	responses.BaseResponse
}
