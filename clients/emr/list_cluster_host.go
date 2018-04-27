package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListClusterHostRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	HostName        string `position:"Query" name:"HostName"`
	HostInstanceId  string `position:"Query" name:"HostInstanceId"`
	PrivateIp       string `position:"Query" name:"PrivateIp"`
	PageSize        int    `position:"Query" name:"PageSize"`
	ComponentName   string `position:"Query" name:"ComponentName"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	PageNumber      int    `position:"Query" name:"PageNumber"`
}

func (req *ListClusterHostRequest) Invoke(client *sdk.Client) (resp *ListClusterHostResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ListClusterHost", "", "")
	resp = &ListClusterHostResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListClusterHostResponse struct {
	responses.BaseResponse
	RequestId  string
	PageNumber int
	PageSize   int
	Total      int
	HostList   ListClusterHostHostList
}

type ListClusterHostHost struct {
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

type ListClusterHostHostList []ListClusterHostHost

func (list *ListClusterHostHostList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterHostHost)
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
