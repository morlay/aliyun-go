package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId     common.String
	PageNumber    common.Integer
	PageSize      common.Integer
	Total         common.Integer
	ComponentList ListClusterHostComponentComponentList
}

type ListClusterHostComponentComponent struct {
	ServiceName          common.String
	ServiceDisplayName   common.String
	ComponentName        common.String
	ComponentDisplayName common.String
	Status               common.String
	NeedRestart          bool
	HostId               common.String
	ServerStatus         common.String
	HostName             common.String
	PublicIp             common.String
	PrivateIp            common.String
	Role                 common.String
	InstanceType         common.String
	Cpu                  common.Integer
	Memory               common.Integer
	HostInstanceId       common.String
	SerialNumber         common.String
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
