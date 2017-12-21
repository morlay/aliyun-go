package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type StartMixStreamsServiceRequest struct {
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

func (r StartMixStreamsServiceRequest) Invoke(client *sdk.Client) (response *StartMixStreamsServiceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		StartMixStreamsServiceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "StartMixStreamsService", "", "")

	resp := struct {
		*responses.BaseResponse
		StartMixStreamsServiceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.StartMixStreamsServiceResponse

	err = client.DoAction(&req, &resp)
	return
}

type StartMixStreamsServiceResponse struct {
	RequestId          string
	MixStreamsInfoList StartMixStreamsServiceMixStreamsInfoList
}

type StartMixStreamsServiceMixStreamsInfo struct {
	DomainName string
	AppName    string
	StreamName string
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
