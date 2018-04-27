package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RunClusterServiceActionRequest struct {
	requests.RpcRequest
	ResourceOwnerId             int64  `position:"Query" name:"ResourceOwnerId"`
	OnlyRestartStaleConfigNodes string `position:"Query" name:"OnlyRestartStaleConfigNodes"`
	NodeCountPerBatch           int    `position:"Query" name:"NodeCountPerBatch"`
	ClusterId                   string `position:"Query" name:"ClusterId"`
	CustomCommand               string `position:"Query" name:"CustomCommand"`
	ComponentNameList           string `position:"Query" name:"ComponentNameList"`
	ServiceActionName           string `position:"Query" name:"ServiceActionName"`
	IsRolling                   string `position:"Query" name:"IsRolling"`
	TotlerateFailCount          int    `position:"Query" name:"TotlerateFailCount"`
	ServiceName                 string `position:"Query" name:"ServiceName"`
	Comment                     string `position:"Query" name:"Comment"`
	HostIdList                  string `position:"Query" name:"HostIdList"`
	TurnOnMaintenanceMode       string `position:"Query" name:"TurnOnMaintenanceMode"`
}

func (req *RunClusterServiceActionRequest) Invoke(client *sdk.Client) (resp *RunClusterServiceActionResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "RunClusterServiceAction", "", "")
	resp = &RunClusterServiceActionResponse{}
	err = client.DoAction(req, resp)
	return
}

type RunClusterServiceActionResponse struct {
	responses.BaseResponse
	RequestId string
}
