open System

let input =
    """1000
2000
3000

4000

5000
6000

7000
8000
9000

10000"""

let elfCalories =
    input.Split "\n\n"
    |> Array.map (fun s -> s.Split "\n" |> Array.map (int))
    |> Array.map (Array.fold (+) 0)

let res = elfCalories |> Array.max

printfn $"Part 1 : {res}"

let res2 = elfCalories |> Array.sortDescending |> Array.take 3 |> Array.sum
printfn $"Part 2 {res2}"
