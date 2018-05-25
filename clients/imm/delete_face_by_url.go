package imm

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteFaceByUrlRequest struct {
	requests.RpcRequest
	Project string `position:"Query" name:"Project"`
	SetId   string `position:"Query" name:"SetId"`
	SrcUri  string `position:"Query" name:"SrcUri"`
}

func (req *DeleteFaceByUrlRequest) Invoke(client *sdk.Client) (resp *DeleteFaceByUrlResponse, err error) {
	req.InitWithApiInfo("imm", "2017-09-06", "DeleteFaceByUrl", "imm", "")
	resp = &DeleteFaceByUrlResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteFaceByUrlResponse struct {
	responses.BaseResponse
	RequestId common.String
}
