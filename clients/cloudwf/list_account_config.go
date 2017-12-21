package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
