package lubanruler

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetResultByLubanrulerRequest struct {
	requests.RpcRequest
	AonePipelineId int    `position:"Query" name:"AonePipelineId"`
	Token          string `position:"Query" name:"Token"`
}

func (req *GetResultByLubanrulerRequest) Invoke(client *sdk.Client) (resp *GetResultByLubanrulerResponse, err error) {
	req.InitWithApiInfo("Lubanruler", "2017-12-28", "GetResultByLubanruler", "", "")
	resp = &GetResultByLubanrulerResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetResultByLubanrulerResponse struct {
	responses.BaseResponse
	Code      string
	Message   string
	RunStatus string
	RequestId string
	Data      GetResultByLubanrulerData
}

type GetResultByLubanrulerData struct {
	AonePipelineId         int
	AppName                string
	ScmUrl                 string
	ScmBranch              string
	TaskStatus             string
	BlockerCnt             int
	CriticalCnt            int
	InfoCnt                int
	MajorCnt               int
	MinorCnt               int
	Complexity             string
	DuplicatedLinesDensity string
	ReliabilityRating      string
	SecurityRating         string
	SqaleRating            string
	LineOfCode             string
	Functions              string
	ResultUrl              string
	CommentLinesDensity    string
}
