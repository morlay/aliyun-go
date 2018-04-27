package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryMetricListRequest struct {
	requests.RpcRequest
	Cursor           string `position:"Query" name:"Cursor"`
	Callby_cms_owner string `position:"Query" name:"Callby_cms_owner"`
	ResourceOwnerId  int64  `position:"Query" name:"ResourceOwnerId"`
	Period           string `position:"Query" name:"Period"`
	Length           string `position:"Query" name:"Length"`
	Project          string `position:"Query" name:"Project"`
	EndTime          string `position:"Query" name:"EndTime"`
	Express          string `position:"Query" name:"Express"`
	StartTime        string `position:"Query" name:"StartTime"`
	Metric           string `position:"Query" name:"Metric"`
	Page             string `position:"Query" name:"Page"`
	Dimensions       string `position:"Query" name:"Dimensions"`
}

func (req *QueryMetricListRequest) Invoke(client *sdk.Client) (resp *QueryMetricListResponse, err error) {
	req.InitWithApiInfo("Cms", "2018-03-08", "QueryMetricList", "cms", "")
	resp = &QueryMetricListResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryMetricListResponse struct {
	responses.BaseResponse
	Code       string
	Message    string
	RequestId  string
	Cursor     string
	Datapoints string
	Period     string
}
