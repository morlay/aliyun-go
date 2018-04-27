package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryMetricDataRequest struct {
	requests.RpcRequest
	Callby_cms_owner string `position:"Query" name:"Callby_cms_owner"`
	ResourceOwnerId  int64  `position:"Query" name:"ResourceOwnerId"`
	Period           string `position:"Query" name:"Period"`
	Metric           string `position:"Query" name:"Metric"`
	Length           string `position:"Query" name:"Length"`
	Project          string `position:"Query" name:"Project"`
	EndTime          string `position:"Query" name:"EndTime"`
	Express          string `position:"Query" name:"Express"`
	StartTime        string `position:"Query" name:"StartTime"`
	Dimensions       string `position:"Query" name:"Dimensions"`
}

func (req *QueryMetricDataRequest) Invoke(client *sdk.Client) (resp *QueryMetricDataResponse, err error) {
	req.InitWithApiInfo("Cms", "2018-03-08", "QueryMetricData", "cms", "")
	resp = &QueryMetricDataResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryMetricDataResponse struct {
	responses.BaseResponse
	Code       string
	Message    string
	RequestId  string
	Datapoints string
	Period     string
}
