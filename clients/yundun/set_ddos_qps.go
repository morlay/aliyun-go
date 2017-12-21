package yundun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetDdosQpsRequest struct {
	requests.RpcRequest
	InstanceId   string `position:"Query" name:"InstanceId"`
	InstanceType string `position:"Query" name:"InstanceType"`
	QpsPosition  int    `position:"Query" name:"QpsPosition"`
	Level        int    `position:"Query" name:"Level"`
}

func (req *SetDdosQpsRequest) Invoke(client *sdk.Client) (resp *SetDdosQpsResponse, err error) {
	req.InitWithApiInfo("Yundun", "2015-04-16", "SetDdosQps", "", "")
	resp = &SetDdosQpsResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetDdosQpsResponse struct {
	responses.BaseResponse
	RequestId string
}
