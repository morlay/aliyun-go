package ehpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpgradeClientRequest struct {
	requests.RpcRequest
	ClientVersion string `position:"Query" name:"ClientVersion"`
	ClusterId     string `position:"Query" name:"ClusterId"`
}

func (req *UpgradeClientRequest) Invoke(client *sdk.Client) (resp *UpgradeClientResponse, err error) {
	req.InitWithApiInfo("EHPC", "2017-07-14", "UpgradeClient", "ehs", "")
	resp = &UpgradeClientResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpgradeClientResponse struct {
	responses.BaseResponse
	RequestId string
}
