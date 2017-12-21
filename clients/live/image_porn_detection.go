package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ImagePornDetectionRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	ImageUrl      string `position:"Query" name:"ImageUrl"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r ImagePornDetectionRequest) Invoke(client *sdk.Client) (response *ImagePornDetectionResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ImagePornDetectionRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("live", "2016-11-01", "ImagePornDetection", "", "")

	resp := struct {
		*responses.BaseResponse
		ImagePornDetectionResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ImagePornDetectionResponse

	err = client.DoAction(&req, &resp)
	return
}

type ImagePornDetectionResponse struct {
	RequestId string
	Label     string
	Rate      float32
}
