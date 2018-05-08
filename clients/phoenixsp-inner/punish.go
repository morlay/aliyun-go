package phoenixsp_inner

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId common.String
	Code      common.String
	Message   common.String
	Success   bool
	Data      common.String
}
