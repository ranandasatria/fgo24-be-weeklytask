package models

import (
	"context"
	"ewallet_be/utils"
)

type Wallet struct {
	IDWallet int     `json:"idWallet" db:"id_wallet"`
	IDUser   int     `json:"idUser" db:"id_user"`
	Balance  float64 `json:"balance"`
}

func CreateWalletForUser(userID int) error {
	conn, err := utils.ConnectDB()
	if err != nil {
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(context.Background(),
		`INSERT INTO wallets (id_user, balance) VALUES ($1, 0)`,
		userID,
	)
	return err
}


func GetWalletByUserID(userID int) (Wallet, error) {
	conn, err := utils.ConnectDB()
	if err != nil {
		return Wallet{}, err
	}
	defer conn.Release()

	var wallet Wallet
	err = conn.QueryRow(
		context.Background(),
		`SELECT id_wallet, id_user, balance FROM wallets WHERE id_user = $1`,
		userID,
	).Scan(&wallet.IDWallet, &wallet.IDUser, &wallet.Balance)

	if err != nil {
		return Wallet{}, err
	}

	return wallet, nil
}
