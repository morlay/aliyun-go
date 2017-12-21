package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ReleaseReadWriteSplittingConnectionRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r ReleaseReadWriteSplittingConnectionRequest) Invoke(client *sdk.Client) (response *ReleaseReadWriteSplittingConnectionResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ReleaseReadWriteSplittingConnectionRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "ReleaseReadWriteSplittingConnection", "rds", "")

	resp := struct {
		*responses.BaseResponse
		ReleaseReadWriteSplittingConnectionResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ReleaseReadWriteSplittingConnectionResponse

	err = client.DoAction(&req, &resp)
	return
}

type ReleaseReadWriteSplittingConnectionResponse struct {
	RequestId string
}
