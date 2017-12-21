package slb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetLogsDownloadAttributeRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	LogType              string `position:"Query" name:"LogType"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	RoleName             string `position:"Query" name:"RoleName"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OSSBucketName        string `position:"Query" name:"OSSBucketName"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
}

func (r SetLogsDownloadAttributeRequest) Invoke(client *sdk.Client) (response *SetLogsDownloadAttributeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SetLogsDownloadAttributeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Slb", "2014-05-15", "SetLogsDownloadAttribute", "slb", "")

	resp := struct {
		*responses.BaseResponse
		SetLogsDownloadAttributeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SetLogsDownloadAttributeResponse

	err = client.DoAction(&req, &resp)
	return
}

type SetLogsDownloadAttributeResponse struct {
	RequestId string
}
