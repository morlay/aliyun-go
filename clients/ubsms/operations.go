package ubsms

import (
	"encoding/json"

	"github.com/morlay/aliyun-go/core"
)

func (c *UbsmsClient) NotifyUserBusinessCommand(req *NotifyUserBusinessCommandArgs) (resp *NotifyUserBusinessCommandResponse, err error) {
	resp = &NotifyUserBusinessCommandResponse{}
	err = c.Request("NotifyUserBusinessCommand", req, resp)
	return
}

type NotifyUserBusinessCommandArgs struct {
	Uid         string
	ServiceCode string
	Cmd         string
	Region      string
	InstanceId  string
	ClientToken string
	Password    string
}
type NotifyUserBusinessCommandResponse struct {
	RequestId string
	Success   core.Bool
}

func (c *UbsmsClient) SetUserBusinessStatuses(req *SetUserBusinessStatusesArgs) (resp *SetUserBusinessStatusesResponse, err error) {
	resp = &SetUserBusinessStatusesResponse{}
	err = c.Request("SetUserBusinessStatuses", req, resp)
	return
}

type SetUserBusinessStatusesArgs struct {
	Uid           string
	ServiceCode   string
	StatusKey1    string
	StatusValue1  string
	StatusKey2    string
	StatusValue2  string
	StatusKey3    string
	StatusValue3  string
	StatusKey4    string
	StatusValue4  string
	StatusKey5    string
	StatusValue5  string
	StatusKey6    string
	StatusValue6  string
	StatusKey7    string
	StatusValue7  string
	StatusKey8    string
	StatusValue8  string
	StatusKey9    string
	StatusValue9  string
	StatusKey10   string
	StatusValue10 string
	Password      string
}
type SetUserBusinessStatusesResponse struct {
	RequestId string
	Success   core.Bool
}

func (c *UbsmsClient) DescribeBusinessStatus(req *DescribeBusinessStatusArgs) (resp *DescribeBusinessStatusResponse, err error) {
	resp = &DescribeBusinessStatusResponse{}
	err = c.Request("DescribeBusinessStatus", req, resp)
	return
}

type DescribeBusinessStatusUserBusinessStatus struct {
	Uid         string
	ServiceCode string
	Statuses    DescribeBusinessStatusStatusList
}

type DescribeBusinessStatusStatus struct {
	StatusKey   string
	StatusValue string
}
type DescribeBusinessStatusArgs struct {
	CallerBid string
	Password  string
}
type DescribeBusinessStatusResponse struct {
	RequestId              string
	Success                core.Bool
	UserBusinessStatusList DescribeBusinessStatusUserBusinessStatusList
}

type DescribeBusinessStatusStatusList []DescribeBusinessStatusStatus

func (list *DescribeBusinessStatusStatusList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeBusinessStatusStatus)
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

type DescribeBusinessStatusUserBusinessStatusList []DescribeBusinessStatusUserBusinessStatus

func (list *DescribeBusinessStatusUserBusinessStatusList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeBusinessStatusUserBusinessStatus)
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
