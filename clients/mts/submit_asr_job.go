package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SubmitAsrJobRequest struct {
	Input                string `position:"Query" name:"Input"`
	UserData             string `position:"Query" name:"UserData"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	AsrConfig            string `position:"Query" name:"AsrConfig"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PipelineId           string `position:"Query" name:"PipelineId"`
}

func (r SubmitAsrJobRequest) Invoke(client *sdk.Client) (response *SubmitAsrJobResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SubmitAsrJobRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Mts", "2014-06-18", "SubmitAsrJob", "", "")

	resp := struct {
		*responses.BaseResponse
		SubmitAsrJobResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SubmitAsrJobResponse

	err = client.DoAction(&req, &resp)
	return
}

type SubmitAsrJobResponse struct {
	RequestId string
	JobId     string
}
