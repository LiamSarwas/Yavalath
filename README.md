# Yavalath
Yavalath is a game engine and AI for the Yavalath board game. Yavalath is a game played on a 5 tile wide hex grid, with players who alternate placing pieces on hexes. The game is won by connecting four of your own hexes without connecting three first. If you connect three without connecting four, your opponent wins. These competing objectives make for an interesting challenge for AI to develop winning strategies.

## Engine 
This engine uses Monte Carlo Tree Search to expand the game tree, searching every turn for SearchDuration seconds, a const in the main.go file. With only 1 second of search time the engine performs rather poorly, though it still beats a completely random agent most of the time. Once it has 10+ seconds of search time per turn the engine improves greatly, and while still beatable is far more interesting to play against.



## Coordinate System

The game is played with moved described by a pairs (x, y) on an axial coordinate system as shown below. 
```
          y
          -4  
           -3  
            -2   
             -1               
x -4 -3 -2 -1  0  1  2  3  4  
                1   
                 2  
                  3  
                   4
```

## Completed Game

An example of a finished game in which Player 1 has constructed four in a row without getting three in a row.
```
    0 0 0 0 0    
   0 0 0 0 0 0   
  0 0 0 0 0 0 0  
 0 0 0 0 0 0 0 0 
0 0 0 0 2 0 0 0 0
 1 2 0 1 0 0 0 0 
  0 0 1 2 0 0 0  
   2 1 0 0 0 0   
    1 2 0 1 0 
```
