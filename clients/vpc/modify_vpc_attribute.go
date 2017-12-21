package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyVpcAttributeRequest struct {
	VpcName              string `position:"Query" name:"VpcName"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	VpcId                string `position:"Query" name:"VpcId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Description          string `position:"Query" name:"Description"`
	UserCidr             string `position:"Query" name:"UserCidr"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r ModifyVpcAttributeRequest) Invoke(client *sdk.Client) (response *ModifyVpcAttributeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyVpcAttributeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Vpc", "2016-04-28", "ModifyVpcAttribute", "vpc", "")

	resp := struct {
		*responses.BaseResponse
		ModifyVpcAttributeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyVpcAttributeResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyVpcAttributeResponse struct {
	RequestId string
}
