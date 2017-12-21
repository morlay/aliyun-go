package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RequestServiceOfCloudDBARequest struct {
	ServiceRequestParam string `position:"Query" name:"ServiceRequestParam"`
	DBInstanceId        string `position:"Query" name:"DBInstanceId"`
	ServiceRequestType  string `position:"Query" name:"ServiceRequestType"`
}

func (r RequestServiceOfCloudDBARequest) Invoke(client *sdk.Client) (response *RequestServiceOfCloudDBAResponse, err error) {
	req := struct {
		*requests.RpcRequest
		RequestServiceOfCloudDBARequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "RequestServiceOfCloudDBA", "rds", "")

	resp := struct {
		*responses.BaseResponse
		RequestServiceOfCloudDBAResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.RequestServiceOfCloudDBAResponse

	err = client.DoAction(&req, &resp)
	return
}

type RequestServiceOfCloudDBAResponse struct {
	RequestId string
	ListData  string
	AttrData  string
}
