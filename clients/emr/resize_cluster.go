package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ResizeClusterRequest struct {
	requests.RpcRequest
	NewMasterInstances int    `position:"Query" name:"NewMasterInstances"`
	Period             int    `position:"Query" name:"Period"`
	AutoRenew          string `position:"Query" name:"AutoRenew"`
	CoreInstanceType   string `position:"Query" name:"CoreInstanceType"`
	NewCoreInstances   int    `position:"Query" name:"NewCoreInstances"`
	NewTaskInstances   int    `position:"Query" name:"NewTaskInstances"`
	ClusterId          string `position:"Query" name:"ClusterId"`
	ChargeType         string `position:"Query" name:"ChargeType"`
}

func (req *ResizeClusterRequest) Invoke(client *sdk.Client) (resp *ResizeClusterResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ResizeCluster", "", "")
	resp = &ResizeClusterResponse{}
	err = client.DoAction(req, resp)
	return
}

type ResizeClusterResponse struct {
	responses.BaseResponse
	RequestId   string
	ClusterId   string
	EmrOrderId  string
	CoreOrderId string
}
