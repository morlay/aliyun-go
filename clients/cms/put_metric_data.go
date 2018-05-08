package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type PutMetricDataRequest struct {
	requests.RpcRequest
	Callby_cms_owner string `position:"Query" name:"Callby_cms_owner"`
	Body             string `position:"Query" name:"Body"`
}

func (req *PutMetricDataRequest) Invoke(client *sdk.Client) (resp *PutMetricDataResponse, err error) {
	req.InitWithApiInfo("Cms", "2018-03-08", "PutMetricData", "cms", "")
	resp = &PutMetricDataResponse{}
	err = client.DoAction(req, resp)
	return
}

type PutMetricDataResponse struct {
	responses.BaseResponse
	Code      common.String
	Message   common.String
	RequestId common.String
	Success   bool
}
