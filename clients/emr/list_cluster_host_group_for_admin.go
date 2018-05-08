package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListClusterHostGroupForAdminRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	HostGroupId     string `position:"Query" name:"HostGroupId"`
	PageSize        int    `position:"Query" name:"PageSize"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	PageNumber      int    `position:"Query" name:"PageNumber"`
}

func (req *ListClusterHostGroupForAdminRequest) Invoke(client *sdk.Client) (resp *ListClusterHostGroupForAdminResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ListClusterHostGroupForAdmin", "", "")
	resp = &ListClusterHostGroupForAdminResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListClusterHostGroupForAdminResponse struct {
	responses.BaseResponse
	RequestId  common.String
	PageNumber common.Integer
	PageSize   common.Integer
	Total      common.Integer
	HostList   ListClusterHostGroupForAdminHostList
}

type ListClusterHostGroupForAdminHost struct {
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

type ListClusterHostGroupForAdminHostList []ListClusterHostGroupForAdminHost

func (list *ListClusterHostGroupForAdminHostList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterHostGroupForAdminHost)
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
