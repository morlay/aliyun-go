package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CancelJobRequest struct {
	JobId                string `position:"Query" name:"JobId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r CancelJobRequest) Invoke(client *sdk.Client) (response *CancelJobResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CancelJobRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Mts", "2014-06-18", "CancelJob", "", "")

	resp := struct {
		*responses.BaseResponse
		CancelJobResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CancelJobResponse

	err = client.DoAction(&req, &resp)
	return
}

type CancelJobResponse struct {
	RequestId string
	JobId     string
}
