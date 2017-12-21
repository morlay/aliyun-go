package cloudauth

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CompareFacesRequest struct {
	SourceImageType  string `position:"Query" name:"SourceImageType"`
	ResourceOwnerId  int64  `position:"Query" name:"ResourceOwnerId"`
	TargetImageType  string `position:"Query" name:"TargetImageType"`
	SourceImageValue string `position:"Query" name:"SourceImageValue"`
	TargetImageValue string `position:"Query" name:"TargetImageValue"`
}

func (r CompareFacesRequest) Invoke(client *sdk.Client) (response *CompareFacesResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CompareFacesRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cloudauth", "2017-11-17", "CompareFaces", "cloudauth", "")

	resp := struct {
		*responses.BaseResponse
		CompareFacesResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CompareFacesResponse

	err = client.DoAction(&req, &resp)
	return
}

type CompareFacesResponse struct {
	RequestId string
	Success   bool
	Code      string
	Message   string
	Data      CompareFacesData
}

type CompareFacesData struct {
	SimilarityScore      float32
	ConfidenceThresholds string
}
