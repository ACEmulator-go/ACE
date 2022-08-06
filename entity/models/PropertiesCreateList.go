package models

import (
	"github.com/ACEmulator-go/ACE/entity/enums"
)

type PropertiesCreateList struct {
	DatabaseRecordID uint
	DestinationType  enums.DestinationType
	WeenieClassID    uint
	StackSize        int
	Palette          int8
	Shade            float32
	TryToBond        bool
}

func (prop PropertiesCreateList) Clone() PropertiesCreateList {
	return PropertiesCreateList{
		DatabaseRecordID: prop.DatabaseRecordID,
		DestinationType:  prop.DestinationType,
		WeenieClassID:    prop.WeenieClassID,
		StackSize:        prop.StackSize,
		Palette:          prop.Palette,
		Shade:            prop.Shade,
		TryToBond:        prop.TryToBond,
	}
}
