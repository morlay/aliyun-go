package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteLogConfigRequest struct {
	requests.RpcRequest
	LogType string `position:"Query" name:"LogType"`
}

func (req *DeleteLogConfigRequest) Invoke(client *sdk.Client) (resp *DeleteLogConfigResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DeleteLogConfig", "apigateway", "")
	resp = &DeleteLogConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteLogConfigResponse struct {
	responses.BaseResponse
	RequestId common.String
}
