package csb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CommitSuccessedServicesRequest struct {
	requests.RpcRequest
	CsbName  string `position:"Query" name:"CsbName"`
	Services string `position:"Body" name:"Services"`
}

func (req *CommitSuccessedServicesRequest) Invoke(client *sdk.Client) (resp *CommitSuccessedServicesResponse, err error) {
	req.InitWithApiInfo("CSB", "2017-11-18", "CommitSuccessedServices", "CSB", "")
	resp = &CommitSuccessedServicesResponse{}
	err = client.DoAction(req, resp)
	return
}

type CommitSuccessedServicesResponse struct {
	responses.BaseResponse
	Code      int
	Message   string
	RequestId string
}
