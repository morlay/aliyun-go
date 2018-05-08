package aegis

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type UpgradeInstanceRequest struct {
	requests.RpcRequest
	InstanceId  string `position:"Query" name:"InstanceId"`
	ClientToken string `position:"Query" name:"ClientToken"`
	VmNumber    int    `position:"Query" name:"VmNumber"`
	OwnerId     int64  `position:"Query" name:"OwnerId"`
	VersionCode int    `position:"Query" name:"VersionCode"`
}

func (req *UpgradeInstanceRequest) Invoke(client *sdk.Client) (resp *UpgradeInstanceResponse, err error) {
	req.InitWithApiInfo("aegis", "2016-11-11", "UpgradeInstance", "vipaegis", "")
	resp = &UpgradeInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpgradeInstanceResponse struct {
	responses.BaseResponse
	OrderId   common.String
	RequestId common.String
}
