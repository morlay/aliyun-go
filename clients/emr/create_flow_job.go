package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateFlowJobRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	RunConf         string `position:"Query" name:"RunConf"`
	EnvConf         string `position:"Query" name:"EnvConf"`
	Description     string `position:"Query" name:"Description"`
	Type            string `position:"Query" name:"Type"`
	Params          string `position:"Query" name:"Params"`
	ParamConf       string `position:"Query" name:"ParamConf"`
	FailAct         string `position:"Query" name:"FailAct"`
	Mode            string `position:"Query" name:"Mode"`
	RetryInterval   int64  `position:"Query" name:"RetryInterval"`
	Name            string `position:"Query" name:"Name"`
	MaxRetry        int    `position:"Query" name:"MaxRetry"`
	ProjectId       string `position:"Query" name:"ProjectId"`
	ParentCategory  string `position:"Query" name:"ParentCategory"`
}

func (req *CreateFlowJobRequest) Invoke(client *sdk.Client) (resp *CreateFlowJobResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "CreateFlowJob", "", "")
	resp = &CreateFlowJobResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateFlowJobResponse struct {
	responses.BaseResponse
	RequestId common.String
	Id        common.String
}
