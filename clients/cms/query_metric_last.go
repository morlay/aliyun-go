package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type QueryMetricLastRequest struct {
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

func (req *QueryMetricLastRequest) Invoke(client *sdk.Client) (resp *QueryMetricLastResponse, err error) {
	req.InitWithApiInfo("Cms", "2018-03-08", "QueryMetricLast", "cms", "")
	resp = &QueryMetricLastResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryMetricLastResponse struct {
	responses.BaseResponse
	Code       common.String
	Message    common.String
	RequestId  common.String
	Cursor     common.String
	Datapoints common.String
	Period     common.String
}
