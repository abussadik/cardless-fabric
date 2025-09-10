package main

import (
  "encoding/json"
  "fmt"
  "time"
  "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct{ contractapi.Contract }

type Wallet struct{ OwnerID string; Balance int64; WalletID string }
type Tx struct{ TxID, Type, From, To, Currency, Status, Timestamp string; Amount int64 }

func (s *SmartContract) CreateWallet(ctx contractapi.TransactionContextInterface, ownerID, walletID string) error {
  key := "WALLET_"+walletID
  if b,_ := ctx.GetStub().GetState(key); b != nil { return fmt.Errorf("wallet exists") }
  w := Wallet{OwnerID: ownerID, Balance: 0, WalletID: walletID}
  b,_ := json.Marshal(w); return ctx.GetStub().PutState(key, b)
}
func (s *SmartContract) ProcessPayment(ctx contractapi.TransactionContextInterface, txID, txType, fromW, toW, currency string, amount int64) error {
  if amount <= 0 { return fmt.Errorf("invalid amount") }
  fw, err := s.getWallet(ctx, fromW); if err != nil { return err }
  tw, err := s.getWallet(ctx, toW);   if err != nil { return err }
  if fw.Balance < amount { return fmt.Errorf("insufficient funds") }
  fw.Balance -= amount; tw.Balance += amount
  if err := s.putWallet(ctx, fw); err != nil { return err }
  if err := s.putWallet(ctx, tw); err != nil { return err }
  tx := Tx{TxID: txID, Type: txType, From: fromW, To: toW, Amount: amount, Currency: currency, Status: "CONFIRMED", Timestamp: time.Now().UTC().Format(time.RFC3339)}
  tb,_ := json.Marshal(tx); if err := ctx.GetStub().PutState("TX_"+txID, tb); err != nil { return err }
  return ctx.GetStub().SetEvent("PaymentConfirmed", tb)
}
func (s *SmartContract) GetBalance(ctx contractapi.TransactionContextInterface, walletID string) (int64, error) {
  w, err := s.getWallet(ctx, walletID); if err != nil { return 0, err }
  return w.Balance, nil
}
func (s *SmartContract) getWallet(ctx contractapi.TransactionContextInterface, walletID string) (*Wallet, error) {
  b, err := ctx.GetStub().GetState("WALLET_"+walletID); if err != nil || b == nil { return nil, fmt.Errorf("wallet not found") }
  var w Wallet; _ = json.Unmarshal(b, &w); return &w, nil
}
func (s *SmartContract) putWallet(ctx contractapi.TransactionContextInterface, w *Wallet) error {
  b,_ := json.Marshal(w); return ctx.GetStub().PutState("WALLET_"+w.WalletID, b)
}
func main(){ cc,_ := contractapi.NewChaincode(new(SmartContract)); if err := cc.Start(); err != nil { panic(err) } }
