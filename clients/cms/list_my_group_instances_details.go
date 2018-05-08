package cms

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	req.InitWithApiInfo("Cms", "2018-03-08", "ListMyGroupInstancesDetails", "cms", "")
	resp = &ListMyGroupInstancesDetailsResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListMyGroupInstancesDetailsResponse struct {
	responses.BaseResponse
	RequestId    common.String
	Success      bool
	ErrorCode    common.Integer
	ErrorMessage common.String
	PageNumber   common.Integer
	PageSize     common.Integer
	Total        common.Integer
	Resources    ListMyGroupInstancesDetailsResourceList
}

type ListMyGroupInstancesDetailsResource struct {
	AliUid       common.Long
	InstanceName common.String
	InstanceId   common.String
	Desc         common.String
	NetworkType  common.String
	Category     common.String
	Tags         ListMyGroupInstancesDetailsTagList
	Region       ListMyGroupInstancesDetailsRegion
	Vpc          ListMyGroupInstancesDetailsVpc
}

type ListMyGroupInstancesDetailsTag struct {
	Key   common.String
	Value common.String
}

type ListMyGroupInstancesDetailsRegion struct {
	RegionId         common.String
	AvailabilityZone common.String
}

type ListMyGroupInstancesDetailsVpc struct {
	VpcInstanceId     common.String
	VswitchInstanceId common.String
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
