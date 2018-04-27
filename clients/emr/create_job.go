package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateJobRequest struct {
	requests.RpcRequest
	RunParameter    string `position:"Query" name:"RunParameter"`
	RetryInterval   int    `position:"Query" name:"RetryInterval"`
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Name            string `position:"Query" name:"Name"`
	Type            string `position:"Query" name:"Type"`
	MaxRetry        int    `position:"Query" name:"MaxRetry"`
	FailAct         string `position:"Query" name:"FailAct"`
}

func (req *CreateJobRequest) Invoke(client *sdk.Client) (resp *CreateJobResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "CreateJob", "", "")
	resp = &CreateJobResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateJobResponse struct {
	responses.BaseResponse
	RequestId string
	Id        string
}
