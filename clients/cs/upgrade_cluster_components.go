package cs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpgradeClusterComponentsRequest struct {
	requests.RoaRequest
	ComponentId string `position:"Path" name:"ComponentId"`
	ClusterId   string `position:"Path" name:"ClusterId"`
}

func (req *UpgradeClusterComponentsRequest) Invoke(client *sdk.Client) (resp *UpgradeClusterComponentsResponse, err error) {
	req.InitWithApiInfo("CS", "2015-12-15", "UpgradeClusterComponents", "/clusters/[ClusterId]/components/[ComponentId]/upgrade", "", "")
	req.Method = "POST"

	resp = &UpgradeClusterComponentsResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpgradeClusterComponentsResponse struct {
	responses.BaseResponse
}
