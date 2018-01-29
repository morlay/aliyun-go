package tesladam

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ActionDiskRmaRequest struct {
	requests.RpcRequest
	DiskName    string `position:"Query" name:"DiskName"`
	ExecutionId string `position:"Query" name:"ExecutionId"`
	DiskSlot    string `position:"Query" name:"DiskSlot"`
	Hostname    string `position:"Query" name:"Hostname"`
	DiskMount   string `position:"Query" name:"DiskMount"`
	DiskReason  string `position:"Query" name:"DiskReason"`
	DiskSn      string `position:"Query" name:"DiskSn"`
}

func (req *ActionDiskRmaRequest) Invoke(client *sdk.Client) (resp *ActionDiskRmaResponse, err error) {
	req.InitWithApiInfo("TeslaDam", "2018-01-18", "ActionDiskRma", "", "")
	resp = &ActionDiskRmaResponse{}
	err = client.DoAction(req, resp)
	return
}

type ActionDiskRmaResponse struct {
	responses.BaseResponse
	Status  bool
	Message string
	Result  string
}
