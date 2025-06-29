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

type Transfer struct {
	IDTransfer      int       `json:"idtransfer" db:"id_transfer"`
	IDSenderWallet  int       `json:"idsenderwallet" db:"id_sender_wallet"`
	IDReceiverWallet int      `json:"idreceiverwallet" db:"id_receiver_wallet"`
	Amount          float64   `json:"amount"`
	Notes           string    `json:"notes"`
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


func CreateTransfer(t Transfer) error {
	conn, err := utils.ConnectDB()
	if err != nil {
		return err
	}
	defer conn.Release()

	var currentBalance float64
	err = conn.QueryRow(context.Background(),
		`SELECT balance FROM wallets WHERE id_wallet = $1`, t.IDSenderWallet).
		Scan(&currentBalance)
	if err != nil {
		return fmt.Errorf("failed to fetch sender wallet: %v", err)
	}

	if currentBalance < t.Amount {
		return fmt.Errorf("insufficient balance")
	}

	tx, err := conn.Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())

	_, err = tx.Exec(context.Background(),
		`UPDATE wallets SET balance = balance - $1 WHERE id_wallet = $2`,
		t.Amount, t.IDSenderWallet)
	if err != nil {
		return err
	}

	_, err = tx.Exec(context.Background(),
		`UPDATE wallets SET balance = balance + $1 WHERE id_wallet = $2`,
		t.Amount, t.IDReceiverWallet)
	if err != nil {
		return err
	}

	row := tx.QueryRow(context.Background(), `
		INSERT INTO transfers (id_sender_wallet, id_receiver_wallet, amount, notes)
		VALUES ($1, $2, $3, $4)
		RETURNING id_transfer, created_at
	`, t.IDSenderWallet, t.IDReceiverWallet, t.Amount, t.Notes)

	err = row.Scan(&t.IDTransfer, &t.CreatedAt)
	if err != nil {
		return err
	}

	return tx.Commit(context.Background())
}
