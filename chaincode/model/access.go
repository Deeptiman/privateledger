package model


type LedgerAccess int32

const (
	ALL					LedgerAccess = 1
	WRITE				LedgerAccess = 2	
	READ 				LedgerAccess = 3
	DELETE				LedgerAccess = 4
	REMOVEACCESS		LedgerAccess = 5
	NOACCESS 			LedgerAccess = 6	 
)

func(access LedgerAccess) String() string {
	modes := [...]string{"", "ALL","WRITE", "READ","DELETE", "REMOVE ACCESS","NO ACCESS"}

	return modes[access]
}

func(access LedgerAccess) Int() int32 {
	modes := [...]int32{0,1,2,3,4,5,6}

	return modes[access]
}

func(access LedgerAccess) NoAccess() bool {
	switch access {
		case NOACCESS:
			return true
		default:
			return false
	}
}


func(access LedgerAccess) AllAccess() bool {
	switch access {
		case ALL:
			return true
		default:
			return false
	}
}


func(access LedgerAccess) ReadAccess() bool {
	switch access {
		case READ:
			return true
		default:
			return false
	}
}

func(access LedgerAccess) WriteAccess() bool {
	switch access {
		case WRITE:
			return true
		default:
			return false
	}
}

func(access LedgerAccess) DeleteAccess() bool {
	switch access {
		case DELETE:
			return true
		default:
			return false
	}
}

func(access LedgerAccess) RemoveAccess() bool {
	switch access {
		case REMOVEACCESS:
			return true
		default:
			return false
	}
}