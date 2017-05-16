package main

import (
  "github.com/hschendel/stl"
  "fmt"
)

func main() {
	obj, errRead := stl.ReadFile("teapot.stl")
	if errRead != nil {
		panic(errRead)
	}

  show_object(toVba(obj))
}

func toVba(obj *stl.Solid) []float32 {
  var r []float32
  for _, triangle := range obj.Triangles {
    for _, vertex := range triangle.Vertices {
      for _, a := range vertex {
        fmt.Printf("%+v\n", a)
        r = append(r, a)
      }
      r = append(r, 1.0)
      r = append(r, 0.0)
    }
  }
  return r
}
