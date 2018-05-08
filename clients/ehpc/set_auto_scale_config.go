package ehpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SetAutoScaleConfigRequest struct {
	requests.RpcRequest
	ShrinkIdleTimes         int    `position:"Query" name:"ShrinkIdleTimes"`
	GrowTimeoutInMinutes    int    `position:"Query" name:"GrowTimeoutInMinutes"`
	ClusterId               string `position:"Query" name:"ClusterId"`
	EnableAutoGrow          string `position:"Query" name:"EnableAutoGrow"`
	EnableAutoShrink        string `position:"Query" name:"EnableAutoShrink"`
	MaxNodesInCluster       int    `position:"Query" name:"MaxNodesInCluster"`
	ExcludeNodes            string `position:"Query" name:"ExcludeNodes"`
	ShrinkIntervalInMinutes int    `position:"Query" name:"ShrinkIntervalInMinutes"`
	ExtraNodesGrowRatio     int    `position:"Query" name:"ExtraNodesGrowRatio"`
	GrowIntervalInMinutes   int    `position:"Query" name:"GrowIntervalInMinutes"`
	GrowRatio               int    `position:"Query" name:"GrowRatio"`
}

func (req *SetAutoScaleConfigRequest) Invoke(client *sdk.Client) (resp *SetAutoScaleConfigResponse, err error) {
	req.InitWithApiInfo("EHPC", "2017-07-14", "SetAutoScaleConfig", "ehs", "")
	resp = &SetAutoScaleConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetAutoScaleConfigResponse struct {
	responses.BaseResponse
	RequestId common.String
}
