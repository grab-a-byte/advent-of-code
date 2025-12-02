using System.Data;

var lines = File.ReadAllLines("./input.txt");
var orderedRules = lines.TakeWhile((x) => !string.IsNullOrWhiteSpace(x));
var input = lines.Skip(orderedRules.Count() + 1);

Dictionary<int, HashSet<int>> rules = [];

foreach (var rule in orderedRules)
{
    var parts = rule.Split('|');
    var key = int.Parse(parts[0]);
    var value = int.Parse(parts[1]);
    if (rules.TryGetValue(key, out var set))
    {
        set.Add(value);
    }
    else
    {
        rules[key] = new([value]);
    }
}

List<List<int>> validRows = [];
List<List<int>> invalidRows = [];
foreach (var row in input)
{
    var parts = row.Split(',').Select(int.Parse).ToList();
    if (ValidateRow(parts))
    {
        validRows.Add(parts);
    }
    else
    {
        invalidRows.Add(parts);
    }
}

var result = validRows.Select(x => x[x.Count / 2]).Sum();
Console.WriteLine($"Day Five - Part 1: {result}");

//TODO: Part 2
List<List<int>> fixedRows= [];
foreach(var badRow in invalidRows){
    var newRow = badRow;
    while(!ValidateRow(newRow)){
    newRow.Sort(delegate(int x, int y)
        {
            if(rules.TryGetValue(x, out var check)){
                if(check.Contains(y)) return -1;
                if(!check.Contains(y)) return 0;
            }
            return 1;
        });
    }
    fixedRows.Add(newRow);
}

foreach(var x in fixedRows){
        Console.Write('[');
        foreach(var y in x){
            Console.Write($"{y},");
        }
        Console.Write(']');
        Console.WriteLine();
}

var result2 = fixedRows.Select(x => x[x.Count / 2]).Sum();
Console.WriteLine($"Day Five - Part 2: {result2}");


bool ValidateRow(List<int> parts){
    for (int i = 0; i < parts.Count; i++)
    {
        var before = parts.Take(i);
        if (rules.TryGetValue(parts[i], out var checks))
        {
            if (!before.All(x => !checks.Contains(x)))
            {
                return false;
            }
        }
    }
    return true;
}