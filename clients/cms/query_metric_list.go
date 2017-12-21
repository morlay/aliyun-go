package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryMetricListRequest struct {
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

func (r QueryMetricListRequest) Invoke(client *sdk.Client) (response *QueryMetricListResponse, err error) {
	req := struct {
		*requests.RpcRequest
		QueryMetricListRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cms", "2017-03-01", "QueryMetricList", "cms", "")

	resp := struct {
		*responses.BaseResponse
		QueryMetricListResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.QueryMetricListResponse

	err = client.DoAction(&req, &resp)
	return
}

type QueryMetricListResponse struct {
	Code       string
	Message    string
	Success    bool
	RequestId  string
	Cursor     string
	Datapoints string
	Period     string
}
