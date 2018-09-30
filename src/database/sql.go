package database

const (
	// ContainTable is the swift table create sql
	ContainTable = `CREATE TABLE container (
		id INTEGER PRIMARY KEY,
		name TEXT DEFAULT '',
		put_timestamp TEXT not null DEFAULT '',
		delete_timestamp TEXT not null DEFAULT '',
		object_count INTEGER,
		bytes_used INTEGER,
		deleted INTEGER DEFAULT 0,
		storage_policy_index INTEGER DEFAULT 0
	);`
	// AccountTable is the swift table create sql
	AccountTable = `CREATE TABLE account_stat (	
		account TEXT default '',
		created_at TEXT default '',
		put_timestamp TEXT DEFAULT '',
		delete_timestamp TEXT DEFAULT '',
		container_count INTEGER,
		object_count INTEGER DEFAULT 0,
		bytes_used INTEGER DEFAULT 0,
		hash TEXT NOT NULL DEFAULT '',
		id TEXT DEFAULT '',
		status TEXT DEFAULT '',
		status_changed_at TEXT DEFAULT '',
		metadata TEXT DEFAULT ''
	);`
	// PolicyTable is the swift table create sql
	PolicyTable = `CREATE TABLE policy_stat (
		storage_policy_index INTEGER PRIMARY KEY,
		container_count INTEGER DEFAULT 0,
		object_count INTEGER DEFAULT 0,
		bytes_used INTEGER DEFAULT 0
	);`

)