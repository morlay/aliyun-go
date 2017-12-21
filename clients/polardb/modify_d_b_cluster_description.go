package polardb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyDBClusterDescriptionRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	DBClusterDescription string `position:"Query" name:"DBClusterDescription"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string `position:"Query" name:"DBClusterId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r ModifyDBClusterDescriptionRequest) Invoke(client *sdk.Client) (response *ModifyDBClusterDescriptionResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyDBClusterDescriptionRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("polardb", "2017-08-01", "ModifyDBClusterDescription", "polardb", "")

	resp := struct {
		*responses.BaseResponse
		ModifyDBClusterDescriptionResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyDBClusterDescriptionResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyDBClusterDescriptionResponse struct {
	RequestId string
}
