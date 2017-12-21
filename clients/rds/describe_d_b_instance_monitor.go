package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDBInstanceMonitorRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeDBInstanceMonitorRequest) Invoke(client *sdk.Client) (response *DescribeDBInstanceMonitorResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeDBInstanceMonitorRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeDBInstanceMonitor", "rds", "")

	resp := struct {
		*responses.BaseResponse
		DescribeDBInstanceMonitorResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeDBInstanceMonitorResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeDBInstanceMonitorResponse struct {
	RequestId string
	Period    string
}
