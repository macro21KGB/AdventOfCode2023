const fs = require("fs")

/**
 * @typedef {{id: number, rounds: string[], isImpossible: boolean}} Set;
 * @typedef {{color: string, number: number}} Round;
 *
 */


const MAX_RED = 12;
const MAX_GREEN = 13;
const MAX_BLUE = 14;

/**
 * @param {Round[][]} rounds
 */
const findMininumAmountOfDice = (rounds) => {

  const minAmount = {
    red: 0,
    blue: 0,
    green: 0
  }

  rounds.forEach(round => {
    round.forEach(item => {
      if (item.color == 'red' && item.number > minAmount.red)
        minAmount.red = item.number
      if (item.color == 'green' && item.number > minAmount.green)
        minAmount.green = item.number
      if (item.color == 'blue' && item.number > minAmount.blue)
        minAmount.blue = item.number
    })
  });

  console.log(minAmount)
  return minAmount.red * minAmount.blue * minAmount.green;
}

/**
 * @param {string} line
 * @returns Set
 */
const parseLine = (line) => {
  const id = +line.split(":")[0].replace("Game ", "");
  const sets = line.split(":")[1].split(";").map(el => el.trim())

  return {
    id,
    sets,
    isImpossible: false
  }
}


/**
 * @param { string[]} rounds
 * @returns {{color: string, number : number }}
 */
const parseRound = (rounds) => {
  const regex = /\d+ (red|green|blue)/g
  const roundsArray = []
  rounds.forEach((el) => {

    const amountOfColors = [];
    el.split(",").forEach(el => {
      const matched = el.match(regex)[0]
      const number = +matched.split(" ")[0]
      const color = matched.split(" ")[1]

      amountOfColors.push({
        number,
        color
      })

      roundsArray.push(amountOfColors)

    })

  });

  return roundsArray;

}

fs.readFile('./input.txt', 'utf8', (err, data) => {
  if (err) throw err;
  const lines = data.split("\n");

  let powerOfRoundSum = 0;
  lines.forEach(element => {
    try {
      const currentSet = parseLine(element);
      const rounds = parseRound(currentSet.sets)
      let illegalPull = false

      rounds.forEach((set) => {
        set.forEach(element => {
          if (illegalPull)
            return;
          switch (element.color) {
            case 'green':
              if (element.number > MAX_GREEN)
                illegalPull = true
              break;
            case 'blue':
              if (element.number > MAX_BLUE)
                illegalPull = true
              break;
            case 'red':
              if (element.number > MAX_RED)
                illegalPull = true
              break;
          }

        });
      })


      const power = findMininumAmountOfDice(rounds)
      powerOfRoundSum += power;


    } catch {
      return;
    }

  });

  console.log(powerOfRoundSum)
});



