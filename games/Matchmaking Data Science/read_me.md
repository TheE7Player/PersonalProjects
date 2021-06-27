# Analysing Data Science with CSGO Matches

> This is basic information to be gathered and does not give any additional information on performance metrics

So, Covid is still a thing. A year onwards can be boring and unmotivating. But I like to analyse data from time to time, looking for patterns and key points in information.

So, why not do it from one of my favourite games of *all* time... **CS:GO**.

---

## Problem

I want to be able to find patterns or key points to figure out which maps I need to spend more time on, find out when the best time to look for a competitive game when I play matchmaking and find how bad I am truly at, for winning games.

I will need to find a way to gain this information. Lucky, Steam supplies all competitive games within 5 years of history (This does mean I am unable to fetch data from November 2017, and that goes higher every year - remember, 5 year’s worth of data!).

> This is due to Valve deleting match data as of the new laws regarding how one entity can hold information on a person, view it from [here](https://www.reddit.com/r/GlobalOffensive/comments/bguq4y/valve_apparently_deleted_all_match_data_prior_to/).

But also, I need to find the best tool or program to analyse the data I will scrape from. So I will need to find the best tool to allow me do perform simple tasks such as storing, fetching, filtering and applying methods to find such patterns etc.

### Finding the best tool

You can view your history of matchmaking games from a range of 5 years from the recent game played (so about 350+ games worth of data!) from [here]().

I know I will have to web scrap the data, I know how to use `beautifulsoup4` with `Python` - but that will mean extra scripting, where I can do it from the browser indirectly, which is the native way to achieve it using `Web Snippets` (Chrome browsers).

Let us evaluate between both ideal tools to achieve this:

## Web Scraping Tool

**Python**

- Easy syntax and methods to iterate through data

- Accessible without need of extra tools or compilers

- Scripting language, does not require a `main` entry-point

- Ability to work of OS layers, which is good for storing files or reading

- Good libraries available such as `beautifulsoup4` to scrape web pages

- Can be a bit annoying, need to learn the library before hand

- Interpreter, needs to run procedurally from top to bottom - slow.

**Javascript** (Using Web Snippet):

- Available straight out the bat, no installs or dependencies required (just a web browser needed)

- Code/Script with `JavaScript`, meaning we can fetch data from the page directly

- Dynamic just like `Python`, still runs top to bottom

- Uses the `V8` engine, quick fast with direct access to web page

- Unable to access `OS layers` natively but can print results to the console (You can with `Node.js`, but I want to just use the web browser as it is!)

So I will use `JavaScript` to scrap the data from the page. Easy as that, no extra installations required - just a browser!

## Data Analysing Program/Tool

**Python**

- Multiple Library support for `Data Science` needs

- But requires time and commitment for the chosen library

- Easy methods to store or manipulate data for patterns

- Slow, as it is an interpreter

**.Net (C#)**

- Good language with an optimised `compiler`

- Good use of memory, comes with a `GCC`

- Syntax can be annoying and can take ages to achieve a goal depending on complexity

**Java**

- Just like `C#`, but different in syntax method and control

- Unable to make good usage of memory variables (Value only)

- Somewhat old, and tends to be slow depending on complexity

- Doesn't support good usage of things that `C#` has to offer

**C/C++**

- Super quick, no bloat added on top of the exe

- Lower layer control, but is more syntax heavy

- No `GCC` available, must control memory with memory management

- Lower in executable size depending on `compiler` and `complexity of code`

**Go**

- Interesting syntax structure

- Interesting way of managing memory

- Forces good programming conventions

- Forces you to consider about what variables to use

- Typically faster while also being a compiler language

I *do* have experience on most of these available languages, such as `C#`, `Java` and a little of `C++`, but using `Go` will be the best option to go for. `Python` would have been second if it included usage of optimisations and a compiler.

---

### Solution

I will use `Go` for the analysing of the data and `JavaScript` as the data collector to feed the information for `Go`.

---

# Grabbing the data (Web Scrapping)

> You can find the snippet used [here]([PersonalProjects/Get_MM_Details.js at 23c80d2accdce2d7ebdf3ba64efd808f5256653f · TheE7Player/PersonalProjects · GitHub](https://github.com/TheE7Player/PersonalProjects/blob/23c80d2accdce2d7ebdf3ba64efd808f5256653f/games/Matchmaking%20Data%20Science/Web%20Snippet/Get_MM_Details.js)).

The first step is to grab the data. We can just open the saved html from Steam when visiting the personal game data when logged in.

> You can right click the page when you have it open and save it to a location to web scrap it or save it for later usage!

You can create a snippet quick easily, a simple google search for the browser you are going to use will tell you how to access it. Snippets stay in the browser, but it is recommended to save it externally from time to time.

Once you run the script, the result will be put into a `CSV` format, with a `comma` as its `delimiter`. You can then put it into a csv file (or any other suitable program) to build up a data set for analysing.

I will not go over *each* part of the script but will go over the main concept or logic that makes the thing run.

### Step 1: Finding myself in each game

It is important to find yourself in each game, you may be the top or the bottom or in between. But I will need to solve the following:

- What If I change my name in 5 years of the data?

I do not change my nickname often funny enough, but the issue still raises up. So, we need to dig into how we can solve it. It is easy, we just need to find our steam URL and compare it to see if we found ourselves - it means we can ignore other checks!

Let us look at the first part, finding myself in the data:

```javascript
var Target = document.getElementsByClassName("whiteLink persona_name_text_content")

if (Target.length > 1)
    throw "Target element contains more than 1 instance, this shouldn't happen!";

if (Target.length == 0)
    throw "Target element was not found at all - this is a problem huston!";

Target = Target[0].href
console.log(`Assuming of validating from user ${Target}... Starting...`)
```

We create a variable called `Target` which looks for an element with the class name `whiteLink persona_name_text_content` which holds the url of our current user that is signed in. This will hold the url to look for, remember that the URL remains the same through out!

We then evaluate the length; it should only be just `1`.

We then re-assign `Target` to its `href` value - the `url` of our target. Now we start!

> Note: Before saving, I clicked a button which said `Load More Matches`, this is important as it only shows 9 matches by default... I kept clicking it until the button does not exist.

### Step 2: Finding the matches and iteration

Now we need to find the table and iterate it over...

```javascript
// Now we get all the matches!
var Matches = document.querySelector("tbody").children

if (Matches.length < 1)
    throw "Expected table to look at was not found"

console.log(`Found table - now processing ${Matches.length - 1}...`)
```

Thankfully, there is only 1 table that exists in the web page. So, we can do a simple `querySelector` and target an `tbody` element and grab its children’s elements. Then we apply a simple validator to check if it was successful.

Now we must create a `2D Array` - just like excel - to hold the `csv` data. So we create an `array of arrays` variable called `output` with `pre-defined headers`.

```javascript
output = [["Map", "Date", "Time", "Waiting Time", "Duration", "Result", "Score",
           "Ping", "Kills", "Assists", "Death", "MVPS", "HS Percentage", "Points"]]
```

Now we make good usage of caching, so we now create 2 extra objects to hold each match’s map and game data...

```javascript
mapData = null, gameData = null
```

Now, we can iterate over each table object in the table we found earlier on.

> Note: There are 2 global variables create at the near-top as well...
> 
> ```javascript
> var idx = 1
> var game = null
> ```

```java
console.log("PROCESSING... PLEASE WAIT")
while ( (idx + 1) < Matches.length)
{
    // Fetch each game
    game = Matches[idx]

    // Grab data from (that) game
    mapData = GetMapData(game)
    gameData = GetOtherData(game)

    // Prepare it to add to the list (push)
    output.push([
        // Add the row
        mapData['map'], mapData['date'], mapData['time'], mapData['waittime'], mapData['duration'], gameData['result'], gameData['score'],
        gameData['info'][0], gameData['info'][1], gameData['info'][2], gameData['info'][3], gameData['info'][4], gameData['info'][5], gameData['info'][6]
    ])  

    idx += 1
}
```

We can ignore the extra functions for now, but we have a simple `while` loop to goes through each table element and increments a pointer by 1 if it has not reached the end yet.

The `game` variable holds that current iterations match, to pass onto the other functions.

`GetMapData` returns a `Dictionary` or `Map` of the Map details which contains the time, date, map, waiting time and game time (game duration).

`GetOtherData` returns another `Dictionary` or `Map` but instead, returns the performance/score of the game played, such as the Ping, Score, Kills, Deaths etc

Then we push that gathered information into the `output` variable earlier on and push an array onto it.

We should not also forget the increment the pointer! Or we will crash the browser.

Remember, you can view the entire script [here]([PersonalProjects/Get_MM_Details.js at 23c80d2accdce2d7ebdf3ba64efd808f5256653f · TheE7Player/PersonalProjects · GitHub](https://github.com/TheE7Player/PersonalProjects/blob/23c80d2accdce2d7ebdf3ba64efd808f5256653f/games/Matchmaking%20Data%20Science/Web%20Snippet/Get_MM_Details.js)). The other functions make it easier to get the information, it just finds it and stores it.

---

### Step 3: Using GO to analyse the Data

> The scripts used for this section is [here]([PersonalProjects/games/Matchmaking Data Science/Script at main · TheE7Player/PersonalProjects · GitHub](https://github.com/TheE7Player/PersonalProjects/tree/main/games/Matchmaking%20Data%20Science/Script)).
> 
> It also includes the data set called `csgo_data.csv`

I like how `Go` operates data and forces `code conventions` as a standard - so no messing around with variables that do not get used!

The only thing I will be going over in this section is to assign the data and to filter it. You can view the code in its entirety to see how it works, it will be too long to write out each section of code.

The 2 scripts that are the brains are the `setup` and `filter` folder, or as `Go` calls it: `Packages`.

## Setup Package

This package holds the logic to read the `csv` file and stores it into an `2D array of arrays of String` - fun stuff.

`Go` has an interesting library to do this straight out the bat called `encoding/csv`.

So the first step is to create the package and import the require libraries:

```go
package setup

import (
    "encoding/csv"
    "fmt"
    "os"
    "strings"
)
```

Nothing too complex here, but it can be overwhelming if your new to `Go`.

`package` tells `Go` that this file here is a part of the `setup` `package` (folder).

We then tell `Go` which dependencies we need to run this script...

- `encoding/csv` library is the library to read in the csv file

- `fmt` library is to print stuff out into the output stream (console window)

- `os` library is used to perform `OS` operations like reading in this example

- `strings` library is used to do conversions of data into strings if we need to

We have 2 functions: `Setup` and `Setup_Table`:

#### Setup Function

```go
func Setup(file string) [][]string {
	// Open File
	csvFile, err := os.Open(file)

	// Show an error to console if an error has occured (err is nil, if not)
	if err != nil {
		fmt.Println("Error happen will reading csv: ", err)
	}

	// Create the reader object using the file as IO stream
	reader := csv.NewReader(csvFile)

	// Get the data into a variable called 'data', ignoring the error variable using underscore '_'
	data, _ := reader.ReadAll()

	// Finally, we return back the data to the caller
	return data
}
```

What I like about `Go` is the ability to hold `Error Variables`. We do not need to wrap it around a `try catch` block at all... which is beneficial in the performance run.



So we call the `OS` library and find in the file location where:

`csvFile` will contain the csv metadata of the file

`err` will be `nil` if no errors happened while attempting to read, if so, this will show why we got an error doing so.



We now create a reader object - `csv.NewReader`- to read the text of the data, instead of its metadata!



We then use that reader object to feed into the `data` variable, at this point we do not care about an error - so we tell `Go` to simply ignore it but applying an underscore instead (`_`). Now we return `data` back to the `caller` of the `function`.



### Setup Table Function

This here is how we construct the information...

We feed in the new `2D Array` into a `Map`. Which will contain:

`Key` ~ The map name of the match (`Dust II`, `Mirage`, `Nuke`, `Overpass` etc)

`Value` ~ The information from that match (`Ping`, `Kills` etc)  of which, is an `String Array`


