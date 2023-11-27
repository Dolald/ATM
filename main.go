package main

import (
	"fmt"
	"math/rand"
	"time"
)

type ATM struct {
	balanceOfATM    int
	balanceOfСlient int
	whatIsLanguage  bool
}

func (c *ATM) ChangeLanguage() {
	if !c.whatIsLanguage {
		fmt.Println("\n1 - sпоменять на русский\n2 - поменять на английский\n")
	} else {
		fmt.Println("\n1 - change the language to Russian\n2 - change the language to English\n")
	}
	var scaned int
	fmt.Scan(&scaned)

	if scaned == 1 {
		c.whatIsLanguage = false
	} else if scaned == 2 {
		c.whatIsLanguage = true
	}
}

func (c *ATM) SetClientBalance() {
	c.balanceOfСlient = 1000 + rand.Intn(49000)
}

func (c *ATM) CheckBalance() {
	time.Sleep(1 * time.Second)
	if !c.whatIsLanguage {
		fmt.Println("Ваш баланс:", c.balanceOfСlient)
	} else {
		fmt.Println("Your balance:", c.balanceOfСlient)
	}
}

func (c *ATM) Deposit() {
	var deposit int
	if !c.whatIsLanguage {
		fmt.Println("Введите сумму для пополнения")
	} else {
		fmt.Println("Enter the amount for replenishment")
	}
	fmt.Scan(&deposit)

	c.balanceOfATM += deposit
	c.balanceOfСlient += deposit

	if !c.whatIsLanguage {
		fmt.Print("Идёт распознавание купюр\n")
	} else {
		fmt.Print("Recognition of banknotes is on processing")
	}

	sleep()

	if !c.whatIsLanguage {
		fmt.Println("\nВы успешно пополнили:", deposit)
	} else {
		fmt.Println("\nYou have successfully replenished:", deposit)
	}
}

func sleep() {
	fmt.Print("Подождите")
	time.Sleep(1 * time.Second)
	fmt.Print(". ")
	time.Sleep(1 * time.Second)
	fmt.Print(". ")
	time.Sleep(1 * time.Second)
	fmt.Print(".")
	time.Sleep(2 * time.Second)
}

func (c *ATM) SetBalance() {
	for {
		n := 10000 + rand.Intn(90000)
		if n%100 == 0 {
			c.balanceOfATM = n
			break
		}
	}
	//fmt.Println("Банкомат имеет столько денег:", c.balanceOfATM)
}

func (c *ATM) Withdraw() {
	var withdrow int
	if !c.whatIsLanguage {
		fmt.Println("Введите сумму для снятия наличных")
	} else {
		fmt.Println("Enter the amount to withdrow cash")
	}
	fmt.Scan(&withdrow)

	if withdrow%10 == 0 {

		if withdrow <= c.balanceOfATM {
			if IsPinCodeCorrectForFunctionalaty(c.whatIsLanguage) {
				c.balanceOfATM -= withdrow
				c.balanceOfСlient -= withdrow

				if !c.whatIsLanguage {
					fmt.Print("\nИдёт распознавание купюр\n")
				} else {
					fmt.Print("\nRecognition of banknotes is on processing")
				}

				sleep()

				if !c.whatIsLanguage {
					fmt.Println("\nВы успешно сняли", withdrow)
				} else {
					fmt.Println("\nYou have successfully withdrowed:", withdrow)
				}
			} else {
				main()
			}
		} else {
			if !c.whatIsLanguage {
				fmt.Println("У банкомата недостаточно средств для выдачи такой суммы")
			} else {
				fmt.Println("ATM has so much money:", c.balanceOfСlient)
			}
			return
		}
	} else {
		if !c.whatIsLanguage {
			fmt.Println("Банкомата не выдаёт такие суммы, выберите другую сумму")
		} else {
			fmt.Println("ATM doesn't issue such amounts, enter other amount", c.balanceOfСlient)
		}
	}
}
func RussOrEngForFunctionality(stract ATM) {
	if !stract.whatIsLanguage {
		display2()
	} else {
		displayEng()
	}
}
func functionality(stract ATM) {
	var num int
	if !stract.whatIsLanguage {
		display2()
	} else {
		displayEng()
	}
	for {
		fmt.Scan(&num)

		switch num {
		case 1:
			stract.Withdraw()
			back(stract)
			RussOrEngForFunctionality(stract)

		case 2:
			stract.Deposit()
			back(stract)
			RussOrEngForFunctionality(stract)

		case 3:
			stract.CheckBalance()
			back(stract)
			RussOrEngForFunctionality(stract)

		case 4:
			stract.ChangeLanguage()
			back(stract)
			RussOrEngForFunctionality(stract)
		}

		if num == 5 {
			if !stract.whatIsLanguage {
				display()
			} else {
				displayEng1()
			}
			var i int
			fmt.Scan(&i)

			switch i {
			case 1:
				if IsPinCodeCorrect(stract) {
					functionality(stract)
				}

			case 2:
				time.Sleep(1 * time.Second)
				stract.ChangeLanguage()

				if !stract.whatIsLanguage {
					display()
				} else {
					displayEng1()
				}
				funcOfDisplay(stract)

			case 3:
				break
			}
			break
		}
	}
}

func funcOfDisplay(stract ATM) {
	stract.SetBalance()
	stract.SetClientBalance()

	var displayNum int
	fmt.Scan(&displayNum)

	switch displayNum {
	case 1:
		if !stract.whatIsLanguage {
			fmt.Println("Идёт считывание карты")
			sleep()
		} else {
			fmt.Println("Recognition of banknotes is on processing")

		}
		time.Sleep(1 * time.Second)
		if IsPinCodeCorrect(stract) {
			functionality(stract)
		}
	case 2:
		stract.whatIsLanguage = false
		stract.ChangeLanguage()
		if !stract.whatIsLanguage {
			display()
		} else {
			displayEng1()
		}
		funcOfDisplay(stract)
	case 3:
		break
	}
}

func IsPinCodeCorrect(stract ATM) bool {
	var pincode, mistake int

	for {
		if !stract.whatIsLanguage {
			fmt.Println("\nВведите пинкод")
		} else {
			fmt.Println("\nEnter the pincode")
		}
		fmt.Scan(&pincode)

		if pincode != 1973 {
			mistake++

			switch mistake {
			case 1:
				if !stract.whatIsLanguage {
					fmt.Println("\nНе правильный пинкод")
				} else {
					fmt.Println("The pincode is incorrect")
				}
			case 2:
				if !stract.whatIsLanguage {
					fmt.Println("\nНе правильный пинкод")
				} else {
					fmt.Println("The pincode is incorrect")
				}
			}

			if mistake == 3 {
				if !stract.whatIsLanguage {
					fmt.Println("\nВаша карта заблокирована :-<\n\nРазблокировать можно по номеру - 8-965-460-43-48\n")
				} else {
					fmt.Println("Your card is blocked :-<\n\nTo unblock you should call - 8-965-460-43-48\n ")
				}
				return false
			}
		} else {
			time.Sleep(1 * time.Second)
			if !stract.whatIsLanguage {
				fmt.Println("Добро пожаловать ʕ ᵔᴥᵔ ʔ Сладенький (ﾉ◕ヮ◕)ﾉ*:･ﾟ✧✧✧✧✧✧✧✧✧✧✧✧✧✧✧✧✧✧✧✧✧✧")
			} else {
				fmt.Println("Wellcome ʕ ᵔᴥᵔ ʔ sweet (ﾉ◕ヮ◕)ﾉ*:･ﾟ✧✧✧✧✧✧✧✧✧✧✧✧✧✧✧✧✧✧✧✧✧✧")
			}
			return true
		}
	}
}

func IsPinCodeCorrectForFunctionalaty(whatIsLanguage bool) bool {
	if !whatIsLanguage {
		fmt.Println("Введите пинкод")
	} else {
		fmt.Println("Enter the pincode")
	}

	var pincode int
	fmt.Scan(&pincode)

	if pincode != 1973 {
		if !whatIsLanguage {
			fmt.Println("Не правильный пинкод")
		} else {
			fmt.Println("The pincode is incorrect")
		}
		return false
	}
	return true
}

func back(stract ATM) {
	if !stract.whatIsLanguage {
		fmt.Println("1 - Назад")
	} else {
		fmt.Println("1 - back")
	}

	var i int
	fmt.Scan(&i)

	if i == 1 {
		return
	}
}

func display() {
	fmt.Printf("\nВставьте карту или приложите к NFC чипу \n1 - вставить карту или приложите к NFC чипу \n2 - поменять язык \n3 - выйти\n")
}

func displayEng1() {
	fmt.Printf("\nPut your card or apply to the NFC chip \n1 - Put the card or apply to the NFC chip \n2 - change the language \n3 - exit\n")
}

func display2() {
	fmt.Println("\n1 - снять наличные \n2 - пополнить карту \n3 - посмотреть баланс \n4 - поменять язык \n5 - вытащить карту")
}

func displayEng() {
	fmt.Println("\n1 - withdrow money \n2 - top up the card \n3 - check the balance \n4 - change the language \n5 - take back the card")
}

func main() {
	stract := ATM{}
	display()
	funcOfDisplay(stract)
}
