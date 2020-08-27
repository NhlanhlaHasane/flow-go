package chunks

import (
	"fmt"

	"github.com/dapperlabs/flow-go/engine/verification/utils"
	chunkmodels "github.com/dapperlabs/flow-go/model/chunks"
	"github.com/dapperlabs/flow-go/model/flow"
	"github.com/dapperlabs/flow-go/module"
)

const (
	// chunkAssignmentAlpha represents number of verification
	// DISCLAIMER: alpha down there is not a production-level value
	chunkAssignmentAlpha = 1
)

// ChunkAssignment handles all the logic related to chunk assignments
type ChunkAssignment struct {
	assigner module.ChunkAssigner
}

// NewChunkAssignment created a new Assignment struct with PublicAssignment as the assigner
func NewChunkAssignment() (*ChunkAssignment, error) {
	// we are using a public assignment scheme here
	assigner, err := NewPublicAssignment(chunkAssignmentAlpha)
	if err != nil {
		return nil, fmt.Errorf("cannot construct new public assignment")
	}

	assignment := &ChunkAssignment{
		assigner: assigner,
	}

	return assignment, nil
}

// TODO: should we check if myID is of the role verification?

// MyChunks creates an assignment using the Execution result to generate an RNG then
// returns the chunks assigned to a specific flow identifier.
func (ca *ChunkAssignment) MyChunks(myID flow.Identifier, verifiers flow.IdentityList, result *flow.ExecutionResult) (flow.ChunkList, error) {
	assignment, err := ca.assign(verifiers, result)
	if err != nil {
		return nil, fmt.Errorf("could not create chunk assignment %w", err)
	}

	// indices of chunks assigned to this node
	chunkIndices := assignment.ByNodeID(myID)

	// mine keeps the list of chunks assigned to this node
	mine := make(flow.ChunkList, 0, len(chunkIndices))
	for _, index := range chunkIndices {
		chunk, ok := result.Chunks.ByIndex(index)
		if !ok {
			return nil, fmt.Errorf("chunk out of range requested: %v", index)
		}

		mine = append(mine, chunk)
	}

	return mine, nil
}

// MyChunksForAssignment ...
func (ca *ChunkAssignment) MyChunksForAssignment(myID flow.Identifier, result *flow.ExecutionResult, assignment chunkmodels.Assignment) (flow.ChunkList, error) {
	// indices of chunks assigned to this node
	chunkIndices := assignment.ByNodeID(myID)

	// mine keeps the list of chunks assigned to this node
	mine := make(flow.ChunkList, 0, len(chunkIndices))
	for _, index := range chunkIndices {
		chunk, ok := result.Chunks.ByIndex(index)
		if !ok {
			return nil, fmt.Errorf("chunk out of range requested: %v", index)
		}

		mine = append(mine, chunk)
	}

	return mine, nil
}

// TODO: should we store the assignment for future calls? This could possible go wrong as
// other execution results are used.

// assign generates the assignment using the execution result to seed the RNG
func (ca *ChunkAssignment) assign(verifiers flow.IdentityList, result *flow.ExecutionResult) (*chunkmodels.Assignment, error) {
	rng, err := utils.NewChunkAssignmentRNG(result)
	if err != nil {
		return nil, fmt.Errorf("could not generate random generator: %w", err)
	}

	assignment, err := ca.assigner.Assign(verifiers, result.Chunks, rng)
	if err != nil {
		return nil, fmt.Errorf("could not create chunk assignment %w", err)
	}

	return assignment, nil
}