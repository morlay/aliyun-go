package ehpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetAutoScaleConfigRequest struct {
	requests.RpcRequest
	ClusterId string `position:"Query" name:"ClusterId"`
}

func (req *GetAutoScaleConfigRequest) Invoke(client *sdk.Client) (resp *GetAutoScaleConfigResponse, err error) {
	req.InitWithApiInfo("EHPC", "2017-07-14", "GetAutoScaleConfig", "ehs", "")
	resp = &GetAutoScaleConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetAutoScaleConfigResponse struct {
	responses.BaseResponse
	RequestId               string
	Uid                     string
	ClusterId               string
	ClusterType             string
	EnableAutoGrow          bool
	EnableAutoShrink        bool
	GrowIntervalInMinutes   int
	ShrinkIntervalInMinutes int
	ShrinkIdleTimes         int
	GrowTimeoutInMinutes    int
	ExtraNodesGrowRatio     int
	GrowRatio               int
	MaxNodesInCluster       int
	ExcludeNodes            string
}
