package model

import (
	"kamgo/model/orm"
)

type kamLocation struct {
	orm.Model    `gorm:"-" json:"-"`
	ID           string `json:"id" json:"id"`
	Ruid         string `json:"ruid" json:"ruid"`
	Username     string `json:"username" json:"username"`
	Domain       string `json:"domain" json:"domain"`
	Contact      string `json:"contact" json:"contact"`
	Received     string `json:"received" json:"received"`
	Path         string `json:"path" json:"path"`
	Expires      string `json:"expires" json:"expires"`
	Q            string `json:"q" json:"q"`
	Callid       string `json:"callid" json:"callid"`
	Cseq         string `json:"cseq" json:"cseq"`
	LastModified string `json:"last_modified" json:"last_modified"`
	Flags        string `json:"flags" json:"flags"`
	Cflags       string `json:"cflags" json:"cflags"`
	UserAgent    string `json:"user_agent" json:"user_agent"`
	Socket       string `json:"socket" json:"socket"`
	Methods      string `json:"methods" json:"methods"`
	Instance     string `json:"instance" json:"instance"`
	RegID        string `json:"reg_id" json:"reg_id"`
	ServerID     string `json:"server_id" json:"server_id"`
	ConnectionID string `json:"connection_id" json:"connection_id"`
	Keepalive    string `json:"keepalive" json:"keepalive"`
	Partition    string `json:"partition" json:"partition"`
}

func GetKamLocation(username, domain string) (*kamLocation, error) {
	var tempSub = new(kamLocation)
	if err := KamDB().Where(kamLocation{Domain: domain, Username: username}).First(tempSub).Error; err != nil {
		return nil, err
	}
	return tempSub, nil
}

func (kamLoc kamLocation) TableName() string {
	return "location"
}


