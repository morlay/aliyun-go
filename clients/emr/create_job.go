package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateJobRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Name            string `position:"Query" name:"Name"`
	Type            string `position:"Query" name:"Type"`
	RunParameter    string `position:"Query" name:"RunParameter"`
	FailAct         string `position:"Query" name:"FailAct"`
}

func (r CreateJobRequest) Invoke(client *sdk.Client) (response *CreateJobResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateJobRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Emr", "2016-04-08", "CreateJob", "", "")

	resp := struct {
		*responses.BaseResponse
		CreateJobResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateJobResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateJobResponse struct {
	RequestId string
	Id        string
}
