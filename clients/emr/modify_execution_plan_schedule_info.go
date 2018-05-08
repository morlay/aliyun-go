package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyExecutionPlanScheduleInfoRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	TimeInterval    int    `position:"Query" name:"TimeInterval"`
	DayOfWeek       string `position:"Query" name:"DayOfWeek"`
	Id              string `position:"Query" name:"Id"`
	StartTime       int64  `position:"Query" name:"StartTime"`
	Strategy        string `position:"Query" name:"Strategy"`
	TimeUnit        string `position:"Query" name:"TimeUnit"`
	DayOfMonth      string `position:"Query" name:"DayOfMonth"`
}

func (req *ModifyExecutionPlanScheduleInfoRequest) Invoke(client *sdk.Client) (resp *ModifyExecutionPlanScheduleInfoResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ModifyExecutionPlanScheduleInfo", "", "")
	resp = &ModifyExecutionPlanScheduleInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyExecutionPlanScheduleInfoResponse struct {
	responses.BaseResponse
	RequestId common.String
}
