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
    |> fun (index, _) -> index + 4

printfn "Part 1 Test 1: %i" (getMarker 4 "mjqjpqmgbljsphdztnvjfqwrcgsmlb")
printfn "Part 1 Test 2: %i" (getMarker 4 "bvwbjplbgvbhsrlpgdmjqwftvncz")
printfn "Part 1 Test 3: %i" (getMarker 4 "nppdvjthqldpwncqszvftbrmjlhg")
printfn "Part 1 Test 4: %i" (getMarker 4 "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg")
printfn "Part 1 Test 5: %i" (getMarker 4 "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw")

let res1 =
    File.ReadAllText "2022/F#/day-06.txt"
    |> getMarker 4
