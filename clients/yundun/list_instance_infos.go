package yundun

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId  string
	PageNumber int
	PageSize   int
	TotalCount int
	InfosList  ListInstanceInfosInstanceInfoList
}

type ListInstanceInfosInstanceInfo struct {
	Region       string
	RegionName   string
	RegionEnName string
	InstanceName string
	InstanceId   string
	Ip           string
	InternetIp   string
	IntranetIp   string
	Ddos         int
	HostEvent    int
	SecureCheck  int
	AegisStatus  int
	Waf          int
	IsLock       bool
	LockType     string
	UnLockTimes  int
	TriggerTime  string
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
