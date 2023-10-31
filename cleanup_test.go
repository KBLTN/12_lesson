package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestExampleCleanup(t *testing.T) {
	x := 0
	Convey("A", t, func() {
		x++ //setup
		Convey("A-B", func() {
			x++
			Convey("A-B-C1", func() {
				So(x, ShouldEqual, 2)
				So(x, ShouldEqual, 2)
				So(x, ShouldEqual, 2)
			})
			Convey("A-B-C2", func() {
				So(x, ShouldEqual, 4)
			})
			Convey("A-B-C3", func() {
				So(x, ShouldEqual, 6)
			})
		})
		Reset(func() {
			t.Log("finish")
		})
	})
}

//func TestExampleCleanup(t *testing.T) {
//	fmt.Println("SETUP")
//
//	t.Cleanup(func() {
//		fmt.Println("TEARDOWN ON CLEANUP")
//	})
//
//	t.Run("FIRST", func(t *testing.T) {
//		fmt.Println("ok")
//	})
//
//	t.Run("SECOND", func(t *testing.T) {
//		fmt.Println("ok")
//	})
//
//	t.Run("THIRD", func(t *testing.T) {
//		fmt.Println("fatal test")
//	})
//
//	fmt.Println("TEARDOWN AT END")
//
//	t.Cleanup(func() {
//		fmt.Println("TEARDOWN ON CLEANUP")
//	})
//}
