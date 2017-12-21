package cloudphoto

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetMeRequest struct {
	requests.RpcRequest
	LibraryId string `position:"Query" name:"LibraryId"`
	StoreName string `position:"Query" name:"StoreName"`
	FaceId    int64  `position:"Query" name:"FaceId"`
}

func (req *SetMeRequest) Invoke(client *sdk.Client) (resp *SetMeResponse, err error) {
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "SetMe", "cloudphoto", "")
	resp = &SetMeResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetMeResponse struct {
	responses.BaseResponse
	Code      string
	Message   string
	RequestId string
	Action    string
}
