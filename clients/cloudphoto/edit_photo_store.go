package cloudphoto

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type EditPhotoStoreRequest struct {
	AutoCleanEnabled string `position:"Query" name:"AutoCleanEnabled"`
	StoreName        string `position:"Query" name:"StoreName"`
	Remark           string `position:"Query" name:"Remark"`
	DefaultQuota     int64  `position:"Query" name:"DefaultQuota"`
	AutoCleanDays    int    `position:"Query" name:"AutoCleanDays"`
}

func (r EditPhotoStoreRequest) Invoke(client *sdk.Client) (response *EditPhotoStoreResponse, err error) {
	req := struct {
		*requests.RpcRequest
		EditPhotoStoreRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "EditPhotoStore", "cloudphoto", "")

	resp := struct {
		*responses.BaseResponse
		EditPhotoStoreResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.EditPhotoStoreResponse

	err = client.DoAction(&req, &resp)
	return
}

type EditPhotoStoreResponse struct {
	Code      string
	Message   string
	RequestId string
	Action    string
}
