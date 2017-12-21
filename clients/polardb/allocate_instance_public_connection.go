package polardb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AllocateInstancePublicConnectionRequest struct {
	ResourceOwnerId        int64  `position:"Query" name:"ResourceOwnerId"`
	ConnectionStringPrefix string `position:"Query" name:"ConnectionStringPrefix"`
	ResourceOwnerAccount   string `position:"Query" name:"ResourceOwnerAccount"`
	Port                   string `position:"Query" name:"Port"`
	OwnerAccount           string `position:"Query" name:"OwnerAccount"`
	DBInstanceId           string `position:"Query" name:"DBInstanceId"`
	OwnerId                int64  `position:"Query" name:"OwnerId"`
}

func (r AllocateInstancePublicConnectionRequest) Invoke(client *sdk.Client) (response *AllocateInstancePublicConnectionResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AllocateInstancePublicConnectionRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("polardb", "2017-08-01", "AllocateInstancePublicConnection", "polardb", "")

	resp := struct {
		*responses.BaseResponse
		AllocateInstancePublicConnectionResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AllocateInstancePublicConnectionResponse

	err = client.DoAction(&req, &resp)
	return
}

type AllocateInstancePublicConnectionResponse struct {
	RequestId        string
	DBInstanceId     string
	ConnectionString string
	IPType           string
	Port             string
}
