package imm

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteFaceSetRequest struct {
	requests.RpcRequest
	LazyMode   string `position:"Query" name:"LazyMode"`
	Project    string `position:"Query" name:"Project"`
	SetId      string `position:"Query" name:"SetId"`
	CheckEmpty string `position:"Query" name:"CheckEmpty"`
}

func (req *DeleteFaceSetRequest) Invoke(client *sdk.Client) (resp *DeleteFaceSetResponse, err error) {
	req.InitWithApiInfo("imm", "2017-09-06", "DeleteFaceSet", "imm", "")
	resp = &DeleteFaceSetResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteFaceSetResponse struct {
	responses.BaseResponse
	RequestId string
}
