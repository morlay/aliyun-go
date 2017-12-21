package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DisableAlarmRequest struct {
	Callby_cms_owner string `position:"Query" name:"Callby_cms_owner"`
	Id               string `position:"Query" name:"Id"`
}

func (r DisableAlarmRequest) Invoke(client *sdk.Client) (response *DisableAlarmResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DisableAlarmRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cms", "2017-03-01", "DisableAlarm", "cms", "")

	resp := struct {
		*responses.BaseResponse
		DisableAlarmResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DisableAlarmResponse

	err = client.DoAction(&req, &resp)
	return
}

type DisableAlarmResponse struct {
	Success   bool
	Code      string
	Message   string
	RequestId string
}
