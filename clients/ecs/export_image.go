package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ExportImageRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ImageId              string `position:"Query" name:"ImageId"`
	OSSBucket            string `position:"Query" name:"OSSBucket"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OSSPrefix            string `position:"Query" name:"OSSPrefix"`
	RoleName             string `position:"Query" name:"RoleName"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ImageFormat          string `position:"Query" name:"ImageFormat"`
}

func (r ExportImageRequest) Invoke(client *sdk.Client) (response *ExportImageResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ExportImageRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "ExportImage", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		ExportImageResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ExportImageResponse

	err = client.DoAction(&req, &resp)
	return
}

type ExportImageResponse struct {
	RequestId string
	TaskId    string
	RegionId  string
}
