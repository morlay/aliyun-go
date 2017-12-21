package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeTaskInfoRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	TaskId               int64  `position:"Query" name:"TaskId"`
}

func (r DescribeTaskInfoRequest) Invoke(client *sdk.Client) (response *DescribeTaskInfoResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeTaskInfoRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeTaskInfo", "rds", "")

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
	CreateTime         string
	TaskAction         string
	DBName             string
	TaskErrorCode      string
	Progress           string
	ExpectedFinishTime string
	TaskErrorMessage   string
	ProgressInfo       string
	Status             string
}
