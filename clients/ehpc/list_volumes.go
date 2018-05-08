package ehpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListVolumesRequest struct {
	requests.RpcRequest
	PageSize   int `position:"Query" name:"PageSize"`
	PageNumber int `position:"Query" name:"PageNumber"`
}

func (req *ListVolumesRequest) Invoke(client *sdk.Client) (resp *ListVolumesResponse, err error) {
	req.InitWithApiInfo("EHPC", "2017-07-14", "ListVolumes", "ehs", "")
	resp = &ListVolumesResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListVolumesResponse struct {
	responses.BaseResponse
	RequestId  common.String
	TotalCount common.Integer
	PageNumber common.Integer
	PageSize   common.Integer
	Volumes    ListVolumesVolumeInfoList
}

type ListVolumesVolumeInfo struct {
	RegionId         common.String
	ClusterId        common.String
	ClusterName      common.String
	VolumeId         common.String
	VolumeType       common.String
	VolumeProtocol   common.String
	VolumeMountpoint common.String
	RemoteDirectory  common.String
}

type ListVolumesVolumeInfoList []ListVolumesVolumeInfo

func (list *ListVolumesVolumeInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListVolumesVolumeInfo)
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
