package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteVpcRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	VpcId                string `position:"Query" name:"VpcId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DeleteVpcRequest) Invoke(client *sdk.Client) (response *DeleteVpcResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteVpcRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Vpc", "2016-04-28", "DeleteVpc", "vpc", "")

	resp := struct {
		*responses.BaseResponse
		DeleteVpcResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteVpcResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteVpcResponse struct {
	RequestId string
}
