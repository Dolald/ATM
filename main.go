package main

import (
	"fmt"
	"sort"
	"strconv"
	// печатание чека, давайте сохраним природу ?
)

type language string

var (
	languageRussian language = "russian"
	languageEnglish language = "english"
)

type Clients struct {
	amountOfWithdrawnMoney int
	balance                int
	CardNumber             string
	Name                   string
	PinCode                int
}

type ATM struct {
	balance         int
	banknoteToCount map[int]int
	Clients         []Clients
	language        language
}

func (a *ATM) ChangeLanguage(enterNumberOfLanguage int) {
	switch enterNumberOfLanguage {
	case 1:
		a.language = languageRussian
	case 2:
		a.language = languageEnglish
	}
}

func (c *Clients) CheckBalance() {
	fmt.Println("Ваш баланс:", c.balance)
}

func (a *ATM) Deposit(deposit int) {
	a.balance += deposit
	a.balanceOfСlient += deposit
	fmt.Println("Идёт распознавание купюр")
	fmt.Println("Вы успешно пополнили:", deposit)
}

func (a *ATM) creatNewClient() {
	var pin int
	var cardNumber, name string
	fmt.Println("Введите номер карты, имя и пинкод")
	fmt.Scan(&cardNumber, &pin, &name)

	Newclient := Clients{
		balance:    0,
		CardNumber: cardNumber,
		Name:       name,
		PinCode:    pin,
	}

	a.Clients = append(a.Clients, Newclient)
}

func (a *ATM) Withdraw() {
	for {
		var withdraw, withdrawn, сhoosingBanknote, сhoosingChange, estimatedWithdrawal, countBanknotes, pinCode int
		var displayBanknotes string
		arr := make([]int, 0)
		estimatedWithdrawalAmount := make(map[int]int) //преполагаемая сумма вывода

		fmt.Println("Выберите сумму для снятия") //выводим существующие банкноты
		for banknote := range a.banknoteToCount {
			displayBanknotes += strconv.Itoa(banknote) + " "
		}

		fmt.Println("1 - ввести другую сумму\n", displayBanknotes)

		fmt.Scan(&сhoosingBanknote)

		fmt.Println("Введите пинкод")
		fmt.Scan(&pinCode)
		if pinCode != a.Clients.PinCode[a.Clients.CardNumber] {
			fmt.Println("Не правильный пинкод")
			return
		}

		if сhoosingBanknote != 1 { //если выбор банкноты один из 5
			a.banknoteToCount[сhoosingBanknote] -= 1
			fmt.Println("Вы успешно сняли:", сhoosingBanknote)
			return
		}

		fmt.Println("Введите сумму для снятия наличных\nВведите сумму до 300000\nДоступны номиналы купюр:", displayBanknotes)
		fmt.Scan(&withdraw)

		withdrawn = withdraw

		switch {
		case withdraw%100 != 0 || withdraw > a.balance:
			fmt.Println("Банкомата не выдаёт такие суммы, выберите другую сумму")
			return
		case withdraw > B.balanceOfClient:
			fmt.Println("У вас недостаточно денег для снятия такой суммы")
			return
		case withdraw+amountOfWithdrawnMoneyAmountOfClient[correctPin] > 300000:
			fmt.Println("Максимальная сумма снятия с карты в день - 300 000")
			return
		}

		if _, ex := a.banknoteToCount[100]; !ex { // if нет 100, идёт проверка вводимого числа на кратность 100, но не 500 или вывод < 500
			if withdraw%100 == 0 && withdraw%500 != 0 || withdraw < 500 {
				fmt.Println("Введите другую сумму")
				return
			}
		}

		displayBanknotes = ""

		if withdraw >= 500 && withdraw <= 120000 {
			fmt.Println("Купюры какого размена желаете ?\n1 - Всё равно") //
			for banknote := range a.banknoteToCount {
				if withdraw >= banknote && withdraw < 6000 { // 6000, т.к. это максимальная сумма по минимальным купюром в данном кейсе
					displayBanknotes += strconv.Itoa(banknote) + " "
				} else if withdraw <= banknote*60 && withdraw >= 6000 {
					displayBanknotes += strconv.Itoa(banknote) + " "
				}
			}
			fmt.Println(displayBanknotes)
			fmt.Scan(&сhoosingChange)

		} else if withdraw < 500 || withdraw > 120000 || сhoosingChange == 1 {
			сhoosingChange = 5000
		}

		for banknote, ammount := range a.banknoteToCount { // сумирование предполагаемого вывода включительно до выбранной банкноты
			arr = append(arr, banknote) //заполняем массив банкнотами
			estimatedWithdrawalAmount[banknote] = ammount
			if banknote <= сhoosingChange {
				estimatedWithdrawal += banknote * ammount
			}
		}

		if withdraw > estimatedWithdrawal {
			fmt.Println("Выберите сумму поменьше")
			return
		}

		sort.Sort(sort.Reverse(sort.IntSlice(arr)))

		for _, banknote := range arr { // процесс снятия денег с банкомата
			if withdraw != 0 && banknote <= сhoosingChange {
				estimatedWithdrawalAmount, withdraw = subtractionATMBanknotes(banknote, withdraw, estimatedWithdrawalAmount)
			}
		}

		for banknote, v := range estimatedWithdrawalAmount { //  проверка количества снимаемых банкнот
			if banknote <= сhoosingChange {
				countBanknotes += a.banknoteToCount[banknote] - v
				if countBanknotes > 60 {
					fmt.Println("банкомат не может выдать такое количество купюр")
					return
				}
			}
		}

		for banknote, v := range estimatedWithdrawalAmount { //окончательная передача информации о выводе денег
			a.banknoteToCount[banknote] = v
		}

		a.balance -= withdrawn
		fmt.Println("Выдача купюр. . .\nВы успешно сняли", withdrawn)
		amountOfWithdrawnMoneyOfClient[1973] += withdrawn //amountOfWithdrawnMoneyOfClient += withdrawn

		fmt.Println(a.banknoteToCount)
	}
}

func (a *ATM) checkBanknotesAmount() bool {
	var countBeforeATMBlock int
	for banknote, v := range a.banknoteToCount {
		if v == 0 { // Проверка количества банкнот каждого номинала
			countBeforeATMBlock++
			delete(a.banknoteToCount, banknote)
		}
	}
	if countBeforeATMBlock >= 2 {
		fmt.Println("Банкомат заблокирован")
		return false
	}
	return true
}

func subtractionATMBanknotes(banknote int, withdraw int, estimatedWithdrawalAmount map[int]int) (map[int]int, int) {
	tempEstimatedWithdrawalAmount := estimatedWithdrawalAmount[banknote]
	estimatedWithdrawalAmount[banknote] -= withdraw / banknote

	if estimatedWithdrawalAmount[banknote] < 0 {
		estimatedWithdrawalAmount[banknote] = 0
	}
	if withdraw/banknote != 0 {
		withdraw -= (tempEstimatedWithdrawalAmount - estimatedWithdrawalAmount[banknote]) * banknote
	}
	return estimatedWithdrawalAmount, withdraw
}

func (a *ATM) actionsOfATM() {
	fmt.Println("1 - снять наличные \n2 - пополнить карту \n3 - посмотреть баланс \n4 - поменять язык \n5 - вытащить карту")
	var enterNum int
	fmt.Scan(&enterNum)
	switch enterNum {

	case 1:
		a.Withdraw()
		getBack()
		a.actionsOfATM()

	case 2:
		var deposit int
		fmt.Println("Введите сумму для пополнения")
		fmt.Scan(&deposit)
		a.Deposit(deposit)
		getBack()
		a.actionsOfATM()

	case 3:
		a.CheckBalance()
		getBack()
		a.actionsOfATM()

	case 4:
		var enterNumberOfLanguage int
		fmt.Println("1 - поменять на русский\n2 - поменять на английский")
		fmt.Scan(&enterNumberOfLanguage)
		a.ChangeLanguage(enterNumberOfLanguage)
		getBack()
		a.actionsOfATM()

	case 5:
		fmt.Println("Заберите карту")
		a.checkTheSecondEnteringPincode()
		a.theFirstEntering()
	}
}

func (a *ATM) checkTheFirstEnteringPincode() {
	fmt.Println("Введите пинкод")

}

func (a *ATM) checkTheSecondEnteringPincode() {

}

func (a *ATM) theFirstEntering() {
	fmt.Println(a.balance)

	fmt.Println(a.balanceOfСlient)

	fmt.Println("Вставьте карту или приложите к NFC чипу \n1 - вставить карту или приложите к NFC чипу \n2 - поменять язык \n3 - выйти")
	var displayNum int

	fmt.Scan(&displayNum)
	switch displayNum {

	case 1:
		fmt.Println("Идёт считывание карты")
		a.checkTheFirstEnteringPincode()
		fmt.Println("Добро пожаловать ʕ ᵔᴥᵔ ʔ Сладенький (ﾉ◕ヮ◕)ﾉ*:･ﾟ✧✧✧✧✧✧✧✧✧✧✧✧✧✧✧✧✧✧✧✧✧✧")
		a.actionsOfATM()

	case 2:
		var enterNumberOfLanguage int
		fmt.Println("1 - поменять на русский\n2 - поменять на английский")
		fmt.Scan(&enterNumberOfLanguage)
		a.ChangeLanguage(enterNumberOfLanguage)
		a.theFirstEntering()

	case 3:
		fmt.Println("До свидания")
	}
}

func getBack() { // функциональность кнопки назад
	fmt.Println("1 - Назад")

	var num int
	fmt.Scan(&num)

	if num == 1 {
		return
	}
}

func initializeATM() ATM {
	atm := ATM{
		language:        languageRussian,
		banknoteToCount: make(map[int]int),
		Clients:         []Clients{},
	}

	atm.banknoteToCount[100] = 50
	atm.banknoteToCount[500] = 50
	atm.banknoteToCount[1000] = 50
	atm.banknoteToCount[2000] = 50
	atm.banknoteToCount[5000] = 50

	for banknote, count := range atm.banknoteToCount {
		atm.balance += banknote * count
	}

	return atm
}

func main() {
	atm := initializeATM()
	atm.theFirstEntering()
}
