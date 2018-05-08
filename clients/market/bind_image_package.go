package market

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type BindImagePackageRequest struct {
	requests.RpcRequest
	EcsInstanceId          string `position:"Query" name:"EcsInstanceId"`
	ImagePackageInstanceId string `position:"Query" name:"ImagePackageInstanceId"`
}

func (req *BindImagePackageRequest) Invoke(client *sdk.Client) (resp *BindImagePackageResponse, err error) {
	req.InitWithApiInfo("Market", "2015-11-01", "BindImagePackage", "yunmarket", "")
	resp = &BindImagePackageResponse{}
	err = client.DoAction(req, resp)
	return
}

type BindImagePackageResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
}
