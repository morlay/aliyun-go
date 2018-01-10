package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SubmitMediaDetailJobRequest struct {
	requests.RpcRequest
	Input                string `position:"Query" name:"Input"`
	UserData             string `position:"Query" name:"UserData"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	MediaDetailConfig    string `position:"Query" name:"MediaDetailConfig"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PipelineId           string `position:"Query" name:"PipelineId"`
}

func (req *SubmitMediaDetailJobRequest) Invoke(client *sdk.Client) (resp *SubmitMediaDetailJobResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "SubmitMediaDetailJob", "", "")
	resp = &SubmitMediaDetailJobResponse{}
	err = client.DoAction(req, resp)
	return
}

type SubmitMediaDetailJobResponse struct {
	responses.BaseResponse
	RequestId string
	JobId     string
}
