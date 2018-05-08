package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeISPAreaCityRequest struct {
	requests.RpcRequest
}

func (req *DescribeISPAreaCityRequest) Invoke(client *sdk.Client) (resp *DescribeISPAreaCityResponse, err error) {
	req.InitWithApiInfo("Cms", "2018-03-08", "DescribeISPAreaCity", "cms", "")
	resp = &DescribeISPAreaCityResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeISPAreaCityResponse struct {
	responses.BaseResponse
	Code      common.String
	Message   common.String
	Success   common.String
	RequestId common.String
	Data      common.String
}
