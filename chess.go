package main

import (
    "fmt"
    "github.com/malbrecht/chess"
)


func main() {
    var player1name string
    var player2name string
    fmt.Println("You can quit this match at any point by pressing CTRL C.")
    fmt.Println("Enter player 1's name:")
    fmt.Scanf("%s", &player1name)
    fmt.Println("Enter player 2's name:")
    fmt.Scanf("%s", &player2name)

    board := chess.MustParseFen("")
    i := 1
    for i>0{
        printboard(board)

        if(i%2==0){
            fmt.Println("It is " + player2name + "'s turn!")
        }else{
            fmt.Println("It is " + player1name + "'s turn!")
        }

    board = moves(board)
    i++
    }
}

func moves(b *chess.Board) *chess.Board{
    fmt.Println("This is the list of moves you can make:")
    fmt.Println(b.LegalMoves())
    fmt.Println("Moves are constructed like {to from promotion}.")
    fmt.Println("For the promotion: 0 is no promotion, 2/3 is a pawn, 4/5 is a knight, 6/7 is a bishop, 8/9 is a rook, 10/11 is a queen and 12/13 is a king.")
    fmt.Println("Which piece would you like to move? (Please enter coordinates of a valid move, or no piece will be moved. EXAMPLE: \"a1b1\")")
    var s string
    fmt.Scanf("%s", &s)
    move, _ := b.ParseMove(s)
    return b.MakeMove(move)
}

func printboard(b *chess.Board){
    pieces := b.Piece
    var figurines[] rune
    figurines = append(figurines, '\u1D00', '\u0299', '\u1D04', '\u1D05', '\u1D07', '\u0192', '\u0262', '\u029C')
    for _, piece := range pieces {
        intpiece := int(piece)
        switch intpiece{
            case 2: figurines = append(figurines, '\u265F')
            case 3: figurines = append(figurines, '\u2659') 
            case 4: figurines = append(figurines, '\u265E')
            case 5: figurines = append(figurines, '\u2658')
            case 6: figurines = append(figurines, '\u265D')
            case 7: figurines = append(figurines, '\u2657')
            case 8: figurines = append(figurines, '\u265C')
            case 9: figurines = append(figurines, '\u2656')
            case 10: figurines = append(figurines, '\u265B')
            case 11: figurines = append(figurines, '\u2655')
            case 12: figurines = append(figurines, '\u265A')
            case 13: figurines = append(figurines, '\u2654')
            default: figurines = append(figurines, '\u25A1')
        }
    }
    fmt.Printf(" " + "%q\n", figurines[0:8])
    fmt.Printf("1" + "%q\n", figurines[8:16])
    fmt.Printf("2" + "%q\n", figurines[16:24])
    fmt.Printf("3" + "%q\n", figurines[24:32])
    fmt.Printf("4" + "%q\n", figurines[32:40])
    fmt.Printf("5" + "%q\n", figurines[40:48])
    fmt.Printf("6" + "%q\n", figurines[48:56])
    fmt.Printf("7" + "%q\n", figurines[56:64])
    fmt.Printf("8" + "%q\n", figurines[64:72])
}
