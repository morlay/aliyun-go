package cloudphoto

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SetFaceCoverRequest struct {
	requests.RpcRequest
	LibraryId string `position:"Query" name:"LibraryId"`
	PhotoId   int64  `position:"Query" name:"PhotoId"`
	StoreName string `position:"Query" name:"StoreName"`
	FaceId    int64  `position:"Query" name:"FaceId"`
}

func (req *SetFaceCoverRequest) Invoke(client *sdk.Client) (resp *SetFaceCoverResponse, err error) {
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "SetFaceCover", "cloudphoto", "")
	resp = &SetFaceCoverResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetFaceCoverResponse struct {
	responses.BaseResponse
	Code      common.String
	Message   common.String
	RequestId common.String
	Action    common.String
}
