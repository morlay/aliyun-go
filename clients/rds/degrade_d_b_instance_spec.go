package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DegradeDBInstanceSpecRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	DBInstanceStorage    int    `position:"Query" name:"DBInstanceStorage"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	DBInstanceClass      string `position:"Query" name:"DBInstanceClass"`
}

func (r DegradeDBInstanceSpecRequest) Invoke(client *sdk.Client) (response *DegradeDBInstanceSpecResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DegradeDBInstanceSpecRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "DegradeDBInstanceSpec", "rds", "")

	resp := struct {
		*responses.BaseResponse
		DegradeDBInstanceSpecResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DegradeDBInstanceSpecResponse

	err = client.DoAction(&req, &resp)
	return
}

type DegradeDBInstanceSpecResponse struct {
	RequestId string
}
