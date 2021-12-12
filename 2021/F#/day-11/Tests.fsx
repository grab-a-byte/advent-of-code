#load "./Source.fsx"

open Source

open System

let sampleInput =
    """5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526"""
        .Split(Environment.NewLine)
    |> Array.toList

// let testResult1 = getFlashCount sampleInput 10

// printfn "Test1 : Given sample input, result should be 204 %i" testResult1

let requireFlash =
    """6594254334
3856965822
6375667284
7252447257
7468496589
5278635756
3287952832
7993992245
5957959665
6394862637"""
        .Split(Environment.NewLine)
    |> Array.toList

let sample = getFlashCount requireFlash 1
