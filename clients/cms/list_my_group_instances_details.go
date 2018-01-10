package cms

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListMyGroupInstancesDetailsRequest struct {
	requests.RpcRequest
	Total      string `position:"Query" name:"Total"`
	GroupId    int64  `position:"Query" name:"GroupId"`
	PageSize   int    `position:"Query" name:"PageSize"`
	Category   string `position:"Query" name:"Category"`
	PageNumber int    `position:"Query" name:"PageNumber"`
}

func (req *ListMyGroupInstancesDetailsRequest) Invoke(client *sdk.Client) (resp *ListMyGroupInstancesDetailsResponse, err error) {
	req.InitWithApiInfo("Cms", "2017-03-01", "ListMyGroupInstancesDetails", "cms", "")
	resp = &ListMyGroupInstancesDetailsResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListMyGroupInstancesDetailsResponse struct {
	responses.BaseResponse
	RequestId    string
	Success      bool
	ErrorCode    int
	ErrorMessage string
	PageNumber   int
	PageSize     int
	Total        int
	Resources    ListMyGroupInstancesDetailsResourceList
}

type ListMyGroupInstancesDetailsResource struct {
	AliUid       int64
	InstanceName string
	InstanceId   string
	Desc         string
	NetworkType  string
	Category     string
	Tags         ListMyGroupInstancesDetailsTagList
	Region       ListMyGroupInstancesDetailsRegion
	Vpc          ListMyGroupInstancesDetailsVpc
}

type ListMyGroupInstancesDetailsTag struct {
	Key   string
	Value string
}

type ListMyGroupInstancesDetailsRegion struct {
	RegionId         string
	AvailabilityZone string
}

type ListMyGroupInstancesDetailsVpc struct {
	VpcInstanceId     string
	VswitchInstanceId string
}

type ListMyGroupInstancesDetailsResourceList []ListMyGroupInstancesDetailsResource

func (list *ListMyGroupInstancesDetailsResourceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListMyGroupInstancesDetailsResource)
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

type ListMyGroupInstancesDetailsTagList []ListMyGroupInstancesDetailsTag

func (list *ListMyGroupInstancesDetailsTagList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListMyGroupInstancesDetailsTag)
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
