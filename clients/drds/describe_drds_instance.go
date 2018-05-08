package drds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeDrdsInstanceRequest struct {
	requests.RpcRequest
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
}

func (req *DescribeDrdsInstanceRequest) Invoke(client *sdk.Client) (resp *DescribeDrdsInstanceResponse, err error) {
	req.InitWithApiInfo("Drds", "2017-10-16", "DescribeDrdsInstance", "", "")
	resp = &DescribeDrdsInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDrdsInstanceResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
	Data      DescribeDrdsInstanceData
}

type DescribeDrdsInstanceData struct {
	DrdsInstanceId common.String
	Type           common.String
	RegionId       common.String
	ZoneId         common.String
	Description    common.String
	NetworkType    common.String
	Status         common.String
	CreateTime     common.Long
	Version        common.Long
	Specification  common.String
	Vips           DescribeDrdsInstanceVipList
}

type DescribeDrdsInstanceVip struct {
	IP        common.String
	Port      common.String
	Type      common.String
	VpcId     common.String
	VswitchId common.String
}

type DescribeDrdsInstanceVipList []DescribeDrdsInstanceVip

func (list *DescribeDrdsInstanceVipList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDrdsInstanceVip)
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
