package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteJobRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
}

func (r DeleteJobRequest) Invoke(client *sdk.Client) (response *DeleteJobResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteJobRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Emr", "2016-04-08", "DeleteJob", "", "")

	resp := struct {
		*responses.BaseResponse
		DeleteJobResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteJobResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteJobResponse struct {
	RequestId string
}
