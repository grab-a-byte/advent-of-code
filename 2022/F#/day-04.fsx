let input =
    """2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8"""

// Part 1
let lineToSetRanges (line: string) =
    let parts = line.Split ","
    let partsA = parts[ 0 ].Split "-"
    let partsB = parts[ 1 ].Split "-"
    let parseNums (arr: string array) = (arr[0] |> int, arr[1] |> int)

    let (aStart, aEnd) = parseNums partsA
    let (bStart, bEnd) = parseNums partsB

    let setA = seq { aStart..aEnd } |> Set.ofSeq
    let setB = seq { bStart..bEnd } |> Set.ofSeq

    (setA, setB)

let parsedInput = input.Split "\n" |> Array.map lineToSetRanges

let eitherSubset (a, b) =
    (Set.isSubset a b) || (Set.isSubset b a)

let res1 =
    parsedInput
    |> Array.filter eitherSubset
    |> Array.length

printfn $"Part 1: {res1}"


// Part 2
let setIntersectsNonZero (a, b) = Set.intersect a b |> Set.count > 0

let res2 =
    parsedInput
    |> Array.filter setIntersectsNonZero
    |> Array.length

printfn $"Part 2: {res2}"
