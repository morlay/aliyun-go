package polardb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyDBInstanceDescriptionRequest struct {
	ResourceOwnerId       int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount          string `position:"Query" name:"OwnerAccount"`
	DBInstanceId          string `position:"Query" name:"DBInstanceId"`
	DBInstanceDescription string `position:"Query" name:"DBInstanceDescription"`
	OwnerId               int64  `position:"Query" name:"OwnerId"`
}

func (r ModifyDBInstanceDescriptionRequest) Invoke(client *sdk.Client) (response *ModifyDBInstanceDescriptionResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyDBInstanceDescriptionRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("polardb", "2017-08-01", "ModifyDBInstanceDescription", "polardb", "")

	resp := struct {
		*responses.BaseResponse
		ModifyDBInstanceDescriptionResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyDBInstanceDescriptionResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyDBInstanceDescriptionResponse struct {
	RequestId string
}
