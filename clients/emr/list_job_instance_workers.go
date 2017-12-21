package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListJobInstanceWorkersRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	JobInstanceId   string `position:"Query" name:"JobInstanceId"`
}

func (r ListJobInstanceWorkersRequest) Invoke(client *sdk.Client) (response *ListJobInstanceWorkersResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ListJobInstanceWorkersRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Emr", "2016-04-08", "ListJobInstanceWorkers", "", "")

	resp := struct {
		*responses.BaseResponse
		ListJobInstanceWorkersResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ListJobInstanceWorkersResponse

	err = client.DoAction(&req, &resp)
	return
}

type ListJobInstanceWorkersResponse struct {
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
