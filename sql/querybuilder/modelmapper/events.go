package modelmapper

//CommonQueryEvents common query events struct
type CommonQueryEvents struct {
}

//BeforeInsert event raised before insert
func (e CommonQueryEvents) BeforeInsert() error {
	return nil
}

//AfterInsert event raised after insert
func (e CommonQueryEvents) AfterInsert() error {
	return nil
}

//BeforeUpdate event raised before update
func (e CommonQueryEvents) BeforeUpdate() error {
	return nil
}

//AfterUpdate event raised after update
func (e CommonQueryEvents) AfterUpdate() error {
	return nil
}

//AfterFind event raised before find
func (e CommonQueryEvents) AfterFind() error {
	return nil
}

//AfterDelete event raised after find
func (e CommonQueryEvents) AfterDelete() error {
	return nil
}

//AfterMarshal event raised before find
func (e CommonQueryEvents) AfterMarshal() error {
	return nil
}

//BeforeMarshal event raised after find
func (e CommonQueryEvents) BeforeMarshal() error {
	return nil
}

//AfterUnmarshal event raised before find
func (e CommonQueryEvents) AfterUnmarshal() error {
	return nil
}

//BeforeUnmarshal event raised after find
func (e CommonQueryEvents) BeforeUnmarshal() error {
	return nil
}
