package models

import (
	"context"
	"ewallet_be/utils"
	"fmt"
	"time"
)

type Wallet struct {
	IDWallet int     `json:"idwallet" db:"id_wallet"`
	IDUser   int     `json:"iduser" db:"id_user"`
	Balance  float64 `json:"balance"`
}

type Topup struct {
	IDTopup         int       `json:"idtopup" db:"id_topup"`
	IDWallet        int       `json:"idwallet" db:"id_wallet"`
	Amount          float64   `json:"amount"`
	IDPaymentMethod int       `json:"idpaymentmethod" db:"id_payment_method"`
	AdminFee        float64   `json:"adminfee" db:"admin_fee"`
	Tax             float64   `json:"tax"`
	CreatedAt       time.Time `json:"createdat" db:"created_at"`
}

func CreateTopup(topup Topup) error {
	conn, err := utils.ConnectDB()
	if err != nil {
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(context.Background(), `
		INSERT INTO topups (id_wallet, amount, id_payment_method, admin_fee, tax)
		VALUES ($1, $2, $3, $4, $5)
	`, topup.IDWallet, topup.Amount, topup.IDPaymentMethod, topup.AdminFee, topup.Tax)

	if err != nil {
		return err
	}

	_, err = conn.Exec(context.Background(), `
		UPDATE wallets SET balance = balance + $1
		WHERE id_wallet = $2
	`, topup.Amount, topup.IDWallet)

	if err != nil {
		return fmt.Errorf("failed to update wallet balance: %v", err)
	}

	return nil
}
