package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetProbeDataSubscriberConfigRequest struct {
	Id int64 `position:"Query" name:"Id"`
}

func (r GetProbeDataSubscriberConfigRequest) Invoke(client *sdk.Client) (response *GetProbeDataSubscriberConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetProbeDataSubscriberConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "GetProbeDataSubscriberConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		GetProbeDataSubscriberConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetProbeDataSubscriberConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetProbeDataSubscriberConfigResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
