package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyJobRequest struct {
	requests.RpcRequest
	RunParameter    string `position:"Query" name:"RunParameter"`
	RetryInterval   int    `position:"Query" name:"RetryInterval"`
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Name            string `position:"Query" name:"Name"`
	Id              string `position:"Query" name:"Id"`
	Type            string `position:"Query" name:"Type"`
	MaxRetry        int    `position:"Query" name:"MaxRetry"`
	FailAct         string `position:"Query" name:"FailAct"`
}

func (req *ModifyJobRequest) Invoke(client *sdk.Client) (resp *ModifyJobResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ModifyJob", "", "")
	resp = &ModifyJobResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyJobResponse struct {
	responses.BaseResponse
	RequestId common.String
}
