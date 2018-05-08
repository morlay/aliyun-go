package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type RemoveVpcAccessRequest struct {
	requests.RpcRequest
	VpcId      string `position:"Query" name:"VpcId"`
	InstanceId string `position:"Query" name:"InstanceId"`
	Port       int    `position:"Query" name:"Port"`
}

func (req *RemoveVpcAccessRequest) Invoke(client *sdk.Client) (resp *RemoveVpcAccessResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "RemoveVpcAccess", "apigateway", "")
	resp = &RemoveVpcAccessResponse{}
	err = client.DoAction(req, resp)
	return
}

type RemoveVpcAccessResponse struct {
	responses.BaseResponse
	RequestId common.String
}
