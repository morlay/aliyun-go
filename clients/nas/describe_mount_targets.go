package nas

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeMountTargetsRequest struct {
	requests.RpcRequest
	MountTargetDomain string `position:"Query" name:"MountTargetDomain"`
	PageSize          int    `position:"Query" name:"PageSize"`
	PageNumber        int    `position:"Query" name:"PageNumber"`
	FileSystemId      string `position:"Query" name:"FileSystemId"`
}

func (req *DescribeMountTargetsRequest) Invoke(client *sdk.Client) (resp *DescribeMountTargetsResponse, err error) {
	req.InitWithApiInfo("NAS", "2017-06-26", "DescribeMountTargets", "nas", "")
	resp = &DescribeMountTargetsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeMountTargetsResponse struct {
	responses.BaseResponse
	RequestId    common.String
	TotalCount   common.Integer
	PageSize     common.Integer
	PageNumber   common.Integer
	MountTargets DescribeMountTargetsMountTargetList
}

type DescribeMountTargetsMountTarget struct {
	MountTargetDomain common.String
	NetworkType       common.String
	VpcId             common.String
	VswId             common.String
	AccessGroup       common.String
	Status            common.String
}

type DescribeMountTargetsMountTargetList []DescribeMountTargetsMountTarget

func (list *DescribeMountTargetsMountTargetList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeMountTargetsMountTarget)
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
