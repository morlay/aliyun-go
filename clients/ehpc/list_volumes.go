package ehpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId  string
	TotalCount int
	PageNumber int
	PageSize   int
	Volumes    ListVolumesVolumeInfoList
}

type ListVolumesVolumeInfo struct {
	RegionId         string
	ClusterId        string
	ClusterName      string
	VolumeId         string
	VolumeType       string
	VolumeProtocol   string
	VolumeMountpoint string
	RemoteDirectory  string
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
