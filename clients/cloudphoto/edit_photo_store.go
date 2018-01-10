package cloudphoto

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type EditPhotoStoreRequest struct {
	requests.RpcRequest
	AutoCleanEnabled  string `position:"Query" name:"AutoCleanEnabled"`
	DefaultTrashQuota int64  `position:"Query" name:"DefaultTrashQuota"`
	StoreName         string `position:"Query" name:"StoreName"`
	Remark            string `position:"Query" name:"Remark"`
	DefaultQuota      int64  `position:"Query" name:"DefaultQuota"`
	AutoCleanDays     int    `position:"Query" name:"AutoCleanDays"`
}

func (req *EditPhotoStoreRequest) Invoke(client *sdk.Client) (resp *EditPhotoStoreResponse, err error) {
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "EditPhotoStore", "cloudphoto", "")
	resp = &EditPhotoStoreResponse{}
	err = client.DoAction(req, resp)
	return
}

type EditPhotoStoreResponse struct {
	responses.BaseResponse
	Code      string
	Message   string
	RequestId string
	Action    string
}
