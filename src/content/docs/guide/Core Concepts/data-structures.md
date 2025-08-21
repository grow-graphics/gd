---
title: Builtin Data Structures
slug: guide/data-structures
sidebar:
  order: 20
---

The gdextension interface to the engine supports a fixed set of data types, all of
them have both an equivalent convenience type and the best performance type (which can be
used to reduce allocations and unnecessary type conversions in hot functions).

All of the data types have a package underneath `graphics.gd/variant`. With the exception
of `Object`, all these packages are implemented in pure Go to avoid any coupling or
overhead when making calls to the engine. You can import these packages in any Go project.


| Engine Type        | Convenience Type          | Best Performance Type           |
| ------------------ | ------------------------- | ------------------------------- |
| Variant            | `any`                     | `variant.Any`                   |
| bool               | `bool`                    | `bool`                          |
| int                | `int`                     | `int64`                         |
| float              | `Float.X`                 | `float64`                       |
| String             | `string`                  | `String.Readable`               |
| Vector2            | `Vector2.XY`              | `Vector2.XY`                    |
| Vector2i           | `Vector2i.XY`             | `Vector2i.XY`                   |
| Rect2              | `Rect2.PositionSize`      | `Rect2.PositionSize`            |
| Rect2i             | `Rect2i.PositionSize`     | `Rect2i.PositionSize`           |
| Vector3            | `Vector3.XYZ`             | `Vector3.XYZ`                   |
| Vector3i           | `Vector3i.XYZ`            | `Vector3i.XYZ`                  |
| Transform2D        | `Transform2D`             | `Transform2D`                   |
| Vector4            | `Vector4.XYZW`            | `Vector4.XYZW`                  |
| Vector4i           | `Vector4i.XYZW`           | `Vector4i.XYZW`                 |
| Plane              | `Plane.NormalD`           | `Plane.NormalD`                 |
| Quaternion         | `Quaternion.IJKL`         | `Quaternion.IJKL`               |
| AABB               | `AABB.PositionSize`       | `AABB.PositionSize`             |
| Basis              | `Basis.XYZ`               | `Basis.XYZ`                     |
| Transform3D        | `Transform3D.BasisOrigin` | `Transform3D.BasisOrigin`       |
| Projection         | `Projection.XYZW`         | `Projection.XYZW`               |
| Color              | `Color.RGBA`              | `Color.RGBA`                    |
| StringName         | `string`                  | `String.Name`                   |
| NodePath           | `string`                  | `Path.ToNode`                   |
| Signal             | `chan T`                  | `Signal.Any`                    |
| RID                | `RID.T`                   | `RID.Any`                       |
| Object             | `T.Instance`              | `T.Advanced`                    |
| Callable           | `func(...T) (...T)`       | `Callable.Function`             |
| Dictionary         | `struct/map[T]T`          | `Dictionary.Any`                |
| Array              | `[]T`                     | `Array.Any`                     |
| PackedByteArray    | `[]byte`                  | `Packed.Bytes`                  |
| PackedInt32Array   | `[]int32`                 | `Packed.Array[int32]`           |
| PackedInt64Array   | `[]int64`                 | `Packed.Array[int64]`           |
| PackedFloat32Array | `[]float32`               | `Packed.Array[float32]`         |
| PackedFloat64Array | `[]float64`               | `Packed.Array[float64]`         |
| PackedStringArray  | `[]string`                | `Packed.Strings`                |
| PackedVector2Array | `[]Vector2.XY`            | `Packed.Array[Vector2.XY]`      |
| PackedVector3Array | `[]Vector3.XYZ`           | `Packed.Array[Vector3.XYZ]`     |
| PackedColorArray   | `[]Color.RGBA`            | `Packed.Array[Color.RGBA]`      |
| PackedVector4Array | `[]Vector4.XYZW`          | `Packed.Array[Vector4.XYZW]`    |
