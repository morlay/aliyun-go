package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeDomainBpsDataByTimeStampRequest struct {
	requests.RpcRequest
	IspNames      string `position:"Query" name:"IspNames"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	LocationNames string `position:"Query" name:"LocationNames"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	TimePoint     string `position:"Query" name:"TimePoint"`
}

func (req *DescribeDomainBpsDataByTimeStampRequest) Invoke(client *sdk.Client) (resp *DescribeDomainBpsDataByTimeStampResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainBpsDataByTimeStamp", "", "")
	resp = &DescribeDomainBpsDataByTimeStampResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDomainBpsDataByTimeStampResponse struct {
	responses.BaseResponse
	RequestId   common.String
	DomainName  common.String
	TimeStamp   common.String
	BpsDataList DescribeDomainBpsDataByTimeStampBpsDataModelList
}

type DescribeDomainBpsDataByTimeStampBpsDataModel struct {
	LocationName common.String
	IspName      common.String
	Bps          common.Long
}

type DescribeDomainBpsDataByTimeStampBpsDataModelList []DescribeDomainBpsDataByTimeStampBpsDataModel

func (list *DescribeDomainBpsDataByTimeStampBpsDataModelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainBpsDataByTimeStampBpsDataModel)
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
