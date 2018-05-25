package drds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type QueryInstanceInfoByConnRequest struct {
	requests.RpcRequest
	Port     int    `position:"Query" name:"Port"`
	Host     string `position:"Query" name:"Host"`
	UserName string `position:"Query" name:"UserName"`
}

func (req *QueryInstanceInfoByConnRequest) Invoke(client *sdk.Client) (resp *QueryInstanceInfoByConnResponse, err error) {
	req.InitWithApiInfo("Drds", "2017-10-16", "QueryInstanceInfoByConn", "", "")
	resp = &QueryInstanceInfoByConnResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryInstanceInfoByConnResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
	Data      QueryInstanceInfoByConnData
}

type QueryInstanceInfoByConnData struct {
	DrdsInstanceId     common.String
	Type               common.String
	RegionId           common.String
	ZoneId             common.String
	Description        common.String
	NetworkType        common.String
	Status             common.String
	CreateTime         common.Long
	Version            common.Long
	Specification      common.String
	SpecTypeId         common.String
	SpecTypeName       common.String
	VpcCloudInstanceId common.String
	Vips               QueryInstanceInfoByConnVipList
}

type QueryInstanceInfoByConnVip struct {
	IP        common.String
	Port      common.String
	Type      common.String
	VpcId     common.String
	VswitchId common.String
}

type QueryInstanceInfoByConnVipList []QueryInstanceInfoByConnVip

func (list *QueryInstanceInfoByConnVipList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryInstanceInfoByConnVip)
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
