package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetCasterConfigRequest struct {
	requests.RpcRequest
	SideOutputUrl    string  `position:"Query" name:"SideOutputUrl"`
	CasterId         string  `position:"Query" name:"CasterId"`
	DomainName       string  `position:"Query" name:"DomainName"`
	ProgramEffect    int     `position:"Query" name:"ProgramEffect"`
	ProgramName      string  `position:"Query" name:"ProgramName"`
	OwnerId          int64   `position:"Query" name:"OwnerId"`
	RecordConfig     string  `position:"Query" name:"RecordConfig"`
	UrgentMaterialId string  `position:"Query" name:"UrgentMaterialId"`
	TranscodeConfig  string  `position:"Query" name:"TranscodeConfig"`
	Delay            float32 `position:"Query" name:"Delay"`
	CasterName       string  `position:"Query" name:"CasterName"`
	CallbackUrl      string  `position:"Query" name:"CallbackUrl"`
}

func (req *SetCasterConfigRequest) Invoke(client *sdk.Client) (resp *SetCasterConfigResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "SetCasterConfig", "live", "")
	resp = &SetCasterConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetCasterConfigResponse struct {
	responses.BaseResponse
	RequestId string
	CasterId  string
}
