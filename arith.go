
package main

import "errors"

type Args struct {
	X, Y string
}


type Arith struct{}	

var m = NewMap()

func (t *Arith) Create(args *Args, reply *string) error{
	p1 := args.X
	p2 := args.Y
	
	_, ok :=  m.GetMap(p1)
	if ok {
		return errors.New("First Key exists, please use set")
	}
	
	p2Num, ok := check(p2)
	
	if !ok {
		return errors.New("Second value is not valid number")
	}
	
	m.SetMap(p1, p2Num)
	
	*reply = "Create successfully ."
	return nil
}


func (t *Arith) Delete(args *Args, reply *string) error{
	p1 := args.X
	_, ok :=  m.GetMap(p1)
	if !ok {
		*reply = "Key not exist, no need to delete."
		return nil
	}
	m.DeleteMap(p1)
	*reply = "Delete successfully."
	return nil
}

func (t *Arith) Set(args *Args, reply *string) error{
	p1 := args.X
	p2 := args.Y
	_, ok :=  m.GetMap(p1)
	if !ok {
		return errors.New("First Key haven't been created, please create First")
	}
	p2Num, ok := check(p2)
	if !ok {
		return errors.New("Second value is not valid number")
	}
	m.SetMap(p1, p2Num)
	*reply = "Set successfully, now " + p1 + " is " + p2 + " ."
	return nil
}



func (t *Arith) Add(args *Args, reply *string) error{
	p1 := args.X
	p2 := args.Y
	p1Num, ok := m.GetMap(p1)
	if !ok {
		return errors.New("First Key haven't been created, please create First")
	}
	p2Num, ok := m.GetMap(p2)
	if !ok{
		return errors.New("Second Key haven't been created, please create First")
	}
	*reply = addNum(p1Num, p2Num)
	return nil
	
}


func (t *Arith) Subtract(args *Args, reply *string) error{
	p1 := args.X
	p2 := args.Y
	p1Num, ok := m.GetMap(p1)
	if !ok {
		return errors.New("First Key haven't been created, please create First")
	}
	p2Num, ok := m.GetMap(p2)
	if !ok{
		return errors.New("Second Key haven't been created, please create First")
	}
	*reply = subNum(p1Num, p2Num)
	return nil
}

func (t *Arith) Multiply(args *Args, reply *string) error{
	p1 := args.X
	p2 := args.Y
	p1Num, ok := m.GetMap(p1)
	if !ok {
		return errors.New("First Key haven't been created, please create First")
	}
	p2Num, ok := m.GetMap(p2)
	if !ok{
		return errors.New("Second Key haven't been created, please create First")
	}
	*reply = mutNum(p1Num, p2Num)
	return nil
}

func (t *Arith) Divide(args *Args, reply *string) error{
	p1 := args.X
	p2 := args.Y
	p1Num, ok := m.GetMap(p1)
	if !ok {
		return errors.New("First Key haven't been created, please create First")
	}
	p2Num, ok := m.GetMap(p2)
	if !ok{
		return errors.New("Second Key haven't been created, please create First")
	}
	if p2Num.IntNum == "0"{
		return errors.New("Divided by zeros ")
	}
	*reply = divNum(p1Num, p2Num)
	return nil
}
