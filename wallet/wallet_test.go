package wallet

import "testing"

func assertBalance(t *testing.T, wallet Wallet, expect Bitcoin) {
	t.Helper()
	actual := wallet.Balance()
	if actual != expect {
		t.Errorf("expected: %s, actual: %s", expect, actual)
	}
}

func assertError(t *testing.T, err error, expected error) {
	if err == nil {
		t.Fatal("expecting error but got none")
	}

	if err != expected {
		t.Errorf("expecting error %s but got %s", expected, err)
	}
}

func assertNoError(t *testing.T, err error) {
	if err != nil {
		t.Errorf("expecting no error but got %s", err)
	}
}

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		err := wallet.Withdraw(Bitcoin(3))
		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(7))
	})

	t.Run("Withdraw with overdraft", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		err := wallet.Withdraw(Bitcoin(11))
		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, Bitcoin(10))
	})
}
