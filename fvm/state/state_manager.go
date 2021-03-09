package state

type StateManager struct {
	startState  *State
	activeState *State
}

func NewStateManager(startState *State) *StateManager {
	return &StateManager{
		startState:  startState,
		activeState: startState,
	}
}

func (s *StateManager) State() *State {
	return s.activeState
}

func (s *StateManager) Nest() {
	s.activeState = s.activeState.NewChild()
}

func (s *StateManager) RollUp(merge bool) error {
	var err error
	// TODO merge the register touches
	if merge {
		err = s.activeState.parent.MergeState(s.activeState)
	} else {
		err = s.activeState.parent.MergeTouchLogs(s.activeState)
	}
	if err != nil {
		return err
	}
	// otherwise ignore for now
	if s.activeState.parent != nil {
		s.activeState = s.activeState.parent
	}
	return nil
}

func (s *StateManager) RollUpAll(merge bool) error {
	for {
		if s.activeState == s.startState || s.activeState.parent == nil {
			break
		}
		err := s.RollUp(merge)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *StateManager) ApplyStartStateToLedger() error {
	return s.startState.ApplyDeltaToLedger()
}