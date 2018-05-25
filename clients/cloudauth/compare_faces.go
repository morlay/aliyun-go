package cloudauth

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CompareFacesRequest struct {
	requests.RpcRequest
	SourceImageType  string `position:"Query" name:"SourceImageType"`
	ResourceOwnerId  int64  `position:"Query" name:"ResourceOwnerId"`
	SourceIp         string `position:"Query" name:"SourceIp"`
	TargetImageType  string `position:"Query" name:"TargetImageType"`
	SourceImageValue string `position:"Query" name:"SourceImageValue"`
	TargetImageValue string `position:"Query" name:"TargetImageValue"`
}

func (req *CompareFacesRequest) Invoke(client *sdk.Client) (resp *CompareFacesResponse, err error) {
	req.InitWithApiInfo("Cloudauth", "2018-05-04", "CompareFaces", "cloudauth", "")
	resp = &CompareFacesResponse{}
	err = client.DoAction(req, resp)
	return
}

type CompareFacesResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
	Code      common.String
	Message   common.String
	Data      CompareFacesData
}

type CompareFacesData struct {
	SimilarityScore      common.Float
	ConfidenceThresholds common.String
}
