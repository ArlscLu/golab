package g2s

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/jellycheng/gosupport"
	"github.com/jellycheng/gosupport/dbutils"
	"github.com/jellycheng/gosupport/env"
	"github.com/jinzhu/gorm"
	"github.com/urfave/cli"
	"gopkg.in/jellycheng/cjsgorm.v1"
)

//数据库定义转换为model struct
var Tool = cli.Command{
	Name:   "g2s",
	Usage:  "数据库定义转换为model struct",
	Action: run,
	Flags: []cli.Flag{
		stringFlag("config,c", ".env", "指定配置文件"),
		stringFlag("keyword,k", "", "关键词"),
		stringFlag("tbl,t", "", "社区库的表名"),
	},
}

func run(ctx *cli.Context) {
	envPath := ctx.String("config")
	tableName := ctx.String("tbl")
	g := NewGenerator(envPath, tableName, ctx)
	err := g.Run()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}

func stringFlag(name, value, usage string) cli.StringFlag {
	return cli.StringFlag{
		Name:  name,
		Value: value,
		Usage: usage,
	}
}

type generator struct {
	EnvPath   string
	TableName string
	Ctx       *cli.Context

	config       *gosupport.DataManage
	db           *gorm.DB
	keywordUpper string
	packageName  string
}

func NewGenerator(envPath, tableName string, ctx *cli.Context) *generator {
	return &generator{
		EnvPath:   envPath,
		TableName: tableName,
		Ctx:       ctx,
	}
}
func (g *generator) Run() error {
	err := g.initConfig()
	if err != nil {
		return err
	}
	err = g.initKeyWord()
	if err != nil {
		return err
	}
	err = g.initDb()
	if err != nil {
		return err
	}
	err = g.initPackageName()
	if err != nil {
		return err
	}
	return g.Output()
}
func (g *generator) initKeyWord() error {
	userSet := g.Ctx.String("keyword")
	//不指定时去env查找第一个符合条件的
	if len(userSet) == 0 {
		find, err := g.defaultKeyWord()
		if err != nil {
			return err
		}
		if len(find) == 0 {
			return errors.New("关键词不能为空")
		}
		userSet = find
	}

	g.keywordUpper = strings.ToUpper(userSet)
	return nil
}
func (g *generator) initPackageName() error {
	database := g.config.GetString(g.keywordUpper + "_DB_WRITE_DATABASE")
	splited := strings.Split(database, "_")
	if len(splited) < 2 {
		return errors.New("生成package发生错误")
	}
	g.packageName = splited[1]
	return nil
}
func (g *generator) defaultKeyWord() (string, error) {
	var wantKeys []string
	for k := range g.config.Data {
		if strings.Contains(k, "DB_WRITE_HOST") {
			wantKeys = append(wantKeys, k)
		}
	}
	if len(wantKeys) == 1 {
		splitBy := strings.Split(wantKeys[0], "_")
		return splitBy[0], nil
	}
	return "", errors.New("包含多个DB_WRITE_HOST定义")
}
func (g *generator) Output() error {
	res := dbutils.GenerateGormModel(g.db.DB(),
		g.config.GetString(g.keywordUpper+"_DB_WRITE_DATABASE"),
		[]string{g.TableName},
		map[string]string{"packageName": g.packageName,
			"ignoreField":      "id,is_delete,create_time,update_time,delete_time,modify_time",
			"appendStructCode": "GormCommonField",
			"trimTblPrefix":    "t_",
			"structNameSuffix": "Model"})
	if v, ok := res[g.TableName]; ok {
		fmt.Print(v)
		return nil
	}
	return errors.New("未查到数据库定义")
}
func (g *generator) initConfig() error {
	err := env.LoadEnv2DataManage(g.EnvPath)
	if err != nil {
		return err
	}
	g.config = gosupport.NewGlobalEnvSingleton()
	return nil
}
func (g *generator) initDb() error {
	keywordUpper := g.keywordUpper
	dsn := dbutils.NewMysqlDsn(map[string]interface{}{
		"host":     g.config.GetString(keywordUpper + "_DB_WRITE_HOST"),
		"username": g.config.GetString(keywordUpper + "_DB_WRITE_USERNAME"),
		"password": g.config.GetString(keywordUpper + "_DB_WRITE_PASSWORD"),
		"port":     g.config.GetString(keywordUpper + "_DB_WRITE_PORT"),
		"dbname":   g.config.GetString(keywordUpper + "_DB_WRITE_DATABASE"),
		"charset":  "utf8mb4",
	})
	gorm := cjsgorm.NewMysqlGorm(*dsn)
	if gorm == nil {
		return errors.New("数据库连接错误")
	}
	g.db = gorm
	return nil
}
