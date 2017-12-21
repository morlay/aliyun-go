package dds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RestartDBInstanceRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	NodeId               string `position:"Query" name:"NodeId"`
}

func (r RestartDBInstanceRequest) Invoke(client *sdk.Client) (response *RestartDBInstanceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		RestartDBInstanceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Dds", "2015-12-01", "RestartDBInstance", "dds", "")

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
