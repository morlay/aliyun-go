package cloudphoto

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteEventRequest struct {
	requests.RpcRequest
	EventId   int64  `position:"Query" name:"EventId"`
	LibraryId string `position:"Query" name:"LibraryId"`
	StoreName string `position:"Query" name:"StoreName"`
}

func (req *DeleteEventRequest) Invoke(client *sdk.Client) (resp *DeleteEventResponse, err error) {
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "DeleteEvent", "cloudphoto", "")
	resp = &DeleteEventResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteEventResponse struct {
	responses.BaseResponse
	Code      common.String
	Message   common.String
	RequestId common.String
	Action    common.String
}
