package live

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type StopMixStreamsServiceRequest struct {
	requests.RpcRequest
	SecurityToken  string `position:"Query" name:"SecurityToken"`
	MainDomainName string `position:"Query" name:"MainDomainName"`
	MixStreamName  string `position:"Query" name:"MixStreamName"`
	MixDomainName  string `position:"Query" name:"MixDomainName"`
	OwnerId        int64  `position:"Query" name:"OwnerId"`
	MainAppName    string `position:"Query" name:"MainAppName"`
	MixAppName     string `position:"Query" name:"MixAppName"`
	MainStreamName string `position:"Query" name:"MainStreamName"`
}

func (req *StopMixStreamsServiceRequest) Invoke(client *sdk.Client) (resp *StopMixStreamsServiceResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "StopMixStreamsService", "live", "")
	resp = &StopMixStreamsServiceResponse{}
	err = client.DoAction(req, resp)
	return
}

type StopMixStreamsServiceResponse struct {
	responses.BaseResponse
	RequestId          string
	MixStreamsInfoList StopMixStreamsServiceMixStreamsInfoList
}

type StopMixStreamsServiceMixStreamsInfo struct {
	DomainName string
	AppName    string
	StreamName string
}

type StopMixStreamsServiceMixStreamsInfoList []StopMixStreamsServiceMixStreamsInfo

func (list *StopMixStreamsServiceMixStreamsInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]StopMixStreamsServiceMixStreamsInfo)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
