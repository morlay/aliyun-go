package drds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDrdsInstanceRequest struct {
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
}

func (r DescribeDrdsInstanceRequest) Invoke(client *sdk.Client) (response *DescribeDrdsInstanceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeDrdsInstanceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Drds", "2017-10-16", "DescribeDrdsInstance", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeDrdsInstanceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeDrdsInstanceResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeDrdsInstanceResponse struct {
	RequestId string
	Success   bool
	Data      DescribeDrdsInstanceData
}

type DescribeDrdsInstanceData struct {
	DrdsInstanceId string
	Type           string
	RegionId       string
	ZoneId         string
	Description    string
	NetworkType    string
	Status         string
	CreateTime     int64
	Version        int64
	Specification  string
	Vips           DescribeDrdsInstanceVipList
}

type DescribeDrdsInstanceVip struct {
	IP        string
	Port      string
	Type      string
	VpcId     string
	VswitchId string
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
