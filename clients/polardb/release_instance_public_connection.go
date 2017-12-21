package polardb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ReleaseInstancePublicConnectionRequest struct {
	ResourceOwnerId         int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount    string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount            string `position:"Query" name:"OwnerAccount"`
	DBInstanceId            string `position:"Query" name:"DBInstanceId"`
	OwnerId                 int64  `position:"Query" name:"OwnerId"`
	CurrentConnectionString string `position:"Query" name:"CurrentConnectionString"`
}

func (r ReleaseInstancePublicConnectionRequest) Invoke(client *sdk.Client) (response *ReleaseInstancePublicConnectionResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ReleaseInstancePublicConnectionRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("polardb", "2017-08-01", "ReleaseInstancePublicConnection", "polardb", "")

	resp := struct {
		*responses.BaseResponse
		ReleaseInstancePublicConnectionResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ReleaseInstancePublicConnectionResponse

	err = client.DoAction(&req, &resp)
	return
}

type ReleaseInstancePublicConnectionResponse struct {
	RequestId string
}
