package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteAppRequest struct {
	requests.RpcRequest
	AppId int64 `position:"Query" name:"AppId"`
}

func (req *DeleteAppRequest) Invoke(client *sdk.Client) (resp *DeleteAppResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DeleteApp", "apigateway", "")
	resp = &DeleteAppResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteAppResponse struct {
	responses.BaseResponse
	RequestId common.String
}
