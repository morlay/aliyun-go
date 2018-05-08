package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ExportImageRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ImageId              string `position:"Query" name:"ImageId"`
	OSSBucket            string `position:"Query" name:"OSSBucket"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OSSPrefix            string `position:"Query" name:"OSSPrefix"`
	RoleName             string `position:"Query" name:"RoleName"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ImageFormat          string `position:"Query" name:"ImageFormat"`
}

func (req *ExportImageRequest) Invoke(client *sdk.Client) (resp *ExportImageResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "ExportImage", "ecs", "")
	resp = &ExportImageResponse{}
	err = client.DoAction(req, resp)
	return
}

type ExportImageResponse struct {
	responses.BaseResponse
	RequestId common.String
	TaskId    common.String
	RegionId  common.String
}
