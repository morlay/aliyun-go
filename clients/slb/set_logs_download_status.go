package slb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetLogsDownloadStatusRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	RoleName             string `position:"Query" name:"RoleName"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	LogsDownloadStatus   string `position:"Query" name:"LogsDownloadStatus"`
	Tags                 string `position:"Query" name:"Tags"`
}

func (r SetLogsDownloadStatusRequest) Invoke(client *sdk.Client) (response *SetLogsDownloadStatusResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SetLogsDownloadStatusRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Slb", "2014-05-15", "SetLogsDownloadStatus", "slb", "")

	resp := struct {
		*responses.BaseResponse
		SetLogsDownloadStatusResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SetLogsDownloadStatusResponse

	err = client.DoAction(&req, &resp)
	return
}

type SetLogsDownloadStatusResponse struct {
	RequestId string
}
