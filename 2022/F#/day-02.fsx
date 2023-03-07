open System

//TODO Refactor to handle errors instead of raising exceptions
//TODO Rename functions

let input =
    """A Y
B X
C Z"""

// First is opp
// Second is you

type Throw =
    | Rock
    | Paper
    | Scissors

let getThrow input =
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

let pairs =
    input.Split "\n"
    |> Array.map (fun l -> (l[0], l[2]))

let res =
    pairs
    |> Array.map (fun (a, b) -> (getThrow a, getThrow b))
    |> Array.map (fun (opp, me) -> (score (opp, me)) + (shapeScore me))
    |> Array.reduce (+)

printfn $"Part 1: {res}"

let (|Lose|Win|Draw|) input =
    match input with
    | 'X' -> Lose
    | 'Y' -> Draw
    | 'Z' -> Win
    | _ -> raise (Exception "Invalid Char")

let calcMe =
    function
    | (Rock, Draw) -> Rock
    | (Rock, Win) -> Paper
    | (Rock, Lose) -> Scissors
    | (Scissors, Draw) -> Scissors
    | (Scissors, Win) -> Rock
    | (Scissors, Lose) -> Paper
    | (Paper, Draw) -> Paper
    | (Paper, Win) -> Scissors
    | (Paper, Lose) -> Rock

let calcRound (opp, me) =
    let oppThrow = getThrow opp
    let myThrow = calcMe (oppThrow, me)
    (score (oppThrow, myThrow)) + (shapeScore myThrow)

let res2 = pairs |> Array.map calcRound |> Array.reduce (+)

printfn $"Part 2: {res2}"
