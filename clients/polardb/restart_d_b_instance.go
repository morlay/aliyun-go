package polardb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RestartDBInstanceRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r RestartDBInstanceRequest) Invoke(client *sdk.Client) (response *RestartDBInstanceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		RestartDBInstanceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("polardb", "2017-08-01", "RestartDBInstance", "polardb", "")

	resp := struct {
		*responses.BaseResponse
		RestartDBInstanceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.RestartDBInstanceResponse

	err = client.DoAction(&req, &resp)
	return
}

type RestartDBInstanceResponse struct {
	RequestId string
}
