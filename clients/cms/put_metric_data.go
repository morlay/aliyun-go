package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type PutMetricDataRequest struct {
	Callby_cms_owner string `position:"Query" name:"Callby_cms_owner"`
	Body             string `position:"Query" name:"Body"`
}

func (r PutMetricDataRequest) Invoke(client *sdk.Client) (response *PutMetricDataResponse, err error) {
	req := struct {
		*requests.RpcRequest
		PutMetricDataRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cms", "2017-03-01", "PutMetricData", "cms", "")

	resp := struct {
		*responses.BaseResponse
		PutMetricDataResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.PutMetricDataResponse

	err = client.DoAction(&req, &resp)
	return
}

type PutMetricDataResponse struct {
	Code      string
	Message   string
	RequestId string
	Success   bool
}
