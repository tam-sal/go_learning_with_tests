package pointers

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit positive amount", func(t *testing.T) {
		wallet := Wallet{}
		err := wallet.Deposit(Bitcoin(10))
		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Deposit negative amount", func(t *testing.T) {
		wallet := Wallet{}
		err := wallet.Deposit(Bitcoin(-10))
		assertError(t, err, "the amount must be greater than zero")
	})

	t.Run("Deposit zero amount", func(t *testing.T) {
		wallet := Wallet{}
		err := wallet.Deposit(Bitcoin(0))
		assertError(t, err, "the amount must be greater than zero")
	})

	t.Run("Withdraw valid amount", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(20))
		err := wallet.Withdraw(Bitcoin(10))
		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw more than balance", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(20))
		err := wallet.Withdraw(Bitcoin(30))
		assertError(t, err, "withdrawal amount can't exceed the current balance")
		assertBalance(t, wallet, Bitcoin(20))
	})

	t.Run("Withdraw negative amount", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(20))
		err := wallet.Withdraw(Bitcoin(-10))
		assertError(t, err, "the amount must be greater than zero")
	})

	t.Run("Withdraw zero amount", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(20))
		err := wallet.Withdraw(Bitcoin(0))
		assertError(t, err, "the amount must be greater than zero")
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("Got: %s, Want: %s", got, want)
	}
}

func assertError(t testing.TB, got error, want string) {
	t.Helper()
	if got == nil {
		t.Fatal("Expected an error but didn't get one")
	}

	if got.Error() != want {
		t.Errorf("Got error: %v, Want: %v", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatalf("Unexpected error: %v", got)
	}
}

// ExampleWallet_Deposit shows how to use the Deposit method.
func ExampleWallet_Deposit() {
	wallet := Wallet{}
	err := wallet.Deposit(Bitcoin(10))
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Balance after deposit:", wallet.Balance())
	}
	// Output: Balance after deposit: 10 BTC
}

// ExampleWallet_Deposit_negative shows how the Deposit method handles a negative amount.
func ExampleWallet_Deposit_negative() {
	wallet := Wallet{}
	err := wallet.Deposit(Bitcoin(-10))
	if err != nil {
		fmt.Println("Error:", err)
	}
	// Output: Error: the amount must be greater than zero
}

// ExampleWallet_Deposit_zero shows how the Deposit method handles a zero amount.
func ExampleWallet_Deposit_zero() {
	wallet := Wallet{}
	err := wallet.Deposit(Bitcoin(0))
	if err != nil {
		fmt.Println("Error:", err)
	}
	// Output: Error: the amount must be greater than zero
}

// ExampleWallet_Withdraw shows how to use the Withdraw method.
func ExampleWallet_Withdraw() {
	wallet := Wallet{}
	wallet.Deposit(Bitcoin(20))
	err := wallet.Withdraw(Bitcoin(10))
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Balance after withdrawal:", wallet.Balance())
	}
	// Output: Balance after withdrawal: 10 BTC
}

// ExampleWallet_Withdraw_moreThanBalance shows how the Withdraw method handles withdrawing more than the balance.
func ExampleWallet_Withdraw_moreThanBalance() {
	wallet := Wallet{}
	wallet.Deposit(Bitcoin(20))
	err := wallet.Withdraw(Bitcoin(30))
	if err != nil {
		fmt.Println("Error:", err)
	}
	// Output: Error: withdrawal amount can't exceed the current balance
}

// ExampleWallet_Withdraw_negative shows how the Withdraw method handles a negative amount.
func ExampleWallet_Withdraw_negative() {
	wallet := Wallet{}
	wallet.Deposit(Bitcoin(20))
	err := wallet.Withdraw(Bitcoin(-10))
	if err != nil {
		fmt.Println("Error:", err)
	}
	// Output: Error: the amount must be greater than zero
}

// ExampleWallet_Withdraw_zero shows how the Withdraw method handles a zero amount.
func ExampleWallet_Withdraw_zero() {
	wallet := Wallet{}
	wallet.Deposit(Bitcoin(20))
	err := wallet.Withdraw(Bitcoin(0))
	if err != nil {
		fmt.Println("Error:", err)
	}
	// Output: Error: the amount must be greater than zero
}
