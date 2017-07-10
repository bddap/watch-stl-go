package main

import (
  "github.com/hschendel/stl"
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
        r = append(r, a)
      }
    }
  }
  return r
}
