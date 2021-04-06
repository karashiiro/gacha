package main

// Drop represents a single possible roll from the gacha.
type Drop struct {
	ID   uint32  `gorm:"column:id;not null"`
	Rate float32 `gorm:"column:rate;not null"`
}
