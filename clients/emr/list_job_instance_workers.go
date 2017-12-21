package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListJobInstanceWorkersRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	JobInstanceId   string `position:"Query" name:"JobInstanceId"`
}

func (req *ListJobInstanceWorkersRequest) Invoke(client *sdk.Client) (resp *ListJobInstanceWorkersResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ListJobInstanceWorkers", "", "")
	resp = &ListJobInstanceWorkersResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListJobInstanceWorkersResponse struct {
	responses.BaseResponse
	RequestId          string
	JobInstanceWorkers ListJobInstanceWorkersJobInstanceWorkerInfoList
}

type ListJobInstanceWorkersJobInstanceWorkerInfo struct {
	ApplicationId string
	InstanceInfo  string
	ContainerInfo string
}

type ListJobInstanceWorkersJobInstanceWorkerInfoList []ListJobInstanceWorkersJobInstanceWorkerInfo

func (list *ListJobInstanceWorkersJobInstanceWorkerInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListJobInstanceWorkersJobInstanceWorkerInfo)
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
