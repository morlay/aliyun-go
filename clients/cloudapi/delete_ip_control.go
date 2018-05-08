package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteIpControlRequest struct {
	requests.RpcRequest
	IpControlId string `position:"Query" name:"IpControlId"`
}

func (req *DeleteIpControlRequest) Invoke(client *sdk.Client) (resp *DeleteIpControlResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DeleteIpControl", "apigateway", "")
	resp = &DeleteIpControlResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteIpControlResponse struct {
	responses.BaseResponse
	RequestId common.String
}
