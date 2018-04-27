package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetWaitingRoomConfigRequest struct {
	requests.RpcRequest
	WaitUrl       string `position:"Query" name:"WaitUrl"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	WaitUri       string `position:"Query" name:"WaitUri"`
	MaxQps        int64  `position:"Query" name:"MaxQps"`
	MaxTimeWait   int    `position:"Query" name:"MaxTimeWait"`
	DomainName    string `position:"Query" name:"DomainName"`
	AllowPct      int    `position:"Query" name:"AllowPct"`
	GapTime       int    `position:"Query" name:"GapTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Version       string `position:"Query" name:"Version"`
}

func (req *SetWaitingRoomConfigRequest) Invoke(client *sdk.Client) (resp *SetWaitingRoomConfigResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "SetWaitingRoomConfig", "", "")
	resp = &SetWaitingRoomConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetWaitingRoomConfigResponse struct {
	responses.BaseResponse
	RequestId string
}
