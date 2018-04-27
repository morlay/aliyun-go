package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListClusterHostGroupRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	HostGroupId     string `position:"Query" name:"HostGroupId"`
	PageSize        int    `position:"Query" name:"PageSize"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	PageNumber      int    `position:"Query" name:"PageNumber"`
}

func (req *ListClusterHostGroupRequest) Invoke(client *sdk.Client) (resp *ListClusterHostGroupResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ListClusterHostGroup", "", "")
	resp = &ListClusterHostGroupResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListClusterHostGroupResponse struct {
	responses.BaseResponse
	RequestId  string
	PageNumber int
	PageSize   int
	Total      int
	HostList   ListClusterHostGroupHostList
}

type ListClusterHostGroupHost struct {
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

type ListClusterHostGroupHostList []ListClusterHostGroupHost

func (list *ListClusterHostGroupHostList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterHostGroupHost)
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
