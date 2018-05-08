package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type KillExecutionJobInstanceRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	JobInstanceId   string `position:"Query" name:"JobInstanceId"`
}

func (req *KillExecutionJobInstanceRequest) Invoke(client *sdk.Client) (resp *KillExecutionJobInstanceResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "KillExecutionJobInstance", "", "")
	resp = &KillExecutionJobInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type KillExecutionJobInstanceResponse struct {
	responses.BaseResponse
	RequestId common.String
}
