package phoenixsp_inner

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type PunishRequest struct {
	requests.RpcRequest
	BizPunishRequest string `position:"Query" name:"BizPunishRequest"`
	Sign             string `position:"Query" name:"Sign"`
}

func (req *PunishRequest) Invoke(client *sdk.Client) (resp *PunishResponse, err error) {
	req.InitWithApiInfo("Phoenixsp-inner", "2016-08-05", "Punish", "", "")
	resp = &PunishResponse{}
	err = client.DoAction(req, resp)
	return
}

type PunishResponse struct {
	responses.BaseResponse
	RequestId string
	Code      string
	Message   string
	Success   bool
	Data      string
}
