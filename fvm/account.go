package fvm

import (
	"errors"
	"fmt"

	"github.com/dapperlabs/flow-go/fvm/state"
	"github.com/dapperlabs/flow-go/model/flow"
)

func getAccount(
	vm *VirtualMachine,
	ctx Context,
	ledger state.Ledger,
	chain flow.Chain,
	address flow.Address,
) (*flow.Account, error) {
	addresses := state.NewAddresses(ledger, chain)
	accounts := state.NewAccounts(ledger, addresses)

	account, err := accounts.Get(address)
	if err != nil {
		if errors.Is(err, state.ErrAccountNotFound) {
			return nil, ErrAccountNotFound
		}

		return nil, err
	}

	var result *InvocationResult
	result, err = vm.Invoke(
		ctx,
		getFlowTokenBalanceScript(address, chain.ServiceAddress()),
		ledger,
	)
	if err != nil {
		return nil, err
	}

	var balance uint64

	// TODO: Figure out how to handle this error. Currently if a runtime error occurs, balance will be 0.
	// 1. An error will occur if user has removed their FlowToken.Vault -- should this be allowed?
	// 2. Any other error indicates a bug in our implementation. How can we reliably check the Cadence error?
	if result.Error == nil {
		balance = result.Value.ToGoValue().(uint64)
	}

	account.Balance = balance

	return account, nil
}

const initFlowTokenTransactionTemplate = `
import FlowServiceAccount from 0x%s

transaction {
  prepare(account: AuthAccount) {
    FlowServiceAccount.initDefaultToken(account)
  }
}
`

const getFlowTokenBalanceScriptTemplate = `
import FlowServiceAccount from 0x%s

pub fun main(): UFix64 {
  let acct = getAccount(0x%s)
  return FlowServiceAccount.defaultTokenBalance(acct)
}
`

func initFlowTokenTransaction(accountAddress, serviceAddress flow.Address) InvokableTransaction {
	return Transaction(
		flow.NewTransactionBody().
			SetScript([]byte(fmt.Sprintf(initFlowTokenTransactionTemplate, serviceAddress))).
			AddAuthorizer(accountAddress),
	)
}

func getFlowTokenBalanceScript(accountAddress, serviceAddress flow.Address) InvokableScript {
	return Script([]byte(fmt.Sprintf(getFlowTokenBalanceScriptTemplate, serviceAddress, accountAddress)))
}
