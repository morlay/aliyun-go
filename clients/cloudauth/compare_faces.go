package cloudauth

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CompareFacesRequest struct {
	requests.RpcRequest
	SourceImageType  string `position:"Query" name:"SourceImageType"`
	ResourceOwnerId  int64  `position:"Query" name:"ResourceOwnerId"`
	TargetImageType  string `position:"Query" name:"TargetImageType"`
	SourceImageValue string `position:"Query" name:"SourceImageValue"`
	TargetImageValue string `position:"Query" name:"TargetImageValue"`
}

func (req *CompareFacesRequest) Invoke(client *sdk.Client) (resp *CompareFacesResponse, err error) {
	req.InitWithApiInfo("Cloudauth", "2017-11-17", "CompareFaces", "cloudauth", "")
	resp = &CompareFacesResponse{}
	err = client.DoAction(req, resp)
	return
}

type CompareFacesResponse struct {
	responses.BaseResponse
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
