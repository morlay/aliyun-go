package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId common.String
	HostList  ListFlowClusterHostHostList
}

type ListFlowClusterHostHost struct {
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
