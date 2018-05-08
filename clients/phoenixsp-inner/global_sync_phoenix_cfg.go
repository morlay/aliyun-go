package phoenixsp_inner

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GlobalSyncPhoenixCfgRequest struct {
	requests.RpcRequest
	Data          string `position:"Query" name:"Data"`
	OperationType string `position:"Query" name:"OperationType"`
}

func (req *GlobalSyncPhoenixCfgRequest) Invoke(client *sdk.Client) (resp *GlobalSyncPhoenixCfgResponse, err error) {
	req.InitWithApiInfo("Phoenixsp-inner", "2016-08-05", "GlobalSyncPhoenixCfg", "", "")
	resp = &GlobalSyncPhoenixCfgResponse{}
	err = client.DoAction(req, resp)
	return
}

type GlobalSyncPhoenixCfgResponse struct {
	responses.BaseResponse
	RequestId common.String
	Code      common.String
	Message   common.String
	Success   bool
	Data      common.String
}
