package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListFlowClusterHostRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	ProjectId       string `position:"Query" name:"ProjectId"`
}

func (req *ListFlowClusterHostRequest) Invoke(client *sdk.Client) (resp *ListFlowClusterHostResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ListFlowClusterHost", "", "")
	resp = &ListFlowClusterHostResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListFlowClusterHostResponse struct {
	responses.BaseResponse
	RequestId string
	HostList  ListFlowClusterHostHostList
}

type ListFlowClusterHostHost struct {
	HostId         string
	HostName       string
	PublicIp       string
	PrivateIp      string
	Role           string
	InstanceType   string
	Cpu            int
	Memory         int
	Status         string
	Type           string
	HostInstanceId string
	SerialNumber   string
}

type ListFlowClusterHostHostList []ListFlowClusterHostHost

func (list *ListFlowClusterHostHostList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListFlowClusterHostHost)
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
