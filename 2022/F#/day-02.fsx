open System

let input =
    """A Y
B X
C Z"""

// First is opp
// Second is you

let (|Rock|Paper|Scissors|) input =
    match input with
    | 'A'
    | 'X' -> Rock
    | 'B'
    | 'Y' -> Paper
    | 'C'
    | 'Z' -> Scissors
    | _ -> raise (Exception "unknown type")


let shapeScore =
    function
    | Rock -> 1
    | Paper -> 2
    | Scissors -> 3

// 0 loss, 3 win, 6 win
let score (a, b) =
    match (a, b) with
    | (Rock, Scissors) -> 0
    | (Rock, Rock) -> 3
    | (Rock, Paper) -> 6
    | (Scissors, Paper) -> 0
    | (Scissors, Scissors) -> 3
    | (Scissors, Rock) -> 6
    | (Paper, Rock) -> 0
    | (Paper, Paper) -> 3
    | (Paper, Scissors) -> 6

let res =
    input.Split "\n"
    |> Array.map (fun l -> (l[0], l[2]))
    |> Array.map (fun (opp, me) -> (score (opp, me)) + (shapeScore me))
    |> Array.reduce (+)

printfn $"Part 1: {res}"

//TODO Part 2
