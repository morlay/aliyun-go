package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId  common.String
	PageNumber common.Integer
	PageSize   common.Integer
	Total      common.Integer
	HostList   ListClusterHostHostList
}

type ListClusterHostHost struct {
	HostId         common.String
	HostName       common.String
	PublicIp       common.String
	PrivateIp      common.String
	Role           common.String
	InstanceType   common.String
	Cpu            common.Integer
	Memory         common.Integer
	Status         common.String
	Type           common.String
	HostInstanceId common.String
	SerialNumber   common.String
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
