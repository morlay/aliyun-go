package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListClusterHostComponentRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	HostName        string `position:"Query" name:"HostName"`
	HostInstanceId  string `position:"Query" name:"HostInstanceId"`
	PageSize        int    `position:"Query" name:"PageSize"`
	ComponentName   string `position:"Query" name:"ComponentName"`
	ServiceName     string `position:"Query" name:"ServiceName"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	PageNumber      int    `position:"Query" name:"PageNumber"`
}

func (req *ListClusterHostComponentRequest) Invoke(client *sdk.Client) (resp *ListClusterHostComponentResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ListClusterHostComponent", "", "")
	resp = &ListClusterHostComponentResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListClusterHostComponentResponse struct {
	responses.BaseResponse
	RequestId     string
	PageNumber    int
	PageSize      int
	Total         int
	ComponentList ListClusterHostComponentComponentList
}

type ListClusterHostComponentComponent struct {
	ServiceName          string
	ServiceDisplayName   string
	ComponentName        string
	ComponentDisplayName string
	Status               string
	NeedRestart          bool
	HostId               string
	ServerStatus         string
	HostName             string
	PublicIp             string
	PrivateIp            string
	Role                 string
	InstanceType         string
	Cpu                  int
	Memory               int
	HostInstanceId       string
	SerialNumber         string
}

type ListClusterHostComponentComponentList []ListClusterHostComponentComponent

func (list *ListClusterHostComponentComponentList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterHostComponentComponent)
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
