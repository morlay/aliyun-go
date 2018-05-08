package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteCasterEpisodeRequest struct {
	requests.RpcRequest
	CasterId  string `position:"Query" name:"CasterId"`
	OwnerId   int64  `position:"Query" name:"OwnerId"`
	EpisodeId string `position:"Query" name:"EpisodeId"`
}

func (req *DeleteCasterEpisodeRequest) Invoke(client *sdk.Client) (resp *DeleteCasterEpisodeResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "DeleteCasterEpisode", "live", "")
	resp = &DeleteCasterEpisodeResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteCasterEpisodeResponse struct {
	responses.BaseResponse
	RequestId common.String
	CasterId  common.String
	EpisodeId common.String
}
