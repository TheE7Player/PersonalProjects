console.log("RUNNING MM SCRIPT")
// Get fast information for CS:GO MM collection

// First we need to get the user to target - what is handy is the name tag as a href with the id!
// This means we don't need to do extra work incase X has changed their name tag or smth

// Get the player <a> tag as the name tag title
var Target = document.getElementsByClassName("whiteLink persona_name_text_content")

if (Target.length > 1)
    throw "Target element contains more than 1 instance, this shouldn't happen!";

if (Target.length == 0)
    throw "Target element was not found at all - this is a problem huston!";

Target = Target[0].href
console.log(`Assuming of validating from user ${Target}... Starting...`)


// Now we get all the matches!
var Matches = document.querySelector("tbody").children

if (Matches.length < 1)
    throw "Expected table to look at was not found"

console.log(`Found table - now processing ${Matches.length - 1}...`)

// We start at 1 as we avoid the header (Map, Match Results)
var idx = 1
var game = null

function GetMapData(g)
{
    // Grab the data from the tbody
    var level = g.children[0].children[1].tBodies[0].children[0].outerText
    var time = g.children[0].children[1].tBodies[0].children[1].outerText
    var wait = g.children[0].children[1].tBodies[0].children[2].outerText
    var total = g.children[0].children[1].tBodies[0].children[3].outerText
    var date = null
    // Do processing to clean it up
    
    // Map name clean-up: Remove the words 'Competitive' and 'Scrimmage' if exists
    level = level.replace("Competitive", "")

    // If the map name contains scrimmage, remove it
    if(level.include("Scrimmage"))
        level = level.replace("Scrimmage", "")

    // Finally, remove all the whitespace or gaps in the string (if any)
    level = level.trim()

    // Time cleanup
    var dt_split = time.split(" ")
    date = dt_split[0].trim()
    time = dt_split[1].trim()
    dt_split = null

    // Clean up the wait and duration of the match by turning gaps into an array and select the last element
    wait = wait.split(" ")
    wait = wait[wait.length-1].trim()

    total = total.split(" ")
    total = total[total.length-1].trim()

    // Finally return back as a dictionary
    return {
        map: level,
        date: date, time: time,
        waittime: wait,
        duration: total
    }
}

function GetOtherData(game)
{
    // game.children[1].children[0].children[0].children[6] ~ Score
    // Team A: [1..5] | Team B: [7..11]
    
    // game.children[1].children[0].children[0].children[IDX]
    // to get steam url: -> .children[0].children[0].children[0].href

    var Player = null
    var Score = game.children[1].children[0].children[0].children[6].outerText
    var Won = false

    var scoreSplit = Score.trim().split(":")
    
    if(scoreSplit.length != 2)
        throw "Score split produced higher or lower length than 2 (2 is max!)"

    var TeamA = parseInt(scoreSplit[0])
    var TeamB = parseInt(scoreSplit[1])

    var Team_A_Won = TeamA > TeamB
    var Tie = TeamA == 15 && TeamB == 15

    var gameResult = "Lost"
    var InTeamA = true

    for(var x = 1; x < 12; x++)
    {

        if(x == 6) continue;

        if(x > 6 && InTeamA)
            InTeamA = false

        Player = game.children[1].children[0].children[0].children[x]

        if(Player.children[0].children[0].children[0].href != Target)
          continue

        // Found the player we require, lets add an array over it now

        var mvpCount = 0
        var mvpText = Player.children[5].outerText.trim()

        if (mvpText == "★")
            mvpCount = 1
        else
            try {
                mvpCount = parseInt(mvpText.replace("★", " "))
            } catch(err)
            {
                mvpCount = 0
            }
            
            if(Number.isNaN(mvpCount))
                mvpCount = 0

            // Inner class order: Name (outerText) [0], Ping, K, A, D, MVP Count, HS Per, Score[7]
        Player = [
            Player.children[1].outerText,
            Player.children[2].outerText,
            Player.children[3].outerText,
            Player.children[4].outerText,
            mvpCount,
            (Player.children[6].outerText.trim() == "") ? "0%" : Player.children[6].outerText,
            Player.children[7].outerText
        ]

        break
    }

    if (Team_A_Won && InTeamA)
        gameResult = "Won"
    else if (Tie)
        gameResult = "Tie"

    return {
        info: Player,
        result: gameResult,
        score: Score.trim().replace(" : ", ":")
    }
}

output = [["Map", "Date", "Time", "Waiting Time", "Duration", "Result", "Score",
           "Ping", "Kills", "Assists", "Death", "MVPS", "HS Percentage", "Points"]]
mapData = null, gameData = null


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

console.log("NOW GENERATING STRING")

var longStr = ""

for(var i = 0; i < output.length; i++)
{
    for(var j = 0; j < output[i].length; j++)
    {
        if(j > 0) 
        { longStr = longStr.concat(", ") }
            
        longStr = longStr.concat(output[i][j])          
    }
        
    longStr = longStr.concat("\n")
}

console.log("ENDED MM SCRIPT - Save the following as an csv")

console.log(longStr)