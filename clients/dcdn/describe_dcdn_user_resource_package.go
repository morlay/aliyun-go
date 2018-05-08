package dcdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDcdnUserResourcePackageRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeDcdnUserResourcePackageRequest) Invoke(client *sdk.Client) (resp *DescribeDcdnUserResourcePackageResponse, err error) {
	req.InitWithApiInfo("dcdn", "2018-01-15", "DescribeDcdnUserResourcePackage", "dcdn", "")
	resp = &DescribeDcdnUserResourcePackageResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDcdnUserResourcePackageResponse struct {
	responses.BaseResponse
	RequestId            string
	ResourcePackageInfos DescribeDcdnUserResourcePackageResourcePackageInfoList
}

type DescribeDcdnUserResourcePackageResourcePackageInfo struct {
	CurrCapacity  string
	InitCapacity  string
	CommodityCode string
	DisplayName   string
	InstanceId    string
	Status        string
}

type DescribeDcdnUserResourcePackageResourcePackageInfoList []DescribeDcdnUserResourcePackageResourcePackageInfo

func (list *DescribeDcdnUserResourcePackageResourcePackageInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDcdnUserResourcePackageResourcePackageInfo)
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
