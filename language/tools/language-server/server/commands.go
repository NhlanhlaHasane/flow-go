package server

import (
	"context"
	"errors"
	"fmt"

	"github.com/dapperlabs/flow-go/model/flow"
	"github.com/dapperlabs/flow-go/sdk/keys"

	"github.com/dapperlabs/flow-go/language/tools/language-server/protocol"
)

const (
	CommandSubmitTransaction = "cadence.submitTransaction"
)

// CommandHandler represents the form of functions that handle commands
// submitted from the client using workspace/executeCommand.
type CommandHandler func(conn protocol.Connection, args ...interface{}) (interface{}, error)

// Registers the commands that the server is able to handle.
//
// The best reference I've found for how this works is:
// https://stackoverflow.com/questions/43328582/how-to-implement-quickfix-via-a-language-server
func (s Server) registerCommands(connection protocol.Connection) {
	// Send a message to the client indicating which commands we support
	err := connection.RegisterCapability(&protocol.RegistrationParams{
		Registrations: []protocol.Registration{
			{
				ID:     "test",
				Method: "workspace/executeCommand",
				RegisterOptions: protocol.ExecuteCommandRegistrationOptions{
					ExecuteCommandOptions: protocol.ExecuteCommandOptions{
						Commands: []string{CommandSubmitTransaction},
					},
				},
			},
		},
	})
	if err != nil {
		connection.LogMessage(&protocol.LogMessageParams{
			Type:    protocol.Warning,
			Message: fmt.Sprintf("Failed to register command: %s", err.Error()),
		})
	}

	// Register each command handler function in the server
	s.commands[CommandSubmitTransaction] = s.submitTransaction
}

// submitTransaction handles submitting a transaction defined in the
// source document in VS Code.
//
// There should be exactly 1 argument, the DocumentURI of the file to submit.
func (s Server) submitTransaction(connection protocol.Connection, args ...interface{}) (interface{}, error) {
	connection.ShowMessage(&protocol.ShowMessageParams{
		Type:    protocol.Info,
		Message: fmt.Sprintf("called submit transaction %v", args),
	})

	if len(args) != 1 {
		return nil, errors.New("missing argument")
	}
	uri, ok := args[0].(string)
	if !ok {
		return nil, errors.New("invalid uri argument")
	}
	documentText, ok := s.documents[protocol.DocumentUri(uri)]
	if !ok {
		return nil, fmt.Errorf("could not find document for URI %s", uri)
	}

	tx := flow.Transaction{
		Script:         []byte(documentText),
		Nonce:          s.getNextNonce(),
		ComputeLimit:   10,
		PayerAccount:   s.config.AccountAddr,
		ScriptAccounts: []flow.Address{s.config.AccountAddr},
	}

	sig, err := keys.SignTransaction(tx, s.config.AccountKey)
	if err != nil {
		return nil, err
	}
	tx.AddSignature(s.config.AccountAddr, sig)

	err = s.flowClient.SendTransaction(context.Background(), tx)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
