package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ImagePornDetectionRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	ImageUrl      string `position:"Query" name:"ImageUrl"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *ImagePornDetectionRequest) Invoke(client *sdk.Client) (resp *ImagePornDetectionResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "ImagePornDetection", "live", "")
	resp = &ImagePornDetectionResponse{}
	err = client.DoAction(req, resp)
	return
}

type ImagePornDetectionResponse struct {
	responses.BaseResponse
	RequestId common.String
	Label     common.String
	Rate      common.Float
}
