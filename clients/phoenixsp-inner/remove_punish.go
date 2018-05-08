package phoenixsp_inner

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type RemovePunishRequest struct {
	requests.RpcRequest
	BizPunishRequest string `position:"Query" name:"BizPunishRequest"`
	Sign             string `position:"Query" name:"Sign"`
	PunishResult     string `position:"Query" name:"PunishResult"`
}

func (req *RemovePunishRequest) Invoke(client *sdk.Client) (resp *RemovePunishResponse, err error) {
	req.InitWithApiInfo("Phoenixsp-inner", "2016-08-05", "RemovePunish", "", "")
	resp = &RemovePunishResponse{}
	err = client.DoAction(req, resp)
	return
}

type RemovePunishResponse struct {
	responses.BaseResponse
	RequestId common.String
	Code      common.String
	Message   common.String
	Success   bool
	Data      common.String
}
