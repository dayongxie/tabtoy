// Generated by github.com/davyxu/tabtoy
// Version: 2.8.8
// DO NOT EDIT!!
package table

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
)

// Defined in table: Globals
type ActorType int32

const (

	// 唐僧
	ActorType_Leader ActorType = 0

	// 孙悟空
	ActorType_Monkey ActorType = 1

	// 猪八戒
	ActorType_Pig ActorType = 2

	// 沙僧
	ActorType_Hammer ActorType = 3
)

var (
	ActorTypeMapperValueByName = map[string]int32{
		"Leader": 0,
		"Monkey": 1,
		"Pig":    2,
		"Hammer": 3,
	}

	ActorTypeMapperNameByValue = map[int32]string{
		0: "Leader",
		1: "Monkey",
		2: "Pig",
		3: "Hammer",
	}
)

func (self ActorType) String() string {
	name, _ := ActorTypeMapperNameByValue[int32(self)]
	return name
}

// Defined in table: Config
type Config struct {

	//Sample
	Sample []*SampleDefine
}

// Defined in table: Globals
type Vec2 struct {
	X int32

	Y int32
}

// Defined in table: Sample
type Prop struct {

	// 血量
	HP int32

	// 攻击速率
	AttackRate float32

	// 额外类型
	ExType ActorType
}

// Defined in table: Sample
type AttackParam struct {

	// 攻击值
	Value int32
}

// Defined in table: Sample
type SampleDefine struct {

	//唯一ID
	ID int64

	//名称
	Name string `自定义tag:"支持go的struct tag"`

	//图标ID
	IconID int32

	//攻击率
	NumericalRate float32

	//物品id
	ItemID int32

	//BuffID
	BuffID []int32

	Pos *Vec2

	//类型
	Type ActorType

	//技能ID列表
	SkillID []int32

	//攻击参数
	AttackParam *AttackParam

	//单结构解析
	SingleStruct *Prop

	//字符串结构
	StrStruct []*Prop
}

// Config 访问接口
type ConfigTable struct {

	// 表格原始数据
	Config

	// 索引函数表
	indexFuncByName map[string][]func(*ConfigTable) error

	// 清空函数表
	clearFuncByName map[string][]func(*ConfigTable) error

	// 加载前回调
	preFuncList []func(*ConfigTable) error

	// 加载后回调
	postFuncList []func(*ConfigTable) error

	SampleByID map[int64]*SampleDefine

	SampleByName map[string]*SampleDefine
}

// 从json文件加载
func (self *ConfigTable) Load(filename string) error {

	data, err := ioutil.ReadFile(filename)

	if err != nil {
		return err
	}

	var newTab Config

	// 读取
	err = json.Unmarshal(data, &newTab)
	if err != nil {
		return err
	}

	// 所有加载前的回调
	for _, v := range self.preFuncList {
		if err = v(self); err != nil {
			return err
		}
	}

	// 清除前通知
	for _, list := range self.clearFuncByName {
		for _, v := range list {
			if err = v(self); err != nil {
				return err
			}
		}
	}

	// 复制数据
	self.Config = newTab

	// 生成索引
	for _, list := range self.indexFuncByName {
		for _, v := range list {
			if err = v(self); err != nil {
				return err
			}
		}
	}

	// 所有完成时的回调
	for _, v := range self.postFuncList {
		if err = v(self); err != nil {
			return err
		}
	}

	return nil
}

// 注册外部索引入口, 索引回调, 清空回调
func (self *ConfigTable) RegisterIndexEntry(name string, indexCallback func(*ConfigTable) error, clearCallback func(*ConfigTable) error) {

	indexList, _ := self.indexFuncByName[name]
	clearList, _ := self.clearFuncByName[name]

	if indexCallback != nil {
		indexList = append(indexList, indexCallback)
	}

	if clearCallback != nil {
		clearList = append(clearList, clearCallback)
	}

	self.indexFuncByName[name] = indexList
	self.clearFuncByName[name] = clearList
}

// 注册加载前回调
func (self *ConfigTable) RegisterPreEntry(callback func(*ConfigTable) error) {

	self.preFuncList = append(self.preFuncList, callback)
}

// 注册所有完成时回调
func (self *ConfigTable) RegisterPostEntry(callback func(*ConfigTable) error) {

	self.postFuncList = append(self.postFuncList, callback)
}

// 创建一个Config表读取实例
func NewConfigTable() *ConfigTable {
	return &ConfigTable{

		indexFuncByName: map[string][]func(*ConfigTable) error{

			"Sample": {func(tab *ConfigTable) error {

				// Sample
				for _, def := range tab.Sample {

					if _, ok := tab.SampleByID[def.ID]; ok {
						panic(fmt.Sprintf("duplicate index in SampleByID: %v", def.ID))
					}

					if _, ok := tab.SampleByName[def.Name]; ok {
						panic(fmt.Sprintf("duplicate index in SampleByName: %v", def.Name))
					}

					tab.SampleByID[def.ID] = def
					tab.SampleByName[def.Name] = def

				}

				return nil
			}},
		},

		clearFuncByName: map[string][]func(*ConfigTable) error{

			"Sample": {func(tab *ConfigTable) error {

				// Sample

				tab.SampleByID = make(map[int64]*SampleDefine)
				tab.SampleByName = make(map[string]*SampleDefine)

				return nil
			}},
		},

		SampleByID: make(map[int64]*SampleDefine),

		SampleByName: make(map[string]*SampleDefine),
	}
}
