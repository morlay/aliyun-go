package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyFlowRequest struct {
	requests.RpcRequest
	CronExpr        string `position:"Query" name:"CronExpr"`
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Periodic        string `position:"Query" name:"Periodic"`
	StartSchedule   int64  `position:"Query" name:"StartSchedule"`
	Description     string `position:"Query" name:"Description"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	Graph           string `position:"Query" name:"Graph"`
	CreateCluster   string `position:"Query" name:"CreateCluster"`
	Name            string `position:"Query" name:"Name"`
	EndSchedule     int64  `position:"Query" name:"EndSchedule"`
	Id              string `position:"Query" name:"Id"`
	ProjectId       string `position:"Query" name:"ProjectId"`
	Status          string `position:"Query" name:"Status"`
	ParentCategory  string `position:"Query" name:"ParentCategory"`
}

func (req *ModifyFlowRequest) Invoke(client *sdk.Client) (resp *ModifyFlowResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ModifyFlow", "", "")
	resp = &ModifyFlowResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyFlowResponse struct {
	responses.BaseResponse
	RequestId string
	Data      bool
}
