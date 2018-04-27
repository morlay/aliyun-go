package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteCasterProgramRequest struct {
	requests.RpcRequest
	CasterId string `position:"Query" name:"CasterId"`
	OwnerId  int64  `position:"Query" name:"OwnerId"`
}

func (req *DeleteCasterProgramRequest) Invoke(client *sdk.Client) (resp *DeleteCasterProgramResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "DeleteCasterProgram", "live", "")
	resp = &DeleteCasterProgramResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteCasterProgramResponse struct {
	responses.BaseResponse
	RequestId string
	CasterId  string
}
