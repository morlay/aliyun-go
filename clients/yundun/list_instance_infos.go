package yundun

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListInstanceInfosRequest struct {
	JstOwnerId   int64  `position:"Query" name:"JstOwnerId"`
	PageNumber   int    `position:"Query" name:"PageNumber"`
	PageSize     int    `position:"Query" name:"PageSize"`
	Region       string `position:"Query" name:"Region"`
	EventType    string `position:"Query" name:"EventType"`
	InstanceName string `position:"Query" name:"InstanceName"`
	InstanceType string `position:"Query" name:"InstanceType"`
	InstanceIds  string `position:"Query" name:"InstanceIds"`
}

func (r ListInstanceInfosRequest) Invoke(client *sdk.Client) (response *ListInstanceInfosResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ListInstanceInfosRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Yundun", "2015-04-16", "ListInstanceInfos", "", "")

	resp := struct {
		*responses.BaseResponse
		ListInstanceInfosResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ListInstanceInfosResponse

	err = client.DoAction(&req, &resp)
	return
}

type ListInstanceInfosResponse struct {
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
