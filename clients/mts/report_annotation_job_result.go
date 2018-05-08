package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ReportAnnotationJobResultRequest struct {
	requests.RpcRequest
	Annotation           string `position:"Query" name:"Annotation"`
	JobId                string `position:"Query" name:"JobId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Details              string `position:"Query" name:"Details"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ReportAnnotationJobResultRequest) Invoke(client *sdk.Client) (resp *ReportAnnotationJobResultResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "ReportAnnotationJobResult", "mts", "")
	resp = &ReportAnnotationJobResultResponse{}
	err = client.DoAction(req, resp)
	return
}

type ReportAnnotationJobResultResponse struct {
	responses.BaseResponse
	RequestId common.String
	JobId     common.String
}
