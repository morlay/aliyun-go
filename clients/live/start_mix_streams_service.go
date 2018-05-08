package live

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type StartMixStreamsServiceRequest struct {
	requests.RpcRequest
	MixType        string `position:"Query" name:"MixType"`
	SecurityToken  string `position:"Query" name:"SecurityToken"`
	MainDomainName string `position:"Query" name:"MainDomainName"`
	MixStreamName  string `position:"Query" name:"MixStreamName"`
	MixTemplate    string `position:"Query" name:"MixTemplate"`
	MixDomainName  string `position:"Query" name:"MixDomainName"`
	OwnerId        int64  `position:"Query" name:"OwnerId"`
	MainAppName    string `position:"Query" name:"MainAppName"`
	MixAppName     string `position:"Query" name:"MixAppName"`
	MainStreamName string `position:"Query" name:"MainStreamName"`
}

func (req *StartMixStreamsServiceRequest) Invoke(client *sdk.Client) (resp *StartMixStreamsServiceResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "StartMixStreamsService", "live", "")
	resp = &StartMixStreamsServiceResponse{}
	err = client.DoAction(req, resp)
	return
}

type StartMixStreamsServiceResponse struct {
	responses.BaseResponse
	RequestId          common.String
	MixStreamsInfoList StartMixStreamsServiceMixStreamsInfoList
}

type StartMixStreamsServiceMixStreamsInfo struct {
	DomainName common.String
	AppName    common.String
	StreamName common.String
}

type StartMixStreamsServiceMixStreamsInfoList []StartMixStreamsServiceMixStreamsInfo

func (list *StartMixStreamsServiceMixStreamsInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]StartMixStreamsServiceMixStreamsInfo)
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
