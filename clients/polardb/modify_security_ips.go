package polardb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifySecurityIpsRequest struct {
	ResourceOwnerId           int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount      string `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId               string `position:"Query" name:"DBClusterId"`
	OwnerAccount              string `position:"Query" name:"OwnerAccount"`
	SecurityIps               string `position:"Query" name:"SecurityIps"`
	DBClusterIPArrayName      string `position:"Query" name:"DBClusterIPArrayName"`
	OwnerId                   int64  `position:"Query" name:"OwnerId"`
	DBClusterIPArrayAttribute string `position:"Query" name:"DBClusterIPArrayAttribute"`
}

func (r ModifySecurityIpsRequest) Invoke(client *sdk.Client) (response *ModifySecurityIpsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifySecurityIpsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("polardb", "2017-08-01", "ModifySecurityIps", "polardb", "")

	resp := struct {
		*responses.BaseResponse
		ModifySecurityIpsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifySecurityIpsResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifySecurityIpsResponse struct {
	RequestId string
}
