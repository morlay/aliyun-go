package yundun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetDdosQpsRequest struct {
	InstanceId   string `position:"Query" name:"InstanceId"`
	InstanceType string `position:"Query" name:"InstanceType"`
	QpsPosition  int    `position:"Query" name:"QpsPosition"`
	Level        int    `position:"Query" name:"Level"`
}

func (r SetDdosQpsRequest) Invoke(client *sdk.Client) (response *SetDdosQpsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SetDdosQpsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Yundun", "2015-04-16", "SetDdosQps", "", "")

	resp := struct {
		*responses.BaseResponse
		SetDdosQpsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SetDdosQpsResponse

	err = client.DoAction(&req, &resp)
	return
}

type SetDdosQpsResponse struct {
	RequestId string
}
