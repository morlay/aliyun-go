package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryMetricTopRequest struct {
	requests.RpcRequest
	Callby_cms_owner string `position:"Query" name:"Callby_cms_owner"`
	Period           string `position:"Query" name:"Period"`
	ResourceOwnerId  int64  `position:"Query" name:"ResourceOwnerId"`
	Length           string `position:"Query" name:"Length"`
	Project          string `position:"Query" name:"Project"`
	EndTime          string `position:"Query" name:"EndTime"`
	Orderby          string `position:"Query" name:"Orderby"`
	Express          string `position:"Query" name:"Express"`
	StartTime        string `position:"Query" name:"StartTime"`
	OrderDesc        string `position:"Query" name:"OrderDesc"`
	Metric           string `position:"Query" name:"Metric"`
	Dimensions       string `position:"Query" name:"Dimensions"`
}

func (req *QueryMetricTopRequest) Invoke(client *sdk.Client) (resp *QueryMetricTopResponse, err error) {
	req.InitWithApiInfo("Cms", "2018-03-08", "QueryMetricTop", "cms", "")
	resp = &QueryMetricTopResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryMetricTopResponse struct {
	responses.BaseResponse
	Code       string
	Message    string
	RequestId  string
	Datapoints string
	Period     string
}
