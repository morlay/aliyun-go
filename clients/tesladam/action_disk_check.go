package tesladam

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ActionDiskCheckRequest struct {
	requests.RpcRequest
	DiskMount string `position:"Query" name:"DiskMount"`
	Ip        string `position:"Query" name:"Ip"`
}

func (req *ActionDiskCheckRequest) Invoke(client *sdk.Client) (resp *ActionDiskCheckResponse, err error) {
	req.InitWithApiInfo("TeslaDam", "2018-01-18", "ActionDiskCheck", "", "")
	resp = &ActionDiskCheckResponse{}
	err = client.DoAction(req, resp)
	return
}

type ActionDiskCheckResponse struct {
	responses.BaseResponse
	Status  bool
	Message common.String
	Result  common.String
}
