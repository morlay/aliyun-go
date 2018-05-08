package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ReportFacerecogJobResultRequest struct {
	requests.RpcRequest
	JobId                string `position:"Query" name:"JobId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Facerecog            string `position:"Query" name:"Facerecog"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Details              string `position:"Query" name:"Details"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ReportFacerecogJobResultRequest) Invoke(client *sdk.Client) (resp *ReportFacerecogJobResultResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "ReportFacerecogJobResult", "mts", "")
	resp = &ReportFacerecogJobResultResponse{}
	err = client.DoAction(req, resp)
	return
}

type ReportFacerecogJobResultResponse struct {
	responses.BaseResponse
	RequestId common.String
	JobId     common.String
}
