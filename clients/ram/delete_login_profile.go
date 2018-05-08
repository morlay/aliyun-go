package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteLoginProfileRequest struct {
	requests.RpcRequest
	UserName string `position:"Query" name:"UserName"`
}

func (req *DeleteLoginProfileRequest) Invoke(client *sdk.Client) (resp *DeleteLoginProfileResponse, err error) {
	req.InitWithApiInfo("Ram", "2015-05-01", "DeleteLoginProfile", "", "")
	resp = &DeleteLoginProfileResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteLoginProfileResponse struct {
	responses.BaseResponse
	RequestId common.String
}
