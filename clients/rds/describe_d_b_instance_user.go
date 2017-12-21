package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDBInstanceUserRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ConnectionString     string `position:"Query" name:"ConnectionString"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeDBInstanceUserRequest) Invoke(client *sdk.Client) (response *DescribeDBInstanceUserResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeDBInstanceUserRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeDBInstanceUser", "rds", "")

	resp := struct {
		*responses.BaseResponse
		DescribeDBInstanceUserResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeDBInstanceUserResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeDBInstanceUserResponse struct {
	RequestId      string
	DBInstanceName string
	InternalDBFlag string
}
