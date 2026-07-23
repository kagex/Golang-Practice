package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	userAttack := CriticalHitDecorator(SlowEffectDecorator(DamageBoostDecorator(CriticalHitDecorator(Attack))))
	fmt.Printf("Рыцарь наносит удар:\n%v", userAttack())
}

func Attack() string {
	return "Атака выполнена!"
}

func DamageBoostDecorator(attackFunc func() string) func() string {

	return func() string {
		return "Вам улыбнулась удача, нанесение урона увеличено на 10%!\n" + attackFunc()
	}
}

func CriticalHitDecorator(attackFunc func() string) func() string {
	return func() string {
		if rand.IntN(100) <= 25 {
			return "Критический удар! Урон удвоен!\n" + attackFunc()
		}
		return attackFunc()
	}
}

func SlowEffectDecorator(attackFunc func() string) func() string {
	return func() string {
		return attackFunc() + "\nЦель замедлена на 2 хода!"
	}
}
