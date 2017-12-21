package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type MigrateToOtherZoneRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ZoneId               string `position:"Query" name:"ZoneId"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r MigrateToOtherZoneRequest) Invoke(client *sdk.Client) (response *MigrateToOtherZoneResponse, err error) {
	req := struct {
		*requests.RpcRequest
		MigrateToOtherZoneRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "MigrateToOtherZone", "rds", "")

	resp := struct {
		*responses.BaseResponse
		MigrateToOtherZoneResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.MigrateToOtherZoneResponse

	err = client.DoAction(&req, &resp)
	return
}

type MigrateToOtherZoneResponse struct {
	RequestId string
}
