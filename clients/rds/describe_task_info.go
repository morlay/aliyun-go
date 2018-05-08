package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeTaskInfoRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	TaskId               int64  `position:"Query" name:"TaskId"`
}

func (req *DescribeTaskInfoRequest) Invoke(client *sdk.Client) (resp *DescribeTaskInfoResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeTaskInfo", "rds", "")
	resp = &DescribeTaskInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeTaskInfoResponse struct {
	responses.BaseResponse
	RequestId          common.String
	TaskId             common.String
	BeginTime          common.String
	FinishTime         common.String
	CreateTime         common.String
	TaskAction         common.String
	DBName             common.String
	TaskErrorCode      common.String
	Progress           common.String
	ExpectedFinishTime common.String
	TaskErrorMessage   common.String
	ProgressInfo       common.String
	Status             common.String
}
