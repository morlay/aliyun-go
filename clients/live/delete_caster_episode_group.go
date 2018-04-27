package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteCasterEpisodeGroupRequest struct {
	requests.RpcRequest
	OwnerId   int64  `position:"Query" name:"OwnerId"`
	ProgramId string `position:"Query" name:"ProgramId"`
}

func (req *DeleteCasterEpisodeGroupRequest) Invoke(client *sdk.Client) (resp *DeleteCasterEpisodeGroupResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "DeleteCasterEpisodeGroup", "live", "")
	resp = &DeleteCasterEpisodeGroupResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteCasterEpisodeGroupResponse struct {
	responses.BaseResponse
	RequestId string
}
