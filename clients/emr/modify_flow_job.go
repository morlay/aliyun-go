package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyFlowJobRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	RunConf         string `position:"Query" name:"RunConf"`
	EnvConf         string `position:"Query" name:"EnvConf"`
	Description     string `position:"Query" name:"Description"`
	Params          string `position:"Query" name:"Params"`
	ParamConf       string `position:"Query" name:"ParamConf"`
	FailAct         string `position:"Query" name:"FailAct"`
	Mode            string `position:"Query" name:"Mode"`
	RetryInterval   int64  `position:"Query" name:"RetryInterval"`
	Name            string `position:"Query" name:"Name"`
	Id              string `position:"Query" name:"Id"`
	MaxRetry        int    `position:"Query" name:"MaxRetry"`
	ProjectId       string `position:"Query" name:"ProjectId"`
}

func (req *ModifyFlowJobRequest) Invoke(client *sdk.Client) (resp *ModifyFlowJobResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ModifyFlowJob", "", "")
	resp = &ModifyFlowJobResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyFlowJobResponse struct {
	responses.BaseResponse
	RequestId string
	Data      bool
}
