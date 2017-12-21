package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type PutMetricDataRequest struct {
	requests.RpcRequest
	Callby_cms_owner string `position:"Query" name:"Callby_cms_owner"`
	Body             string `position:"Query" name:"Body"`
}

func (req *PutMetricDataRequest) Invoke(client *sdk.Client) (resp *PutMetricDataResponse, err error) {
	req.InitWithApiInfo("Cms", "2017-03-01", "PutMetricData", "cms", "")
	resp = &PutMetricDataResponse{}
	err = client.DoAction(req, resp)
	return
}

type PutMetricDataResponse struct {
	responses.BaseResponse
	Code      string
	Message   string
	RequestId string
	Success   bool
}
