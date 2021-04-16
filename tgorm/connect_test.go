package tgorm

import "testing"

func TestConnect(t *testing.T) {
	var records []Dict
	db.Table("t_dict").Where("type = ?", "vipgift_type").Find(&records)
	t.Log(records)
}
