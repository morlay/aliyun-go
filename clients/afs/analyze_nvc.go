package afs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AnalyzeNvcRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	Data            string `position:"Query" name:"Data"`
	ScoreJsonStr    string `position:"Query" name:"ScoreJsonStr"`
}

func (req *AnalyzeNvcRequest) Invoke(client *sdk.Client) (resp *AnalyzeNvcResponse, err error) {
	req.InitWithApiInfo("afs", "2018-01-12", "AnalyzeNvc", "", "")
	resp = &AnalyzeNvcResponse{}
	err = client.DoAction(req, resp)
	return
}

type AnalyzeNvcResponse struct {
	responses.BaseResponse
	RequestId string
	BizCode   string
}
