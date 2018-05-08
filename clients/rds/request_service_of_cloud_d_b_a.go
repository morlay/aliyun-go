package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type RequestServiceOfCloudDBARequest struct {
	requests.RpcRequest
	ServiceRequestParam string `position:"Query" name:"ServiceRequestParam"`
	DBInstanceId        string `position:"Query" name:"DBInstanceId"`
	ServiceRequestType  string `position:"Query" name:"ServiceRequestType"`
}

func (req *RequestServiceOfCloudDBARequest) Invoke(client *sdk.Client) (resp *RequestServiceOfCloudDBAResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "RequestServiceOfCloudDBA", "rds", "")
	resp = &RequestServiceOfCloudDBAResponse{}
	err = client.DoAction(req, resp)
	return
}

type RequestServiceOfCloudDBAResponse struct {
	responses.BaseResponse
	RequestId common.String
	ListData  common.String
	AttrData  common.String
}
