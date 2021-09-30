package models

type Machine struct {
	ID     string
	Name   string
	Status string
}

var machines = []Machine{
	{ID: "M1", Name: "Machine1", Status: "Green"},
	{ID: "M2", Name: "Machine2", Status: "Yellow"},
	{ID: "M3", Name: "Machine3", Status: "Blue"},
	{ID: "M4", Name: "Machine4", Status: "Green"},
}

func GetAllMachines() ([]Machine, error) {
	return machines, nil
}

func AddNewMachine(m Machine) error {
	machines = append(machines, m)
	return nil
}
