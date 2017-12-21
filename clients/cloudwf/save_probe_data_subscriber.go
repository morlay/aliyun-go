package cloudwf

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SaveProbeDataSubscriberRequest struct {
	requests.RpcRequest
	ApiUrl         string                                  `position:"Query" name:"ApiUrl"`
	ParamGenScript string                                  `position:"Query" name:"ParamGenScript"`
	Name           string                                  `position:"Query" name:"Name"`
	HttpMethod     string                                  `position:"Query" name:"HttpMethod"`
	Description    string                                  `position:"Query" name:"Description"`
	Id             int64                                   `position:"Query" name:"Id"`
	Type           int                                     `position:"Query" name:"Type"`
	ResourceIdss   *SaveProbeDataSubscriberResourceIdsList `position:"Query" type:"Repeated" name:"ResourceIds"`
}

func (req *SaveProbeDataSubscriberRequest) Invoke(client *sdk.Client) (resp *SaveProbeDataSubscriberResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "SaveProbeDataSubscriber", "", "")
	resp = &SaveProbeDataSubscriberResponse{}
	err = client.DoAction(req, resp)
	return
}

type SaveProbeDataSubscriberResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

type SaveProbeDataSubscriberResourceIdsList []int64

func (list *SaveProbeDataSubscriberResourceIdsList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]int64)
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
