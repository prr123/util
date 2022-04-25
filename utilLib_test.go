// utilLib_test.go
// test program for utilLib.go
//
// author: prr
// date: 24/4/2022
// copyright 2022 prr azul software
//
package utilLib

import (
//    "fmt"
    "testing"
	"os"
//	util "utility/utilLib"

)

func TestCreateFileFolder(t *testing.T) {

	fpath, exist, err := CreateFileFolder("test", "new")
	if err != nil {
		t.Error("should be nil!")
	}
	if exist {
		t.Error("exist should be false!")
	}

	if fpath  != "test/new" {
		t.Error("fpath is wrong!")
	}

	finfo, err :=  os.Stat(fpath)
	if os.IsNotExist(err) {
		t.Error("fpath should exist!")
	}

	if !finfo.IsDir() {
		t.Error("fpath should be dir!")
	}

	fpath, exist, err = CreateFileFolder("test", "new")
	if err != nil {
		t.Error("should be nil!")
	}

	if !exist {
		t.Error("exist should be true!")
	}
	os.RemoveAll(fpath)

}

func TestCreateFileFolder2(t *testing.T) {


	fpath, exist, err := CreateFileFolder("test/test1/test2", "new")
	if err != nil {
		t.Error("should be nil!")
	}
	if exist {
		t.Error("exist should be false!")
	}

	if fpath  != "test/test1/test2/new" {
		t.Error("fpath is wrong!")
	}

	finfo, err :=  os.Stat("test/test1/test2/new")
	if os.IsNotExist(err) {
		t.Error("test/test1/test2/new should exist!")
	}

	if !finfo.IsDir() {
		t.Error("test/test1/test2/new should be dir!")
	}

	os.RemoveAll("test")

}
