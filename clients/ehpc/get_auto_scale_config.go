package ehpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetAutoScaleConfigRequest struct {
	requests.RpcRequest
	ClusterId string `position:"Query" name:"ClusterId"`
}

func (req *GetAutoScaleConfigRequest) Invoke(client *sdk.Client) (resp *GetAutoScaleConfigResponse, err error) {
	req.InitWithApiInfo("EHPC", "2018-04-12", "GetAutoScaleConfig", "ehs", "")
	resp = &GetAutoScaleConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetAutoScaleConfigResponse struct {
	responses.BaseResponse
	RequestId               common.String
	Uid                     common.String
	ClusterId               common.String
	ClusterType             common.String
	EnableAutoGrow          bool
	EnableAutoShrink        bool
	GrowIntervalInMinutes   common.Integer
	ShrinkIntervalInMinutes common.Integer
	ShrinkIdleTimes         common.Integer
	GrowTimeoutInMinutes    common.Integer
	ExtraNodesGrowRatio     common.Integer
	GrowRatio               common.Integer
	MaxNodesInCluster       common.Integer
	ExcludeNodes            common.String
}
