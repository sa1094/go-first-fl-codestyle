package main

import (
	"fmt"
	"math/rand"
	"strings"
)

const Warrior string = "warrior"
const Mage string = "mage"
const Healer string = "healer"

func attack(charName, charClass string) string {

	switch charClass {
	case Warrior:
		return fmt.Sprintf("%s нанес урон противнику равный %d.", charName, 5+randint(3, 5))

	case Mage:
		return fmt.Sprintf("%s нанес урон противнику равный %d.", charName, 5+randint(5, 10))

	case Healer:
		return fmt.Sprintf("%s нанес урон противнику равный %d.", charName, 5+randint(-3, -1))
	}
	return "неизвестный класс персонажа"
}

// обратите внимание на "if else" и на "else".
func defence(charName, charClass string) string {

	switch charClass {
	case Warrior:
		return fmt.Sprintf("%s блокировал %d урона.", charName, 10+randint(5, 10))
	case Mage:
		return fmt.Sprintf("%s блокировал %d урона.", charName, 10+randint(-2, 2))
	case Healer:
		return fmt.Sprintf("%s блокировал %d урона.", charName, 10+randint(2, 5))
	}
	return "неизвестный класс персонажа"
}

// обратите внимание на "if else" и на "else".
func special(charName, charClass string) string {

	switch charClass {
	case Warrior:
		return fmt.Sprintf("%s применил специальное умение `Выносливость %d`", charName, 80+25)
	case Mage:
		return fmt.Sprintf("%s применил специальное умение `Атака %d`", charName, 5+40)
	case Healer:
		return fmt.Sprintf("%s применил специальное умение `Защита %d`", charName, 10+30)
	}
	return "неизвестный класс персонажа"
}

// здесь обратите внимание на имена параметров.
func startTraining(charName, charClass string) string {

	switch charClass {
	case Warrior:
		fmt.Printf("%s, ты Воитель - отличный боец ближнего боя.\n", charName)
	case Mage:
		fmt.Printf("%s, ты Маг - превосходный укротитель стихий.\n", charName)
	case Healer:
		fmt.Printf("%s, ты Лекарь - чародей, способный исцелять раны.\n", charName)
	}

	fmt.Println("Потренируйся управлять своими навыками.")
	fmt.Println("Введи одну из команд: attack — чтобы атаковать противника,")
	fmt.Println("defence — чтобы блокировать атаку противника,")
	fmt.Println("special — чтобы использовать свою суперсилу.")
	fmt.Println("Если не хочешь тренироваться, введи команду skip.")

	var cmd string
	for cmd != "skip" {
		fmt.Print("Введи команду: ")
		fmt.Scanf("%s\n", &cmd)

		if cmd == "attack" {
			fmt.Println(attack(charName, charClass))
		}

		if cmd == "defence" {
			fmt.Println(defence(charName, charClass))
		}

		if cmd == "special" {
			fmt.Println(special(charName, charClass))
		}
	}
	return "тренировка окончена"
}

// обратите внимание на имя функции и имена переменных.
func charClass() string {
	var approveChoice string
	var charClass string

	for approveChoice != "y" {
		fmt.Print("Введи название персонажа, за которого хочешь играть: Воитель — warrior, Маг — mage, Лекарь — healer: ")
		fmt.Scanf("%s\n", &charClass)
		switch charClass {
		case Warrior:
			fmt.Println("Воитель — дерзкий воин ближнего боя. Сильный, выносливый и отважный.")
		case Mage:
			fmt.Println("Маг — находчивый воин дальнего боя. Обладает высоким интеллектом.")
		case Healer:
			fmt.Println("Лекарь — могущественный заклинатель. Черпает силы из природы, веры и духов.")
		}
		fmt.Print("Нажми (Y), чтобы подтвердить выбор, или любую другую кнопку, чтобы выбрать другого персонажа: ")
		fmt.Scanf("%s\n", &approveChoice)
		approveChoice = strings.ToLower(approveChoice)
	}
	return charClass
}

// обратите внимание на имена переменных.
func main() {
	fmt.Println("Приветствую тебя, искатель приключений!")
	fmt.Println("Прежде чем начать игру...")

	var charName string
	fmt.Print("...назови себя: ")
	fmt.Scanf("%s\n", &charName)

	fmt.Printf("Здравствуй, %s\n", charName)
	fmt.Println("Сейчас твоя выносливость — 80, атака — 5 и защита — 10.")
	fmt.Println("Ты можешь выбрать один из трёх путей силы:")
	fmt.Println("Воитель, Маг, Лекарь")

	charClass := charClass()

	fmt.Println(startTraining(charName, charClass))
}

func randint(min, max int) int {
	return rand.Intn(max-min) + min
}
