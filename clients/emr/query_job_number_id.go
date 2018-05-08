package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type QueryJobNumberIdRequest struct {
	requests.RpcRequest
	JobBizId        string `position:"Query" name:"JobBizId"`
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
}

func (req *QueryJobNumberIdRequest) Invoke(client *sdk.Client) (resp *QueryJobNumberIdResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "QueryJobNumberId", "", "")
	resp = &QueryJobNumberIdResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryJobNumberIdResponse struct {
	responses.BaseResponse
	RequestId   common.String
	JobNumberId common.String
}
