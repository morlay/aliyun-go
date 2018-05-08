package phoenixsp_inner

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SyncPhoenixCfgRequest struct {
	requests.RpcRequest
	Data          string `position:"Query" name:"Data"`
	OperationType string `position:"Query" name:"OperationType"`
}

func (req *SyncPhoenixCfgRequest) Invoke(client *sdk.Client) (resp *SyncPhoenixCfgResponse, err error) {
	req.InitWithApiInfo("Phoenixsp-inner", "2016-08-05", "SyncPhoenixCfg", "", "")
	resp = &SyncPhoenixCfgResponse{}
	err = client.DoAction(req, resp)
	return
}

type SyncPhoenixCfgResponse struct {
	responses.BaseResponse
	RequestId common.String
	Code      common.String
	Message   common.String
	Success   bool
	Data      common.String
}
