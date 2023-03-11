package test

import "io"

type Tester interface {
	Test()
}

type Opener interface {
	Open() (reader io.Reader, err error)
}
