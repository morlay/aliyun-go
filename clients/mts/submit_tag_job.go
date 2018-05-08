package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SubmitTagJobRequest struct {
	requests.RpcRequest
	Input                string `position:"Query" name:"Input"`
	UserData             string `position:"Query" name:"UserData"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	TagConfig            string `position:"Query" name:"TagConfig"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PipelineId           string `position:"Query" name:"PipelineId"`
}

func (req *SubmitTagJobRequest) Invoke(client *sdk.Client) (resp *SubmitTagJobResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "SubmitTagJob", "mts", "")
	resp = &SubmitTagJobResponse{}
	err = client.DoAction(req, resp)
	return
}

type SubmitTagJobResponse struct {
	responses.BaseResponse
	RequestId common.String
	JobId     common.String
}
