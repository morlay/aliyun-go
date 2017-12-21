package polardb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeTaskInfoRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	TaskId               string `position:"Query" name:"TaskId"`
}

func (req *DescribeTaskInfoRequest) Invoke(client *sdk.Client) (resp *DescribeTaskInfoResponse, err error) {
	req.InitWithApiInfo("polardb", "2017-08-01", "DescribeTaskInfo", "polardb", "")
	resp = &DescribeTaskInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeTaskInfoResponse struct {
	responses.BaseResponse
	RequestId          string
	TaskId             string
	BeginTime          string
	FinishTime         string
	ExpectedFinishTime string
	TaskAction         string
	Progress           int
	DBName             string
	ProgressInfo       string
}
