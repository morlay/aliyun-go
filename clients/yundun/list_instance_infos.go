package yundun

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListInstanceInfosRequest struct {
	requests.RpcRequest
	JstOwnerId   int64  `position:"Query" name:"JstOwnerId"`
	PageNumber   int    `position:"Query" name:"PageNumber"`
	PageSize     int    `position:"Query" name:"PageSize"`
	Region       string `position:"Query" name:"Region"`
	EventType    string `position:"Query" name:"EventType"`
	InstanceName string `position:"Query" name:"InstanceName"`
	InstanceType string `position:"Query" name:"InstanceType"`
	InstanceIds  string `position:"Query" name:"InstanceIds"`
}

func (req *ListInstanceInfosRequest) Invoke(client *sdk.Client) (resp *ListInstanceInfosResponse, err error) {
	req.InitWithApiInfo("Yundun", "2015-04-16", "ListInstanceInfos", "", "")
	resp = &ListInstanceInfosResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListInstanceInfosResponse struct {
	responses.BaseResponse
	RequestId  common.String
	PageNumber common.Integer
	PageSize   common.Integer
	TotalCount common.Integer
	InfosList  ListInstanceInfosInstanceInfoList
}

type ListInstanceInfosInstanceInfo struct {
	Region       common.String
	RegionName   common.String
	RegionEnName common.String
	InstanceName common.String
	InstanceId   common.String
	Ip           common.String
	InternetIp   common.String
	IntranetIp   common.String
	Ddos         common.Integer
	HostEvent    common.Integer
	SecureCheck  common.Integer
	AegisStatus  common.Integer
	Waf          common.Integer
	IsLock       bool
	LockType     common.String
	UnLockTimes  common.Integer
	TriggerTime  common.String
}

type ListInstanceInfosInstanceInfoList []ListInstanceInfosInstanceInfo

func (list *ListInstanceInfosInstanceInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListInstanceInfosInstanceInfo)
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
