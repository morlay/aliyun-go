package yundun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetDdosAutoRequest struct {
	InstanceId   string `position:"Query" name:"InstanceId"`
	InstanceType string `position:"Query" name:"InstanceType"`
}

func (r SetDdosAutoRequest) Invoke(client *sdk.Client) (response *SetDdosAutoResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SetDdosAutoRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Yundun", "2015-04-16", "SetDdosAuto", "", "")

	resp := struct {
		*responses.BaseResponse
		SetDdosAutoResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SetDdosAutoResponse

	err = client.DoAction(&req, &resp)
	return
}

type SetDdosAutoResponse struct {
	RequestId string
}
