package labdb

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

var sl1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func TestSliceFilter(t *testing.T) {
	n := 0
	for _, x := range sl1 {
		if x%2 == 0 {
			sl1[n] = x
			n++
		}
	}
	sl2 := sl1[:n]
	sl1 := sl1[:n]
	t.Log(sl2, sl1)
}

//list归并  struct的方式
// func groupByReward(tcs []community.TaskConfig) error {
// 	n := 0
// 	tasks := []taskProgressTask{}
// 	for _, v := range tcs {
// 		if v.TaskID != n {
// 			tasks = append(tasks, taskProgressTask{
// 				TaskID:    v.TaskID,
// 				TaskKind:  v.TaskKind,
// 				TaskTitle: v.TaskTitle,
// 				List:      []taskProgressReward{},
// 			})
// 			tasks[n].List = append(tasks[n].List, taskProgressReward{
// 				Reward: Reward{
// 					RewardType:  v.RewardType,
// 					RewardValue: v.RewardValue,
// 				},
// 			})
// 			n++
// 			continue
// 		}
// 		tasks[n-1].List = append(tasks[n-1].List, taskProgressReward{
// 			Reward: Reward{
// 				RewardType:  v.RewardType,
// 				RewardValue: v.RewardValue,
// 			},
// 		})
// 	}
// 	return nil
// }

//list归并 map[string]interface{}方式
// func groupByMap() {
// tasks := []map[string]interface{}{}
// 	var already []int
// 	for _, v := range tcs {
// 		reward := map[string]interface{}{
// 			"id":           v.ID,
// 			"reward_rule":  v.RewardRule,
// 			"reward_type":  v.RewardType,
// 			"reward_value": v.RewardValue,
// 		}
// 		if !utils.IntInSlice(v.TaskID, already) {
// 			task := map[string]interface{}{
// 				"task_id":    v.TaskID,
// 				"task_kind":  v.TaskKind,
// 				"task_title": v.TaskTitle,
// 			}
// 			var list []map[string]interface{}
// 			list = append(list, reward)
// 			task["list"] = list
// 			tasks = append(tasks, task)
// 			already = append(already, v.TaskID)
// 			continue
// 		}
// 		last := len(tasks) - 1
// 		if vv, ok := tasks[last]["list"].([]map[string]interface{}); ok {
// 			tasks[last]["list"] = append(vv, reward)
// 		}
// 	}
// 	return tasks
// }
var base = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var sl0 = base[2:6]

func TestBaseSlice(t *testing.T) {
	base := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	sl1 := base[2:7]
	sl2 := sl1[0:0:0]
	sl3 := []int(nil)
	sl1 = append(sl1, 1)
	sl1 = append(sl1, 1)
	sl1 = append(sl1, 1)
	sl1 = append(sl1, 1)
	sl1 = append(sl1, 1)
	sl1 = append(sl1[:3], sl1[8:]...)
	t.Log(sl1, sl2, sl3)
}

func TestBackSpace(t *testing.T) {
	for n := byte(60); n < 91; n++ {
		os.Stdout.Write([]byte{n, '\b', 'p', '\n'})
	}
}

func TestCopy(t *testing.T) {
	dst := make([]int, 2) //need 2 only
	n := copy(dst, sl0)
	if n != len(dst) {
		t.Errorf("copy not succeed, copyed %d", n)
	}
	t.Log(sl0, dst)
}

func TestParameterFrom(t *testing.T) {
	sl1 := base[3:7]
	t.Log(sl1)
	parameterTo(sl1)
	t.Log(sl1)
}
func parameterTo(s []int) {
	for i := 0; i < len(s); i++ {
		logrus.Printf("%p", &s[i])
	}
	for _, v := range s {
		logrus.Printf("%p", &v)
	}
}

func TestNextToken(t *testing.T) {
	t.Logf("%d", []int{1, 2, 3})
}
