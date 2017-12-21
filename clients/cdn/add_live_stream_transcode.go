package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddLiveStreamTranscodeRequest struct {
	Template      string `position:"Query" name:"Template"`
	App           string `position:"Query" name:"App"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	OwnerAccount  string `position:"Query" name:"OwnerAccount"`
	Domain        string `position:"Query" name:"Domain"`
	Record        string `position:"Query" name:"Record"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Snapshot      string `position:"Query" name:"Snapshot"`
}

func (r AddLiveStreamTranscodeRequest) Invoke(client *sdk.Client) (response *AddLiveStreamTranscodeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AddLiveStreamTranscodeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "AddLiveStreamTranscode", "", "")

	resp := struct {
		*responses.BaseResponse
		AddLiveStreamTranscodeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AddLiveStreamTranscodeResponse

	err = client.DoAction(&req, &resp)
	return
}

type AddLiveStreamTranscodeResponse struct {
	RequestId string
}
