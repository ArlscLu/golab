package tgorm

import (
	"encoding/json"
	"testing"
)

func TestSelectAll(t *testing.T) {
	var records []Dict
	db.Table("t_dict").Where("type = ?", "vipgift_type").Find(&records)
	bytes, err := json.MarshalIndent(records, "", "    ")
	if err != nil {
		t.Log(err.Error())
		return
	}
	t.Log(string(bytes))
}
