package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyJobRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
	Name            string `position:"Query" name:"Name"`
	Type            string `position:"Query" name:"Type"`
	RunParameter    string `position:"Query" name:"RunParameter"`
	FailAct         string `position:"Query" name:"FailAct"`
}

func (r ModifyJobRequest) Invoke(client *sdk.Client) (response *ModifyJobResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyJobRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Emr", "2016-04-08", "ModifyJob", "", "")

	resp := struct {
		*responses.BaseResponse
		ModifyJobResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyJobResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyJobResponse struct {
	RequestId string
}
