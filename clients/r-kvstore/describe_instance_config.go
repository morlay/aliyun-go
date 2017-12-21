
package r-kvstore

import (
    "github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
)


type DescribeInstanceConfigRequest struct {
ResourceOwnerId int64 `position:"Query" name:"ResourceOwnerId"`
InstanceId string `position:"Query" name:"InstanceId"`
SecurityToken string `position:"Query" name:"SecurityToken"`
ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
OwnerAccount string `position:"Query" name:"OwnerAccount"`
OwnerId int64 `position:"Query" name:"OwnerId"`
}

func (r DescribeInstanceConfigRequest) Invoke(client *sdk.Client) (response *DescribeInstanceConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeInstanceConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "DescribeInstanceConfig", "redisa", "")

	resp := struct {
		*responses.BaseResponse
		DescribeInstanceConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
    response = &resp.DescribeInstanceConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeInstanceConfigResponse struct {
RequestId string
Config string
}

