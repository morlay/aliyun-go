package yundun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteBackDoorFileRequest struct {
	requests.RpcRequest
	JstOwnerId int64  `position:"Query" name:"JstOwnerId"`
	InstanceId string `position:"Query" name:"InstanceId"`
	Path       string `position:"Query" name:"Path"`
}

func (req *DeleteBackDoorFileRequest) Invoke(client *sdk.Client) (resp *DeleteBackDoorFileResponse, err error) {
	req.InitWithApiInfo("Yundun", "2015-04-16", "DeleteBackDoorFile", "", "")
	resp = &DeleteBackDoorFileResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteBackDoorFileResponse struct {
	responses.BaseResponse
	RequestId string
}
