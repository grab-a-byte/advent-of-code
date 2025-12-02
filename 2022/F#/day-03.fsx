open System

let input = """vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw"""

let scores = ['a'..'z'] @ ['A'..'Z']

//Part 1
let splitToSets n = (List.splitInto n) >> (fun l -> ((l[0] |> Set.ofList), (l[1] |> Set.ofList)))
let weighting ch = (List.findIndex (fun a -> a = ch) scores) + 1
let seperator = Environment.NewLine //'\010'

let x = input.Split seperator
        |> Array.map List.ofSeq
        |> Array.map (splitToSets 2)
        |> Array.map (fun (a, b) -> Set.intersect a b)
        |> Array.map (Set.toList >> List.head)
        |> Array.map weighting
        |> Array.reduce (+)

printfn $"Part 1 : {x}"

//Part 2
let y = input.Split seperator
        |> Array.map List.ofSeq
        |> Array.chunkBySize 3
        |> Array.map (fun win -> win |> Array.map (fun arr -> arr |> Set.ofList))
        |> Array.map Set.intersectMany
        |> Array.map (Set.toList >> List.head)
        |> Array.map weighting
        |> Array.reduce (+)

printfn $"Part 2 : {y}"