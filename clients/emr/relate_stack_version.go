package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RelateStackVersionRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	MainVersion     string `position:"Query" name:"MainVersion"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	StackVersion    string `position:"Query" name:"StackVersion"`
	UserId          string `position:"Query" name:"UserId"`
}

func (req *RelateStackVersionRequest) Invoke(client *sdk.Client) (resp *RelateStackVersionResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "RelateStackVersion", "", "")
	resp = &RelateStackVersionResponse{}
	err = client.DoAction(req, resp)
	return
}

type RelateStackVersionResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
}
