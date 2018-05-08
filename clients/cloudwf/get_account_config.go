package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetAccountConfigRequest struct {
	requests.RpcRequest
	Id int64 `position:"Query" name:"Id"`
}

func (req *GetAccountConfigRequest) Invoke(client *sdk.Client) (resp *GetAccountConfigResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "GetAccountConfig", "", "")
	resp = &GetAccountConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetAccountConfigResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
	Message   common.String
	Data      common.String
	ErrorCode common.Integer
	ErrorMsg  common.String
}
