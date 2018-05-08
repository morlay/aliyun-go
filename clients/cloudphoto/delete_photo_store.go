package cloudphoto

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeletePhotoStoreRequest struct {
	requests.RpcRequest
	StoreName string `position:"Query" name:"StoreName"`
}

func (req *DeletePhotoStoreRequest) Invoke(client *sdk.Client) (resp *DeletePhotoStoreResponse, err error) {
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "DeletePhotoStore", "cloudphoto", "")
	resp = &DeletePhotoStoreResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeletePhotoStoreResponse struct {
	responses.BaseResponse
	Code      common.String
	Message   common.String
	RequestId common.String
	Action    common.String
}
