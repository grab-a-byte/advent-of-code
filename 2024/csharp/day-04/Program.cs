var lines = File.ReadAllLines("./input.txt");

// test.txt expected 2
// input.txt expected 371
//-----Part One------//
var numLines = lines.Select(x => x.Split(" ").Select(n => int.Parse(n))).ToList();
int resultOne = 0;
foreach (var line in numLines)
{
    if(LineIsSafe(line)) resultOne++;
}
Console.WriteLine($"Day Two - Part One: {resultOne}");

//-------Part Two ---------//
int resultTwo =0;
foreach (var line in numLines)
{
    if(LineIsSafe(line)) {
        resultTwo++;
        continue;
    }
    for(int i = 0; i < line.Count(); i++){
        var start = line.Take(i);
        var end = line.Skip(i+1);
        var shortenedLine = start.Concat(end);
        if(LineIsSafe(shortenedLine)){
            resultTwo++;
            break;
        }
    }
}

Console.WriteLine($"Day Two - Part Two: {resultTwo}");

//Helpers
bool LineIsSafe(IEnumerable<int> line){
    var zipOne = line.SkipLast(1);
    var zipTwo = line.Skip(1);
    var zipped = zipOne.Zip(zipTwo);
    Func<int, int, bool> dirCheck = zipOne.First() < zipTwo.First() ?
        (int a, int b) => a + 1 <= b && a + 3 >= b :
        (int a, int b) => a - 1 >= b && a - 3 <= b;

    return zipped.All((tup) => dirCheck(tup.First, tup.Second));
}
