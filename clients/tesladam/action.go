package tesladam

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ActionRequest struct {
	requests.RpcRequest
	OrderId  int    `position:"Query" name:"OrderId"`
	StepCode string `position:"Query" name:"StepCode"`
}

func (req *ActionRequest) Invoke(client *sdk.Client) (resp *ActionResponse, err error) {
	req.InitWithApiInfo("TeslaDam", "2018-01-18", "Action", "", "")
	resp = &ActionResponse{}
	err = client.DoAction(req, resp)
	return
}

type ActionResponse struct {
	responses.BaseResponse
	Status  bool
	Message string
	Result  string
}
