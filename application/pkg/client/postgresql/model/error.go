package model

import "fmt"

func ErrorCommit(err error) error {
	return fmt.Errorf("failed to commit tx due to error: %v", err)
}

func ErrorRollback(err error) error {
	return fmt.Errorf("failed to rollback tx due to error: %v", err)
}

func ErrorCreateTx(err error) error {
	return fmt.Errorf("failed to create tx due to error: %v", err)
}

func ErrorCreateQuery(err error) error {
	return fmt.Errorf("failed to create sql query due to error: %v", err)
}

func ErrorScan(err error) error {
	return fmt.Errorf("failed to scan due to error: %v", err)
}

func ErrorDoQuery(err error) error {
	return fmt.Errorf("failed to query due to error: %v", err)
}
