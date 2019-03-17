package main

import "fmt"
import "time"
import _ "sync"

func main() {
	s2 := &student{}
	funcs := []func(){
		s2.name,
		s2.age,
		s2.birth,
		s2.arr,
		s2.mapp,
	}

	for _, f := range funcs {
		go func(ff func()) {
			f()
		}(f)
	}

	fmt.Println("name:", s2.Name, "age:", s2.Age, "birth:", s2.Birth, "arr:", s2.Arr, "map:", s2.Map, "ref1:", s2.Ref1, "ref2:", s2.Ref2)

	s3 := &student{}

	s3.name()
	s3.age()
	s3.birth()
	s3.arr()
	s3.mapp()

	fmt.Println("name:", s3.Name, "age:", s3.Age, "birth:", s3.Birth, "arr:", s3.Arr, "map:", s3.Map, "ref1:", s3.Ref1, "ref2:", s3.Ref2)
}

type student struct {
	Name  string
	Age   int
	Birth time.Time
	Arr   []string
	Map   map[int]int

	Ref1 *ref
	Ref2 ref
}

type ref struct {
	A int
}

func (s *student) name() {
	s.Name = "xx"
}

func (s *student) age() {
	s.Age = 18
}

func (s *student) birth() {
	s.Birth = time.Now()
}

func (s *student) arr() {
	s.Arr = make([]string, 2)
	s.Arr[0] = "arr1"
	s.Arr[1] = "arr2"
}

func (s *student) mapp() {
	s.Map = make(map[int]int)
	s.Map[1] = 1
	s.Map[2] = 2

	s.Ref1 = &ref{A: 1}
	s.Ref2 = ref{A: 1}
}
