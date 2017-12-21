package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type LoginDBInstancefromCloudDBARequest struct {
	ServiceRequestParam string `position:"Query" name:"ServiceRequestParam"`
	DBInstanceId        string `position:"Query" name:"DBInstanceId"`
	ServiceRequestType  string `position:"Query" name:"ServiceRequestType"`
}

func (r LoginDBInstancefromCloudDBARequest) Invoke(client *sdk.Client) (response *LoginDBInstancefromCloudDBAResponse, err error) {
	req := struct {
		*requests.RpcRequest
		LoginDBInstancefromCloudDBARequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "LoginDBInstancefromCloudDBA", "rds", "")

	resp := struct {
		*responses.BaseResponse
		LoginDBInstancefromCloudDBAResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.LoginDBInstancefromCloudDBAResponse

	err = client.DoAction(&req, &resp)
	return
}

type LoginDBInstancefromCloudDBAResponse struct {
	RequestId string
	ListData  string
	AttrData  string
}
