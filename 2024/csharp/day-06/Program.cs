var rows = File.ReadAllLines("./input.txt");

var startingPos = new Position(-100, -100);
List<Position> blockers = [];
for(int y = 0; y < rows.Length; y++)
{
    for (int x = 0; x < rows[0].Length; x++)
    {
        if(rows[y][x] == '^')
        {
            startingPos = new Position(x,y);
        }
        else if (rows[y][x] == '#'){
            blockers.Add(new(x,y));
        }
    }
}

List<Direction> directions = [
    new(0, -1),
    new(1, 0),
    new(0, 1),
    new(-1, 0)
];

Direction GetDirection(int number) => directions[number % directions.Count];
List<Position> path = [startingPos];
int directionChangeCount = 0;

while(true){
    var direction = GetDirection(directionChangeCount);
    var moveFrom = path.Last();
    var nextSpot = new Position(moveFrom.X + direction.X, moveFrom.Y + direction.Y);
    if(blockers.Contains(nextSpot)){
        directionChangeCount++;
        continue;
    }
    if(nextSpot.X < 0 || nextSpot.X >= rows[0].Length || nextSpot.Y < 0 || nextSpot.Y >= rows.Length){
        break;
    }
    path.Add(nextSpot);
}

var result = path.Distinct().Count();
Console.WriteLine($"Day 6 - Part 1 : {result}");

record Position(int X, int Y);
record Direction(int X, int Y);