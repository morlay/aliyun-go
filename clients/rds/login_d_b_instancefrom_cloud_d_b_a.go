package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type LoginDBInstancefromCloudDBARequest struct {
	requests.RpcRequest
	ServiceRequestParam string `position:"Query" name:"ServiceRequestParam"`
	DBInstanceId        string `position:"Query" name:"DBInstanceId"`
	ServiceRequestType  string `position:"Query" name:"ServiceRequestType"`
}

func (req *LoginDBInstancefromCloudDBARequest) Invoke(client *sdk.Client) (resp *LoginDBInstancefromCloudDBAResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "LoginDBInstancefromCloudDBA", "rds", "")
	resp = &LoginDBInstancefromCloudDBAResponse{}
	err = client.DoAction(req, resp)
	return
}

type LoginDBInstancefromCloudDBAResponse struct {
	responses.BaseResponse
	RequestId common.String
	ListData  common.String
	AttrData  common.String
}
