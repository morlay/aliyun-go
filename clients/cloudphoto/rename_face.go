package cloudphoto

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type RenameFaceRequest struct {
	requests.RpcRequest
	LibraryId string `position:"Query" name:"LibraryId"`
	StoreName string `position:"Query" name:"StoreName"`
	FaceId    int64  `position:"Query" name:"FaceId"`
	FaceName  string `position:"Query" name:"FaceName"`
}

func (req *RenameFaceRequest) Invoke(client *sdk.Client) (resp *RenameFaceResponse, err error) {
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "RenameFace", "cloudphoto", "")
	resp = &RenameFaceResponse{}
	err = client.DoAction(req, resp)
	return
}

type RenameFaceResponse struct {
	responses.BaseResponse
	Code      common.String
	Message   common.String
	RequestId common.String
	Action    common.String
}
