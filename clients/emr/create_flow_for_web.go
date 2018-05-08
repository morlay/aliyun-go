package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateFlowForWebRequest struct {
	requests.RpcRequest
	CronExpr       string `position:"Query" name:"CronExpr"`
	StartSchedule  int64  `position:"Query" name:"StartSchedule"`
	Name           string `position:"Query" name:"Name"`
	Description    string `position:"Query" name:"Description"`
	EndSchedule    int64  `position:"Query" name:"EndSchedule"`
	ClusterId      string `position:"Query" name:"ClusterId"`
	ProjectId      string `position:"Query" name:"ProjectId"`
	Graph          string `position:"Query" name:"Graph"`
	ParentCategory string `position:"Query" name:"ParentCategory"`
}

func (req *CreateFlowForWebRequest) Invoke(client *sdk.Client) (resp *CreateFlowForWebResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "CreateFlowForWeb", "", "")
	resp = &CreateFlowForWebResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateFlowForWebResponse struct {
	responses.BaseResponse
	RequestId common.String
	Id        common.String
}
