package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	Code      string
	Message   string
	Success   string
	RequestId string
	Data      string
}
