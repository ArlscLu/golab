package df

import (
	"bytes"
	"encoding/json"
	"os"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func newConn() (*gorm.DB, error) {
	dsn := "test_minqijia:test_minqijia_rw@tcp(10.20.1.20:3306)/towngas_community?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
func init() {
	d, err := newConn()
	if err != nil {
		panic(err)
	}
	db = d
}
func TestGormConn(t *testing.T) {
	var rc RewardConfModel
	db.First(&rc)
	t.Logf("%+v", rc)
}

func TestGormSelectSome(t *testing.T) {
	var rcs []RewardConfModel
	db.Where(&RewardConfModel{BonusPoints: 1}).Find(&rcs)
	for i, rc := range rcs {
		t.Logf("找到的第%d个,%+v\n", i+1, rc)
	}
	pt(rcs)
}
func pt(i interface{}) {
	z, _ := json.Marshal(i)
	var b bytes.Buffer
	_ = json.Indent(&b, z, "", "    ")
	os.Stdout.WriteString(b.String())
}

// 奖励配置
type RewardConfModel struct {
	// GormCommonField
	Type        int8   `gorm:"Column:type" json:"type"`
	Name        string `gorm:"Column:name" json:"name"`
	BonusPoints int64  `gorm:"Column:bonus_points" json:"bonus_points"`
	DailyCap    int64  `gorm:"Column:daily_cap" json:"daily_cap"`
	//modify_time

}

func (RewardConfModel) TableName() string {
	return "t_reward_conf"
}
