```
ISLAND OF SECRETS   TIME REMAINING:1000
----------------------------------------
YOUR STRENGTH = 100 YOUR WISDOM = 35
----------------------------------------

YOU ARE ON A LEAFY PATH.
----------------------------------------

LET YOUR QUEST BEGIN

WHAT WILL YOU DO?
```

This is a Go implementation of Island of Secrets, a text adventure game published by Usborne Publishing in 1984. It was published in a book in the form of terse BASIC code which you had to type out and save to tape or disc before being able to play. Obfuscation in the code helped avoid the player from knowing too much about the game before playing. 

A PDF copy of the book is still available from Usborne's website here: https://usborne.com/gb/books/computer-and-coding-books.

All credit must go to the authors: Jenny Tyler and Les Howarth, and copyright holders Usborne Publishing.

The code borrows heavily from the Go version of the similar text adventure game The Mystery of Silver Mountain located at https://github.com/fivegreenapples/go-mountain.

Install via `git clone` and `go build`:
```
git clone https://github.com/jameshollandusa/go-island-of-secrets
cd go-island-of-secrets
go build -o island-of-secrets
./island-of-secrets
```

For Windows users:

```
go build -o island-of-secrets.exe
island-of-secrets.exe
```

# Affordances for an Easier Life

Replicating the original program is all very well but does lead to some slightly frustrating game play. Perhaps that was part of the charm back in the eighties, but if memory serves, it was equally if not more frustrating back then.

For example, the `SAVE GAME` and `CONTINUE A SAVED GAME` operations both require typing out a filename plus a few other keystrokes. So just to save your progress and continue is a bit of a hassle. I've added a few niceties to make game play more pleasant:

1. Input commands are case insensitive.

   This just means there is a `strings.ToUpper(...)` wrapping any input as the BASIC only understands UPPERCASE COMMANDS.

2. You can specify a commands file.

    By writing a sequence of newline separated commands into a file, you can have the program run these before accepting manual input from STDIN. Specify the filename as the first command line argument to the program:

    ```
    ./island-of-secrets commands.txt
    ```

    You can also comment your command file. Any line starting with a `#` will be ignored.

3. You can seed the RNG.

    Some elements of the game are generated at game start using a random number generator. If you use a command file as above, you'll want to seed the RNG so these game elements are predictable. Use a second command line argument:

    ```
    ./island-of-secrets commands.txt 123
    ```

# Notes on Playing the Game

While trying not to give away any spoilers, it's worth noting that you can't complete this game without the book. The book contains some critical information that isn't available just from playing the game.

Good luck!
