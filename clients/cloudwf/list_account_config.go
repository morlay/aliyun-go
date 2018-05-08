package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListAccountConfigRequest struct {
	requests.RpcRequest
	OrderCol    string `position:"Query" name:"OrderCol"`
	Length      int    `position:"Query" name:"Length"`
	SearchEmail string `position:"Query" name:"SearchEmail"`
	PageIndex   int    `position:"Query" name:"PageIndex"`
	OrderDir    string `position:"Query" name:"OrderDir"`
}

func (req *ListAccountConfigRequest) Invoke(client *sdk.Client) (resp *ListAccountConfigResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ListAccountConfig", "", "")
	resp = &ListAccountConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListAccountConfigResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
	Message   common.String
	Data      common.String
	ErrorCode common.Integer
	ErrorMsg  common.String
}
