open System.IO

let input = "mjqjpqmgbljsphdztnvjfqwrcgsmlb"

let charsAreUnique (str: string) =
    (str |> Set.ofSeq |> Set.count) = str.Length

let getMarker windowSize (input: string) =
    input.ToCharArray()
    |> Array.windowed windowSize
    |> Array.map System.String
    |> Array.mapi (fun i x -> (i, charsAreUnique x))
    |> Array.filter (fun (i, b) -> b)
    |> Array.head
    |> fun (index, _) -> index + windowSize

printfn "Part 1 Test 1: %i" (getMarker 4 "mjqjpqmgbljsphdztnvjfqwrcgsmlb")
printfn "Part 1 Test 2: %i" (getMarker 4 "bvwbjplbgvbhsrlpgdmjqwftvncz")
printfn "Part 1 Test 3: %i" (getMarker 4 "nppdvjthqldpwncqszvftbrmjlhg")
printfn "Part 1 Test 4: %i" (getMarker 4 "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg")
printfn "Part 1 Test 5: %i" (getMarker 4 "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw")

let res1 = File.ReadAllText "2022/F#/day-06.txt" |> getMarker 4

printfn $"Part 1 : {res1}"

// Part 2
printfn "Part 2 Test 1: %i" (getMarker 14 "mjqjpqmgbljsphdztnvjfqwrcgsmlb")
printfn "Part 2 Test 2: %i" (getMarker 14 "bvwbjplbgvbhsrlpgdmjqwftvncz")
printfn "Part 2 Test 3: %i" (getMarker 14 "nppdvjthqldpwncqszvftbrmjlhg")
printfn "Part 2 Test 4: %i" (getMarker 14 "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg")
printfn "Part 2 Test 5: %i" (getMarker 14 "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw")

let res2 = File.ReadAllText "2022/F#/day-06.txt" |> getMarker 14
printfn $"Part 2 : {res2}"
