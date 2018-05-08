package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeTaskDetailRequest struct {
	requests.RpcRequest
	TaskId string `position:"Query" name:"TaskId"`
}

func (req *DescribeTaskDetailRequest) Invoke(client *sdk.Client) (resp *DescribeTaskDetailResponse, err error) {
	req.InitWithApiInfo("Cms", "2018-03-08", "DescribeTaskDetail", "cms", "")
	resp = &DescribeTaskDetailResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeTaskDetailResponse struct {
	responses.BaseResponse
	Code      common.String
	Message   common.String
	Success   common.String
	RequestId common.String
	Data      common.String
}
