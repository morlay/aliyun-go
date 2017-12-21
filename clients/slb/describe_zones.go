package slb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeZonesRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
}

func (r DescribeZonesRequest) Invoke(client *sdk.Client) (response *DescribeZonesResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeZonesRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Slb", "2014-05-15", "DescribeZones", "slb", "")

	resp := struct {
		*responses.BaseResponse
		DescribeZonesResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeZonesResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeZonesResponse struct {
	RequestId string
	Zones     DescribeZonesZoneList
}

type DescribeZonesZone struct {
	ZoneId     string
	LocalName  string
	SlaveZones DescribeZonesSlaveZoneList
}

type DescribeZonesSlaveZone struct {
	ZoneId    string
	LocalName string
}

type DescribeZonesZoneList []DescribeZonesZone

func (list *DescribeZonesZoneList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeZonesZone)
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

type DescribeZonesSlaveZoneList []DescribeZonesSlaveZone

func (list *DescribeZonesSlaveZoneList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeZonesSlaveZone)
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
