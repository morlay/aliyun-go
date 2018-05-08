package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetHdfsCapacityStatisticInfoRequest struct {
	requests.RpcRequest
	FromDatetime    string `position:"Query" name:"FromDatetime"`
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	ToDatetime      string `position:"Query" name:"ToDatetime"`
}

func (req *GetHdfsCapacityStatisticInfoRequest) Invoke(client *sdk.Client) (resp *GetHdfsCapacityStatisticInfoResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "GetHdfsCapacityStatisticInfo", "", "")
	resp = &GetHdfsCapacityStatisticInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetHdfsCapacityStatisticInfoResponse struct {
	responses.BaseResponse
	RequestId        common.String
	HdfsCapacityList GetHdfsCapacityStatisticInfoClusterStatHdfsCapacityList
}

type GetHdfsCapacityStatisticInfoClusterStatHdfsCapacity struct {
	CapacityTotal       common.Long
	CapacityTotalGB     common.Long
	CapacityUsed        common.Long
	CapacityUsedGB      common.Long
	CapacityRemaining   common.Long
	CapacityRemainingGB common.Long
	CapacityUsedNonDfs  common.Long
	ClusterBizId        common.String
	DateTime            common.String
}

type GetHdfsCapacityStatisticInfoClusterStatHdfsCapacityList []GetHdfsCapacityStatisticInfoClusterStatHdfsCapacity

func (list *GetHdfsCapacityStatisticInfoClusterStatHdfsCapacityList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetHdfsCapacityStatisticInfoClusterStatHdfsCapacity)
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
