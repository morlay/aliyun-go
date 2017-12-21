package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type EnableAlarmRequest struct {
	Callby_cms_owner string `position:"Query" name:"Callby_cms_owner"`
	Id               string `position:"Query" name:"Id"`
}

func (r EnableAlarmRequest) Invoke(client *sdk.Client) (response *EnableAlarmResponse, err error) {
	req := struct {
		*requests.RpcRequest
		EnableAlarmRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cms", "2017-03-01", "EnableAlarm", "cms", "")

	resp := struct {
		*responses.BaseResponse
		EnableAlarmResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.EnableAlarmResponse

	err = client.DoAction(&req, &resp)
	return
}

type EnableAlarmResponse struct {
	Success   bool
	Code      string
	Message   string
	RequestId string
}
