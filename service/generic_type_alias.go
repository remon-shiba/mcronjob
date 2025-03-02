package service

type Age int

func (a Age) Valid() bool { return a > 0 && a < 100 }
