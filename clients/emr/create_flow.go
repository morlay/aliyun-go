package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateFlowRequest struct {
	requests.RpcRequest
	CronExpr       string `position:"Query" name:"CronExpr"`
	StartSchedule  int64  `position:"Query" name:"StartSchedule"`
	Description    string `position:"Query" name:"Description"`
	ClusterId      string `position:"Query" name:"ClusterId"`
	Type           string `position:"Query" name:"Type"`
	Graph          string `position:"Query" name:"Graph"`
	CreateCluster  string `position:"Query" name:"CreateCluster"`
	Name           string `position:"Query" name:"Name"`
	EndSchedule    int64  `position:"Query" name:"EndSchedule"`
	ProjectId      string `position:"Query" name:"ProjectId"`
	ParentCategory string `position:"Query" name:"ParentCategory"`
}

func (req *CreateFlowRequest) Invoke(client *sdk.Client) (resp *CreateFlowResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "CreateFlow", "", "")
	resp = &CreateFlowResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateFlowResponse struct {
	responses.BaseResponse
	RequestId common.String
	Id        common.String
}
