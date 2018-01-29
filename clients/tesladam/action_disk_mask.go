package tesladam

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ActionDiskMaskRequest struct {
	requests.RpcRequest
	Op        string `position:"Query" name:"Op"`
	DiskMount string `position:"Query" name:"DiskMount"`
	Ip        string `position:"Query" name:"Ip"`
}

func (req *ActionDiskMaskRequest) Invoke(client *sdk.Client) (resp *ActionDiskMaskResponse, err error) {
	req.InitWithApiInfo("TeslaDam", "2018-01-18", "ActionDiskMask", "", "")
	resp = &ActionDiskMaskResponse{}
	err = client.DoAction(req, resp)
	return
}

type ActionDiskMaskResponse struct {
	responses.BaseResponse
	Status  bool
	Message string
	Result  string
}
