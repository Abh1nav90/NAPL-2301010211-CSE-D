open System

// The base class
// shapes can replace with their own calculation.
[<AbstractClass>]
type Shape() =
    abstract member Area : unit -> float
    default this.Area() = 0.0

// A circle calculates its area from its radius.
type Circle(radius: float) =
    inherit Shape()

    override _.Area() =
        Math.PI * radius * radius

// A rectangle calculates its area from its two sides.
type Rectangle(length: float, breadth: float) =
    inherit Shape()

    override _.Area() =
        length * breadth

[<EntryPoint>]
let main _ =
    let gardenCircle = Circle(3.5)
    let floorRectangle = Rectangle(8.0, 6.0)

    printfn "Shape area demonstration"
    printfn "Circle area: %.2f" (gardenCircle.Area())
    printfn "Rectangle area: %.2f" (floorRectangle.Area())

    // Both objects are handled through the common Shape type.
    let allShapes: Shape list =
        [ gardenCircle :> Shape
          floorRectangle :> Shape ]

    printfn "\nAreas calculated through the base-class reference:"

    allShapes
    |> List.iteri (fun index shape ->
        printfn "Shape %d area: %.2f" (index + 1) (shape.Area()))

    0
