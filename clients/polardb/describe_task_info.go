package polardb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeTaskInfoRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	TaskId               string `position:"Query" name:"TaskId"`
}

func (r DescribeTaskInfoRequest) Invoke(client *sdk.Client) (response *DescribeTaskInfoResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeTaskInfoRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("polardb", "2017-08-01", "DescribeTaskInfo", "polardb", "")

	resp := struct {
		*responses.BaseResponse
		DescribeTaskInfoResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeTaskInfoResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeTaskInfoResponse struct {
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
