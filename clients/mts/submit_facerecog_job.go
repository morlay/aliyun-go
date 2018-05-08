package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SubmitFacerecogJobRequest struct {
	requests.RpcRequest
	Input                string `position:"Query" name:"Input"`
	UserData             string `position:"Query" name:"UserData"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	FacerecogConfig      string `position:"Query" name:"FacerecogConfig"`
	PipelineId           string `position:"Query" name:"PipelineId"`
}

func (req *SubmitFacerecogJobRequest) Invoke(client *sdk.Client) (resp *SubmitFacerecogJobResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "SubmitFacerecogJob", "mts", "")
	resp = &SubmitFacerecogJobResponse{}
	err = client.DoAction(req, resp)
	return
}

type SubmitFacerecogJobResponse struct {
	responses.BaseResponse
	RequestId common.String
	JobId     common.String
}
