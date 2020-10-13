package kvdb

import "time"

type Nop struct{}

//Close close database
func (n Nop) Close() error {
	return nil
}

//Set set value by given key
func (n Nop) Set(key string, value []byte) error {
	return ErrFeatureNotSupported
}

//Get get value by given key
func (n Nop) Get(key string) ([]byte, error) {
	return nil, ErrFeatureNotSupported
}

//Delete delete value by given key
func (n Nop) Delete(key string) error {
	return ErrFeatureNotSupported
}

//Next return values after key not more than given limit
func (n Nop) Next(key string, limit int) ([][]byte, error) {
	return nil, ErrFeatureNotSupported
}

//Prev return values before key not more than given limit
func (n Nop) Prev(key string, limit int) ([][]byte, error) {
	return nil, ErrFeatureNotSupported
}

//SetWithTTL set value by given key and ttl
func (n Nop) SetWithTTL(key string, ttl time.Duration) error {
	return ErrFeatureNotSupported
}

//Begin begin new transaction
func (n Nop) Begin() (Transaction, error) {
	return nil, ErrFeatureNotSupported
}

//Features return supported features
func (n Nop) Features() Feature {
	return 0
}
