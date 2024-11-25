package runner

func connect(address string) error {
	if err := polling(address); err != nil {
		return err
	}
	if err := checkHealth(address); err != nil {
		return err
	}
	return nil
}
