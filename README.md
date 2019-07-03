# Text Based Adventure Game

Something I wrote to learn golang during a vacation.

### How to Run

```
go get github.com/dorthu/tbaGame
cd ~/go/src/github.com/dorthu/tbaGame
go run main.go
```

### Controls

Arrow keys move
Space bar interacts

### How To

See any "version 1" room file in `rooms` for an example of how to correctly
format game data.  Rooms are defined as an array of spaces, an array of items,
and a grid associating positions on the screen to spaces and their contents.
The grid is a 2-dimensional yaml array starting at the top-left corner of the
screen and going across the columns to the left.  The screen is 9x5 (cols x rows).
