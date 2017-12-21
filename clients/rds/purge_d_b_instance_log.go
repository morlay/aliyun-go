package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type PurgeDBInstanceLogRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r PurgeDBInstanceLogRequest) Invoke(client *sdk.Client) (response *PurgeDBInstanceLogResponse, err error) {
	req := struct {
		*requests.RpcRequest
		PurgeDBInstanceLogRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "PurgeDBInstanceLog", "rds", "")

	resp := struct {
		*responses.BaseResponse
		PurgeDBInstanceLogResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.PurgeDBInstanceLogResponse

	err = client.DoAction(&req, &resp)
	return
}

type PurgeDBInstanceLogResponse struct {
	RequestId string
}
