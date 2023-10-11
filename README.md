<a name="readme-top"></a>

# GoWordle

## About the project

GoWordle is a terminal version of the popular game Wordle. It is intended to be just a practice project in order to test the new language that I'm learning now, Go.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Getting started

Those are the instructions to setting up the project locally. To get a local copy up and running follow these simple steps.

### Prerequisites

In order to use GoWordle you need to have Go compiler installed on your computer and have acces to a CLI.

### Installation

1. Clone the repo

    ```sh
    git clone https://github.com/jordicido/gowordle.git

2. Execute the command

   ```go
   go build && ./wordle

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## How to use it


You will run the application and will be able to play the game. It will appear the main menu, where you can select what to do. You can play a game, read the instructions of the game, see the match history (if any stored) or exit the application.

## How is this project developed

To develop this project I used two main technologies, Go language and SQLite to store some small amount of data.

I stored in a table words all the possible words of 5 letters that exists in english. In every game one of them is randomly select to be the hidden word. Every time you try to guess the word the system check that this word exists and mark every letter of the word you tried with yellow (in case the letter is misplaced but exists in the word) or green (in case the letter is correctly placed). Your mission is to guess the hidden word within 6 turns.

Every time you finish a game, the data of this game is stored in match_history table in the database, so you can check your game history every time you want.
