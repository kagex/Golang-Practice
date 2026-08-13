package main

import (
	"fmt"
	"slices"
	"sort"
	"strings"
)

type userData map[string][]string

func main() {
	friendsData := userData{
		"Алексей":  {"Иван", "Сергей", "Елена"},
		"Иван":     {"Алексей", "Дмитрий", "Мария"},
		"Сергей":   {"Алексей", "Елена"},
		"Дмитрий":  {"Иван", "Елена", "Ольга"},
		"Елена":    {"Алексей", "Сергей", "Дмитрий"},
		"Мария":    {"Иван", "Ольга"},
		"Ольга":    {"Дмитрий", "Мария"},
		"Анна":     {"Петр"},
		"Петр":     {"Анна", "Сергей"},
		"Светлана": {"Иван", "Елена"},
	}

	friendsCount := countFriends(friendsData)

	names := make([]string, 0, len(friendsCount))

	for name := range friendsCount {
		names = append(names, name)
	}

	sort.Strings(names)
	fmt.Println("Количество друзей:")
	for _, name := range names {
		fmt.Printf("%s: %d \n", name, friendsCount[name])
	}

	fmt.Println("Введите имена двух пользователей через пробел, а я выведу общих друзей этих пользователей")

	var user1, user2 string
	fmt.Scanln(&user1, &user2)

	commonFriends := commonFriends(friendsData, user1, user2)
	slices.Sort(commonFriends)
	fmt.Printf("Общие друзья между пользователями %s и %s: %s.\n", user1, user2, strings.Join(commonFriends, ", "))

	popularUsers, maxFriends := mostPopularUsers(friendsData)
	fmt.Printf("Наиболее популярные пользователи: %s (количество друзей: %d).", strings.Join(popularUsers, ", "), maxFriends)
}

func countFriends(friendsData userData) map[string]int {
	friendsCount := make(map[string]int, len(friendsData))
	for userName, friendList := range friendsData {
		friendsCount[userName] = len(friendList)
	}
	return friendsCount
}

func commonFriends(friendsData userData, user1, user2 string) []string {
	friendsUser1 := make(map[string]struct{})
	for _, friend := range friendsData[user1] {
		friendsUser1[friend] = struct{}{}
	}
	commonFriends := []string{}
	for _, friend := range friendsData[user2] {
		if _, ok := friendsUser1[friend]; ok {
			commonFriends = append(commonFriends, friend)
		}
	}
	return commonFriends
}

func mostPopularUsers(friendsData userData) ([]string, int) {
	maxFriends := 0
	users := []string{}

	for user, friends := range friendsData {
		friendsCount := len(friends)
		if friendsCount > maxFriends {
			users = []string{user}
			maxFriends = friendsCount
		} else if maxFriends == friendsCount {
			users = append(users, user)
		}
	}
	return users, maxFriends
}
