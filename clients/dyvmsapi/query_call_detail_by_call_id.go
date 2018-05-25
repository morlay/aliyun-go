package dyvmsapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type QueryCallDetailByCallIdRequest struct {
	requests.RpcRequest
	CallId               string `position:"Query" name:"CallId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	QueryDate            int64  `position:"Query" name:"QueryDate"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ProdId               int64  `position:"Query" name:"ProdId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *QueryCallDetailByCallIdRequest) Invoke(client *sdk.Client) (resp *QueryCallDetailByCallIdResponse, err error) {
	req.InitWithApiInfo("Dyvmsapi", "2017-05-25", "QueryCallDetailByCallId", "", "")
	resp = &QueryCallDetailByCallIdResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryCallDetailByCallIdResponse struct {
	responses.BaseResponse
	RequestId common.String
	Data      common.String
	Code      common.String
	Message   common.String
}
